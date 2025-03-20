package admin

import "database/sql"

// Products
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Sold  int    `json:"sold"`
}

type PaginationRequest struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	CategoryId int    `json:"category_id"`
	SearchKey  string `json:"search_key"`
}
type Filter struct {
	CategoryId int    `json:"category_id"`
	SearchKey  string `json:"search_key"`
}
type PaginationResponse struct {
	Data       []Product `json:"data"`
	Pagination Metadata  `json:"pagination"`
	Filter     Filter    `json:"filter"`
}

type Metadata struct {
	CurrentPage  int  `json:"current_page"`
	PageSize     int  `json:"page_size"`
	TotalRecords int  `json:"total_records"`
	TotalPages   int  `json:"total_pages"`
	HasNext      bool `json:"has_next"`
	HasPrevious  bool `json:"has_previous"`
}

//-----------------//

// Product details
type ProductDetailsResponse struct {
	ProductDetail   ProductDetail    `json:"product_details"`
	ProductVariants []ProductVariant `json:"product_variant"`
}

type ProductDetail struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	TotalStock int    `json:"total_stock"`
	TotalSold  int    `json:"total_sold"`
}

type ProductVariant struct {
	ID    int64          `json:"id"`
	Name  string         `json:"name"`
	Color sql.NullString `json:"color"`
	Size  sql.NullString `json:"size"`
	Price float64        `json:"price"`
	Stock int            `json:"stock"`
	Sold  int            `json:"sold"`
}

type ProductCategory int

const (
	Racket ProductCategory = 1
	Shoes  ProductCategory = 2
	Shirt  ProductCategory = 3
	Pants  ProductCategory = 4
)
