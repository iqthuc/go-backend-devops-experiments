package cache

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	rdb, err := NewRedisClient()
	defer rdb.Close()
	if err != nil {
		t.Errorf("error connecting to postgres: %v", err)
	}
}
