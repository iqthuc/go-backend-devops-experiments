package product

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type GetProductsQueryParams struct {
	limit   int
	offset  int
	Filters Filters
	SortBy  SortBy
}

type Repository interface {
	GetProducts(ctx context.Context, params GetProductsQueryParams) (GetProductsResult, error)
	GetProductByID(ctx context.Context, id int) (ProductDetail, error)
	GetProductVariantsByID(ctx context.Context, id int) ([]ProductVariant, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetProducts(ctx context.Context, params GetProductsQueryParams) (GetProductsResult, error) {
	// note for sql: add tsvector for name field and add GIN index for tsvector
	var query strings.Builder
	query.WriteString(`
	SELECT p.id, p.name, COALESCE(SUM(stock_quantity),0),COALESCE(SUM(sold_quantity),0), COALESCE(base_price,0), COUNT(*) OVER() AS total_count
	FROM products as p
	LEFT JOIN product_variants as pv
	ON p.id=pv.product_id
	WHERE ($3::int IS NULL OR category_id = $3)
		AND ($4::varchar IS NULL OR name_search @@ plainto_tsquery($4))
		AND ($5::float IS NULL OR base_price >= $5)
		AND ($6::float IS NULL OR base_price <= $6)
	GROUP BY p.id
	`)

	if params.SortBy.Field != "" {
		query.WriteString(fmt.Sprintf(" ORDER BY %s %s", params.SortBy.Field, params.SortBy.Order))
	}
	query.WriteString(" LIMIT $1 OFFSET $2")

	rows, err := r.db.QueryContext(
		ctx,
		query.String(),
		params.limit, //$1
		params.offset,
		sql.NullInt64{Int64: int64(params.Filters.CategoryID), Valid: params.Filters.CategoryID != 0},
		sql.NullString{String: params.Filters.Keyword, Valid: params.Filters.Keyword != ""},
		sql.NullFloat64{Float64: params.Filters.PriceMin, Valid: params.Filters.PriceMin != 0},
		sql.NullFloat64{Float64: params.Filters.PriceMax, Valid: params.Filters.PriceMax != 0},
	)

	if err != nil {
		return GetProductsResult{}, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	products := []Product{}
	totalCount := 0

	for rows.Next() {
		product := Product{}
		count := 0
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Stock,
			&product.Sold,
			&product.BasePrice,
			&count,
		); err != nil {
			return GetProductsResult{}, fmt.Errorf("failed to scan product row: %w", err)
		}
		if len(products) == 0 {
			totalCount = count
		}

		products = append(products, product)
	}

	result := GetProductsResult{
		products:   products,
		totalCount: totalCount,
	}

	if err := rows.Err(); err != nil {
		return GetProductsResult{}, fmt.Errorf("error iterating product rows: %w", err)
	}

	return result, nil
}

func (r *repository) GetProductByID(ctx context.Context, id int) (ProductDetail, error) {
	query := `
	SELECT p.id, p.name, c.name as category, SUM(pv.stock_quantity) as stock, SUM(pv.sold_quantity) as sold, COALESCE(p.base_price,0)
	FROM product_variants as pv
	JOIN products as p
	ON p.id=pv.product_id
	JOIN categories as c
	ON p.category_id=c.id
	JOIN brands
	ON p.brand_id=brands.id
	WHERE p.id=$1
	GROUP BY p.id, c.id, brands.id
	ORDER BY p.id
	`
	var p ProductDetail
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&p.ID, &p.Name, &p.Category, &p.Stock, &p.Sold, &p.BasePrice)
	if err != nil {
		return ProductDetail{}, fmt.Errorf("data base query error: %w", err)
	}

	return p, nil
}
func (r *repository) GetProductVariantsByID(ctx context.Context, id int) ([]ProductVariant, error) {
	query := `
	SELECT pv.id, colors.name as color, sizes.name as size, pv.price , pv.stock_quantity, pv.sold_quantity
	FROM product_variants as pv
	JOIN products as p
	ON p.id=pv.product_id

	LEFT JOIN product_variant_colors as pvc
	ON pv.id = pvc.product_variant_id
	LEFT JOIN colors
	ON colors.id = pvc.color_id

	LEFT JOIN product_variant_sizes as pvs
	ON pv.id = pvs.product_variant_id
	LEFT JOIN sizes
	ON sizes.id = pvs.size_id

	WHERE pv.product_id=$1
	ORDER BY pv.id
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	var productVariants []ProductVariant
	for rows.Next() {
		pv := ProductVariant{}
		if err := rows.Scan(&pv.ID, &pv.Color, &pv.Size, &pv.Price, &pv.Stock, &pv.Sold); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		productVariants = append(productVariants, pv)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}

	return productVariants, nil
}
