// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Brand struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Color struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}

type Product struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	Category_id int64             `json:"category,omitempty"`
	Brand_id    int64             `json:"brand,omitempty"`
	BasePrice   float64           `json:"basePrice"`
	Variants    []*ProductVariant `json:"variants"`
}

type ProductVariant struct {
	ID            int64    `json:"id"`
	Product       *Product `json:"product"`
	StockQuantity int      `json:"stockQuantity"`
	SoldQuantity  int      `json:"soldQuantity"`
	Price         float64  `json:"price"`
	Size          *Size    `json:"size,omitempty"`
	Color         *Color   `json:"color,omitempty"`
}

type Query struct {
}

type Size struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
