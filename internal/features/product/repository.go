package product

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type GetProductsQueryPrarams struct {
	limit      int
	offset     int
	searchKey  string
	categoryId int
}

type Repository interface {
	GetProducts(ctx context.Context, params GetProductsQueryPrarams) (GetProductsResult, error)
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

func (r *repository) GetProducts(ctx context.Context, params GetProductsQueryPrarams) (GetProductsResult, error) {
	query := `
	SELECT id, name, COUNT(*) OVER() AS total_count
	FROM products
	WHERE ($3::int IS NULL OR category_id = $3)
	AND ($4::varchar IS NULL OR name ILIKE '%' || $4 || '%')
	LIMIT $1
	OFFSET $2
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		params.limit,
		params.offset,
		sql.NullInt64{Int64: int64(params.categoryId), Valid: params.categoryId != 0},
		sql.NullString{String: params.searchKey, Valid: params.searchKey != ""},
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
