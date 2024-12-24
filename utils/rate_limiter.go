package utils

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	timestamp map[string]time.Time
	limit     time.Duration
}

func NewRateLimiter(limit time.Duration) *RateLimiter {
	return &RateLimiter{
		timestamp: make(map[string]time.Time),
		limit:     limit,
	}
}

func (r *RateLimiter) Allow(ip string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if lastRequest, found := r.timestamp[ip]; found && time.Since(lastRequest) < r.limit {
		return false
	}
	r.timestamp[ip] = time.Now()
	return true
}

func Middleware(limiter *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if !limiter.Allow(ip) {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
