package middleware

import (
	"go-portfolio/server/api/response"
	"go-portfolio/server/lib/environment"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func newIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}
func RateLimitMiddleware(cfg *environment.Config) gin.HandlerFunc {
	limiter := newIPRateLimiter(rate.Limit(cfg.RATE_LIMITER_RPS), cfg.RATE_LIMITER_BURST)
	log.Print(limiter)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter.mu.Lock()
		if _, found := limiter.ips[ip]; !found {
			limiter.ips[ip] = rate.NewLimiter(limiter.r, limiter.b)
		}
		userLimiter := limiter.ips[ip]
		limiter.mu.Unlock()

		if !userLimiter.Allow() {
			res := response.ErrorResponse(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, res)
			return
		}
		c.Next()
	}
}
