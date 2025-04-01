package dataloader

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/iqthuc/sport-shop/internal/delivery/graph/model"
	"github.com/iqthuc/sport-shop/pkg/constants"
	"github.com/iqthuc/sport-shop/utils"
)

type categoryReader struct {
	db *sql.DB
}

func (cr categoryReader) getCategories(ctx context.Context, categoryIDs []int64) ([]*model.Category, []error) {
	// Tạo placeholders ($1, $2, $3, ...)
	placeholders := utils.BuildPostgresPlaceholders(len(categoryIDs))

	query := fmt.Sprintf("SELECT id, name FROM categories WHERE id IN (%s)", placeholders)
	log.Printf("Gọi tới categories: %s", query)

	// Chuyển đổi categoryIDs thành dạng []any để truyền vào query
	args := make([]any, len(categoryIDs))
	for i, v := range categoryIDs {
		args[i] = v
	}

	rows, err := cr.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Failed to query categories table: %v", err)
		return nil, []error{constants.Error.SomethingWentWrong}
	}
	defer rows.Close()

	categories := make([]*model.Category, 0, len(categoryIDs))
	var errs []error
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Printf("Error scanning brand data: %v", err)
			errs = append(errs, err)
			continue
		}
		categories = append(categories, &category)
	}
	return categories, errs
}

func PrimeCategory(ctx context.Context, brand *model.Category) bool {
	loaders := For(ctx)
	return loaders.CategoryLoader.Prime(brand.ID, brand)
}

func GetCategory(ctx context.Context, brandID int64) (*model.Category, error) {
	loaders := For(ctx)
	return loaders.CategoryLoader.Load(ctx, brandID)
}

func GetCategories(ctx context.Context, brandIDs []int64) ([]*model.Category, error) {
	loaders := For(ctx)
	return loaders.CategoryLoader.LoadAll(ctx, brandIDs)
}
