package admin

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository interface {
	GetProducts(ctx context.Context, limit int, offset int, category_id int) ([]Product, error)
	CountProducts(ctx context.Context, category_id int) (int, error)
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

func (r *repository) GetProducts(ctx context.Context, limit int, offset int, category_id int) ([]Product, error) {
	query := `
	SELECT p.id, name, COALESCE(SUM(stock_quantity),0),COALESCE(SUM(sold_quantity),0)
	FROM products as p
	LEFT JOIN product_variants as v
	ON p.id = v.product_id
	WHERE $3::int IS NULL OR category_id = $3
	GROUP BY p.id
	ORDER BY p.id
	LIMIT $1
	OFFSET $2;
	`
	rows, err := r.db.Query(
		query,
		limit,
		offset,
		sql.NullInt64{Int64: int64(category_id), Valid: category_id != 0},
	)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		product := Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Stock, &product.Sold); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	return products, nil
}

func (r *repository) CountProducts(ctx context.Context, category_id int) (int, error) {
	var total int
	query := `SELECT COUNT(*)
	FROM products
	WHERE $1::int IS NULL OR category_id = $1
	`
	fmt.Println(category_id, "id")
	err := r.db.QueryRowContext(
		ctx,
		query,
		sql.NullInt64{Int64: int64(category_id), Valid: category_id != 0},
	).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetProductByID(ctx context.Context, id int) (ProductDetail, error) {

	query := `
	SELECT p.id, p.name, c.name as category, SUM(pv.stock_quantity) as stock, SUM(pv.sold_quantity) as sold
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
		Scan(&p.ID, &p.Name, &p.Category, &p.TotalStock, &p.TotalSold)
	if err != nil {
		fmt.Println(err)
		return ProductDetail{}, nil
	}
	fmt.Println(p)
	return p, nil
}

func (r *repository) GetProductVariantsByID(ctx context.Context, id int) ([]ProductVariant, error) {
	query := `
	SELECT pv.id, colors.name as color, sizes.name as size, pp.price , pv.stock_quantity, pv.sold_quantity
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

	JOIN product_prices as pp
	ON pv.id = pp.product_variant_id

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
