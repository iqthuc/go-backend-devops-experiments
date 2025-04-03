package cart

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetCart(ctx context.Context, userID int64) (*Cart, error)
	AddToCart(ctx context.Context, userID int64, cartItem *CartItem) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client) Repository {
	collection := client.Database("carts").Collection("carts")
	return &repository{collection: collection}
}

func (r *repository) GetCart(ctx context.Context, userID int64) (*Cart, error) {
	var cart Cart
	filter := bson.M{"user_id": userID}
	err := r.collection.FindOne(ctx, filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get carts: %w", err)

	}
	return &cart, err
}

func (r *repository) AddToCart(ctx context.Context, userID int64, cartItem *CartItem) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$setOnInsert": bson.M{"created_at": time.Now()},
		"$push":        bson.M{"items": cartItem},
	}
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}
