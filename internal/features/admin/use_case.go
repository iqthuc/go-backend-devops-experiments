package admin

import (
	"context"
	"fmt"
)

type UseCase interface {
	GetProducts(ctx context.Context, reg PaginationRequest) (PaginationResponse, error)
	GetProductDetails(ctx context.Context, id int) (ProductDetailsResponse, error)
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (s *useCase) GetProducts(ctx context.Context, req PaginationRequest) (PaginationResponse, error) {
	offset := (req.Page - 1) * req.Limit

	products, err := s.repo.GetProducts(ctx, req.Limit, offset, req.CategoryId, req.SearchKey)
	if err != nil {
		return PaginationResponse{}, fmt.Errorf("failed to get products: %w", err)
	}

	totals, err := s.repo.CountProducts(ctx, req.CategoryId)
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
		Filter: Filter{CategoryId: req.CategoryId,
			SearchKey: req.SearchKey,
		},
	}

	return response, nil
}

func (s *useCase) GetProductDetails(ctx context.Context, id int) (ProductDetailsResponse, error) {
	var data ProductDetailsResponse
	productDetail, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return ProductDetailsResponse{}, fmt.Errorf("failed to get product by id: %w", err)
	}
	data.ProductDetail = productDetail

	productVariants, err := s.repo.GetProductVariantsByID(ctx, id)
	if err != nil {
		return ProductDetailsResponse{}, fmt.Errorf("failed to get product variants by id: %w", err)
	}
	data.ProductVariants = productVariants
	return data, err
}
