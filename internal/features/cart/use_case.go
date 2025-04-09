package cart

import "context"

type UseCase interface {
	GetCart(ctx context.Context, userID int64) (*Cart, error)
	AddToCart(ctx context.Context, userID int64, cartItem *CartItem) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

func (uc *useCase) GetCart(ctx context.Context, userID int64) (*Cart, error) {
	return uc.repo.GetCart(ctx, userID)
}

func (uc *useCase) AddToCart(ctx context.Context, userID int64, cartItem *CartItem) error {
	return uc.repo.AddToCart(ctx, userID, cartItem)
}
