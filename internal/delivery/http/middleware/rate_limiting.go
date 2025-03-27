package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/iqthuc/sport-shop/utils"
)

type rateLimitInfo struct {
	requests  int
	resetTime time.Time // thời điểm reset bộ đếm
}

type RateLimiter struct {
	clients sync.Map
	limit   int
	windows time.Duration // khoảng thời gian giới hạn (vd: 15p)
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		windows: window,
	}
}

func (rl *RateLimiter) RateLimiting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		value, loaded := rl.clients.LoadOrStore(clientIP, &rateLimitInfo{
			requests:  0,
			resetTime: time.Now().Add(rl.windows),
		})
		if !loaded {
			log.Printf("new client access: %s", clientIP)
		}
		clientData := value.(*rateLimitInfo)

		if time.Now().After(clientData.resetTime) {
			clientData.requests = 0
			clientData.resetTime = time.Now().Add(rl.windows)
		}
		if clientData.requests >= rl.limit {
			utils.ErrorJsonResponse(w, http.StatusTooManyRequests, "Rap chậm thôi")
			return
		}
		clientData.requests++
		next.ServeHTTP(w, r)
	})
}
