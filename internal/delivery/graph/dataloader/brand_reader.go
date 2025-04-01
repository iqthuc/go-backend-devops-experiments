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

type brandReader struct {
	db *sql.DB
}

func (br brandReader) getBrands(ctx context.Context, brandIDs []int64) ([]*model.Brand, []error) {
	if len(brandIDs) == 0 {
		return nil, nil
	}
	// Tạo placeholders ($1, $2, $3, ...)
	placeholders := utils.BuildPostgresPlaceholders(len(brandIDs))

	query := fmt.Sprintf("SELECT id, name FROM brands WHERE id IN (%s)", placeholders)
	log.Printf("Gọi tới brands: %s", query)

	// Chuyển đổi brandIDs thành dạng []any để truyền vào query
	args := make([]any, len(brandIDs))
	for i, v := range brandIDs {
		args[i] = v
	}

	rows, err := br.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Failed to query brands table: %v", err)
		return nil, []error{constants.Error.SomethingWentWrong}
	}
	defer rows.Close()

	brands := make([]*model.Brand, 0, len(brandIDs))
	var errs []error
	for rows.Next() {
		var brand model.Brand
		if err := rows.Scan(&brand.ID, &brand.Name); err != nil {
			log.Printf("Error scanning brand data: %v", err)
			errs = append(errs, err)
			continue
		}
		brands = append(brands, &brand)
	}

	return brands, errs
}

func PrimeBrand(ctx context.Context, brand *model.Brand) bool {
	loaders := For(ctx)
	return loaders.BrandLoader.Prime(brand.ID, brand)
}

func GetBrand(ctx context.Context, brandID int64) (*model.Brand, error) {
	loaders := For(ctx)
	return loaders.BrandLoader.Load(ctx, brandID)
}

func GetBrands(ctx context.Context, brandIDs []int64) ([]*model.Brand, error) {
	loaders := For(ctx)
	return loaders.BrandLoader.LoadAll(ctx, brandIDs)
}
