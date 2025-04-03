package database

import (
	"context"
	"testing"

	"github.com/iqthuc/sport-shop/utils"
	_ "github.com/lib/pq"
)

func TestNewPostgres(t *testing.T) {
	utils.InitTestConfig()
	db, err := NewPostgres()
	defer db.Close()
	if err != nil {
		t.Errorf("error connecting to postgres: %v", err)
	}

}

func TestNewMongoDB(t *testing.T) {
	utils.InitTestConfig()
	client, err := NewMongoDB()
	defer client.Disconnect(context.Background())
	if err != nil {
		t.Errorf("error connecting to mongodb: %v", err)
	}

}
