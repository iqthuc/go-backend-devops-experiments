package resolver

import (
	"context"

	"github.com/iqthuc/sport-shop/internal/delivery/graph/dataloader"
	"github.com/iqthuc/sport-shop/internal/delivery/graph/model"
)

// Brand is the resolver for the brand field.
func (r *productResolver) Brand(ctx context.Context, obj *model.Product) (*model.Brand, error) {
	return dataloader.GetBrand(ctx, obj.Brand_id)
}
