package product

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository interface {
	GetProducts(ctx context.Context) ([]Product, error)
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

func (r *repository) GetProducts(ctx context.Context) ([]Product, error) {
	query := `SELECT id, name FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		product := Product{}
		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	return products, nil
}

func (r *repository) GetProductByID(id int) (Product, error) {
	return Product{}, nil
}
