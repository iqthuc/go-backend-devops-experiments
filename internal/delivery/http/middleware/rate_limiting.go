package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iqthuc/sport-shop/utils"
	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	redisClient *redis.Client
	limit       int
	windows     time.Duration // khoảng thời gian giới hạn (vd: 15p)
}

func NewRateLimiter(redisClient *redis.Client, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		redisClient: redisClient,
		limit:       limit,
		windows:     window,
	}
}

// tạo cacheKey cho mỗi clientIP, tăng count mỗi khi có request
// nếu là lần request đầu, đặt thời gian giới hạn
// nếu vượt quá limit, báo lỗi
func (rl *RateLimiter) RateLimiting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		clientIP := r.RemoteAddr
		cacheKey := fmt.Sprintf("rate_limit:%s", clientIP)

		count, err := rl.redisClient.Incr(ctx, cacheKey).Result()
		if err != nil {
			log.Printf("Redis error: %v", err)
			utils.ErrorJsonResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		if count == 1 {
			rl.redisClient.Expire(ctx, cacheKey, rl.windows)
		}

		if count > int64(rl.limit) {
			utils.ErrorJsonResponse(w, http.StatusTooManyRequests, "Rap chậm thôi")
			return
		}

		next.ServeHTTP(w, r)
	})
}
