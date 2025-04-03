package cart

import "context"

type UserCase interface {
	GetCart(ctx context.Context, userID int64) (*Cart, error)
	AddToCart(ctx context.Context, userID int64, cartItem *CartItem) error
}

type userCase struct {
	repo Repository
}

func NewUserCase(repo Repository) UserCase {
	return &userCase{repo: repo}
}

func (uc *userCase) GetCart(ctx context.Context, userID int64) (*Cart, error) {
	return uc.repo.GetCart(ctx, userID)
}

func (uc *userCase) AddToCart(ctx context.Context, userID int64, cartItem *CartItem) error {
	return uc.repo.AddToCart(ctx, userID, cartItem)
}
