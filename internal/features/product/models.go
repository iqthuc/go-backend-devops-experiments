package product

const (
	perPageDefault = 2
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Sold  int    `json:"sold"`
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

type Filters struct {
	Keyword    string `json:"key_word"`
	CategoryID int    `json:"category_id"`
}

type SortBy struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

type Pagination struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
}
