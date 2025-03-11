package product

import (
	"context"
	"fmt"
)

type UserCase interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProductByID(id int) Product
}

type userCase struct {
	repo Repository
}

func NewService(repo Repository) UserCase {
	return &userCase{
		repo: repo,
	}
}

func (s *userCase) GetProducts(ctx context.Context) ([]Product, error) {
	products, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	return products, nil
}

func (s *userCase) GetProductByID(id int) Product {
	product := s.GetProductByID(id)
	return product
}
