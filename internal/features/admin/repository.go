package admin

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository interface {
	GetProducts(ctx context.Context, limit int, offset int) ([]Product, error)
	CountProducts(ctx context.Context) (int, error)
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

func (r *repository) GetProducts(ctx context.Context, limit int, offset int) ([]Product, error) {
	query := `
	SELECT product_variants.id, name, price, stock_quantity
	FROM products
	INNER JOIN product_variants ON products.id = product_variants.product_id
	INNER JOIN product_prices ON product_variants.id = product_prices.product__variant_id
	ORDER BY product_variants.id
	LIMIT $1
	OFFSET $2;
	`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		product := Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	return products, nil
}

func (r *repository) CountProducts(ctx context.Context) (int, error) {
	var total int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM products").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetProductByID(id int) (Product, error) {
	return Product{}, nil
}
