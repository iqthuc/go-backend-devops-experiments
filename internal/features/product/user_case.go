package product

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
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
	repo  Repository
	cache *redis.Client
}

func NewUserCase(repo Repository, redis *redis.Client) UserCase {
	return &userCase{
		repo:  repo,
		cache: redis,
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
	cacheKey := fmt.Sprintf("products:page=%d:filters=%v:sort=%s-%s",
		requestParams.page, requestParams.Filters, requestParams.SortBy.Field, requestParams.SortBy.Order)
	cachedData, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("lấy products từ cache")
		var cachedResponse ProductsResponse
		json.Unmarshal([]byte(cachedData), &cachedResponse)
		return cachedResponse, nil
	}
	log.Println("không có products trong cache, tự lấy products")

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

	cacheData, _ := json.Marshal(response)
	s.cache.Set(ctx, cacheKey, cacheData, time.Minute*10)

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
