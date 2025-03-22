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
	GetProductByID(id int) (Product, error)
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
	SELECT id, name, COUNT(*) OVER() AS total_count
	FROM products
	WHERE ($3::int IS NULL OR category_id = $3)
		AND ($4::varchar IS NULL OR name_search @@ plainto_tsquery($4))
		AND ($5::float IS NULL OR base_price >= $5)
		AND ($6::float IS NULL OR base_price <= $6)
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
		if err := rows.Scan(&product.ID, &product.Name, &count); err != nil {
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

func (r *repository) GetProductByID(id int) (Product, error) {
	return Product{}, nil
}
