package middleware

import (
	"go-portfolio/server/lib/environment"
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(cfg *environment.Config) gin.HandlerFunc {
	log.Printf("Allowed origin is %v", cfg.ALLOWED_ORIGINS)
	origins := strings.Split(cfg.ALLOWED_ORIGINS, ",")
	var allowedOrigins []string

	for _, o := range origins {
		cleanOrigin := strings.TrimSpace(o)
		if cleanOrigin != "" {
			allowedOrigins = append(allowedOrigins, cleanOrigin)
		}
	}

	log.Printf("Final Allowed Origins: %v", allowedOrigins)
	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Admin-Secret"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
