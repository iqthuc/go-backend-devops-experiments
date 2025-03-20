package product

import (
	"context"
	"fmt"
)

type GetProductsRequestParams struct {
	page       int
	searchKey  string
	categoryId int
}

type UserCase interface {
	GetProducts(ctx context.Context, requestParams GetProductsRequestParams) (ProductsResponse, error)
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

func (s *userCase) GetProducts(ctx context.Context, requestParams GetProductsRequestParams) (ProductsResponse, error) {
	perPage := perPageDefault
	queryParams := GetProductsQueryPrarams{
		offset:     (requestParams.page - 1) * perPage,
		limit:      perPage,
		searchKey:  requestParams.searchKey,
		categoryId: requestParams.categoryId,
	}
	result, err := s.repo.GetProducts(ctx, queryParams)
	if err != nil {
		return ProductsResponse{}, fmt.Errorf("failed to get products: %w", err)
	}

	// handle pagination
	currentpage := requestParams.page
	total := result.totalCount
	totalPages := result.totalCount / perPage
	prevPage := max(requestParams.page-1, 0)
	nextPage := min(requestParams.page+1, totalPages)
	pagination := Pagination{
		CurrentPage: currentpage,
		Total:       total,
		TotalPages:  totalPages,
		PerPage:     perPage,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}

	response := ProductsResponse{
		Products: result.products,
		Filters: Filters{
			Keyword: requestParams.searchKey,
		},
		Pagination: pagination,
	}

	return response, nil
}

func (s *userCase) GetProductByID(id int) Product {
	product := s.GetProductByID(id)
	return product
}
