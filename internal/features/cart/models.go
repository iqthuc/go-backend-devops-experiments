package cart

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ProductID int64   `json:"id" bson:"product_id"`
	Name      string  `json:"name" bson:"name"`
	Price     float64 `json:"price" bson:"price"`
	Quantity  int     `json:"quantity" bson:"quantity"`
}

type Cart struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    int64              `json:"user_id" bson:"user_id"`
	Items     []CartItem         `json:"items" bson:"items"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type AddToCartRequestParams struct {
	UserID    int64   `json:"user_id"`
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
