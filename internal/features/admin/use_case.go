package admin

import (
	"context"
	"fmt"
)

type UserCase interface {
	GetProducts(ctx context.Context, reg PaginationRequest) (PaginationResponse, error)
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

func (s *userCase) GetProducts(ctx context.Context, req PaginationRequest) (PaginationResponse, error) {
	offset := (req.Page - 1) * req.Limit

	products, err := s.repo.GetProducts(ctx, req.Limit, offset)
	if err != nil {
		return PaginationResponse{}, fmt.Errorf("failed to get products: %w", err)
	}

	totals, err := s.repo.CountProducts(ctx)
	if err != nil {
		return PaginationResponse{}, fmt.Errorf("failed to count products: %w", err)
	}

	totalPages := (totals + req.Limit - 1) / req.Limit
	response := PaginationResponse{
		Data: products,
		Pagination: Metadata{
			CurrentPage:  req.Page,
			PageSize:     req.Limit,
			TotalRecords: totals,
			TotalPages:   totalPages,
			HasNext:      req.Page < totalPages,
			HasPrevious:  req.Page > 1,
		},
	}

	return response, nil
}

func (s *userCase) GetProductByID(id int) Product {
	product := s.GetProductByID(id)
	return product
}
