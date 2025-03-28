package product

import (
	"context"
	"fmt"
	"strings"
)

type GetProductsRequestParams struct {
	page    int
	Filters Filters
	SortBy  SortBy
}

type UserCase interface {
	GetProducts(ctx context.Context, requestParams GetProductsRequestParams) (ProductsResponse, error)
	GetProductDetails(ctx context.Context, id int) (ProductDetailsResponse, error)
}

type userCase struct {
	repo Repository
}

func NewService(repo Repository) UserCase {
	return &userCase{
		repo: repo,
	}
}

func (s *userCase) GetProducts(ctx context.Context, requestParams GetProductsRequestParams) (ProductsResponse, error) {
	//Input validation and transformation
	perPage := perPageDefault
	offset := max((requestParams.page-1), 0) * perPage
	validSortFields := map[string]bool{
		"id":         true,
		"name":       true,
		"base_price": true,
	}
	if !validSortFields[requestParams.SortBy.Field] {
		requestParams.SortBy.Field = ""
	}
	if strings.ToUpper(requestParams.SortBy.Order) != "DESC" {
		requestParams.SortBy.Order = ""
	}

	queryParams := GetProductsQueryParams{
		offset:  offset,
		limit:   perPage,
		Filters: requestParams.Filters,
		SortBy:  requestParams.SortBy,
	}
	result, err := s.repo.GetProducts(ctx, queryParams)
	if err != nil {
		return ProductsResponse{}, fmt.Errorf("Failed to get products: %w", err)
	}

	// handle pagination
	currentpage := max(requestParams.page, 1)
	total := result.totalCount
	totalPages := (result.totalCount + perPage - 1) / perPage
	prevPage := max(currentpage-1, 0)
	nextPage := min(currentpage+1, totalPages)
	pagination := Pagination{
		CurrentPage: currentpage,
		Total:       total,
		TotalPages:  totalPages,
		PerPage:     perPage,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}

	response := ProductsResponse{
		Products:   result.products,
		Filters:    requestParams.Filters,
		SortBy:     requestParams.SortBy,
		Pagination: pagination,
	}

	return response, nil
}

func (s *userCase) GetProductDetails(ctx context.Context, id int) (ProductDetailsResponse, error) {
	var data ProductDetailsResponse
	productDetail, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return ProductDetailsResponse{}, fmt.Errorf("Failed to get product by id: %w", err)
	}

	productVariants, err := s.repo.GetProductVariantsByID(ctx, id)
	if err != nil {
		return ProductDetailsResponse{}, fmt.Errorf("Failed to get product variants by id: %w", err)
	}
	data.ProductDetail = productDetail
	data.ProductVariants = productVariants
	return data, err
}
