package product

import (
	"encoding/json"
)

const (
	perPageDefault = 5
)

type Product struct {
	ID        int32   `json:"id"`
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	Stock     int     `json:"stock"`
	Sold      int     `json:"sold"`
}

// TotalCount is necessary because the returned products are limited by pagination.
type GetProductsResult struct {
	products   []Product
	totalCount int
}

type ProductsResponse struct {
	Products   []Product  `json:"products"`
	Filters    Filters    `json:"filters"`
	SortBy     SortBy     `json:"sort_by"`
	Pagination Pagination `json:"pagination"`
}

func (pr *ProductsResponse) MarshalJSON() ([]byte, error) {
	type Alias ProductsResponse
	aux := &struct {
		Filters *Filters `json:"filters,omitempty"`
		SortBy  *SortBy  `json:"sort_by,omitempty"`
		*Alias
	}{
		Filters: &pr.Filters,
		SortBy:  &pr.SortBy,
		Alias:   (*Alias)(pr),
	}
	if pr.Filters == (Filters{}) {
		aux.Filters = nil
	}
	if pr.SortBy == (SortBy{}) {
		aux.SortBy = nil
	}

	return json.Marshal(aux)
}

type Filters struct {
	Keyword    string  `json:"key_word,omitempty"`
	CategoryID int     `json:"category_id,omitempty"`
	PriceMin   float64 `json:"price_min,omitempty"`
	PriceMax   float64 `json:"price_max,omitempty"`
}

type SortBy struct {
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}

type Pagination struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
}

// // Product details
type ProductDetailsResponse struct {
	ProductDetail   ProductDetail    `json:"product_details"`
	ProductVariants []ProductVariant `json:"product_variant"`
}

type ProductDetail struct {
	Product
	Category string `json:"category"`
}

type ProductVariant struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Color *string `json:"color"`
	Size  *string `json:"size"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
	Sold  int     `json:"sold"`
}
