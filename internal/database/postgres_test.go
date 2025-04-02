package database

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestNewPostgres(t *testing.T) {
	db, err := NewPostgres()
	defer db.Close()
	if err != nil {
		t.Errorf("error connecting to postgres: %v", err)
	}

}
