package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type LimiterStore struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

func NewLimiterStore() *LimiterStore {
	return &LimiterStore{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (s *LimiterStore) GetLimiter(clientIP string) *rate.Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limiter, exists := s.limiters[clientIP]; exists {
		return limiter
	}

	// 	Rate (r = 5): Bộ giới hạn sinh 5 token mỗi giây.
	// Burst (burst = 10): Xô chứa được tối đa 10 token.
	limiter := rate.NewLimiter(5, 10)
	s.limiters[clientIP] = limiter
	return limiter
}

func RateLimiterMiddleware(store *LimiterStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		limiter := store.GetLimiter(clientIP)
		if !limiter.Allow() {
			// Quá giới hạn, trả về lỗi
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
