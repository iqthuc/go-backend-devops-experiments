package resolver

import (
	"context"

	"github.com/iqthuc/sport-shop/internal/delivery/graph/dataloader"
	"github.com/iqthuc/sport-shop/internal/delivery/graph/model"
)

// Category is the resolver for the category field.
func (r *productResolver) Category(ctx context.Context, obj *model.Product) (*model.Category, error) {
	return dataloader.GetCategory(ctx, obj.Category_id)
}
