package apiserver

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CORSMiddleware allows everything CORS.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// MetricsMiddleware measures and logs how long a request
// took in milliseconds.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Call next to jump into the next handler.
		c.Next()

		took := time.Since(t).String()

		log.
			With(
				zap.String("path", c.Request.RequestURI),
				zap.String("took", took),
			).
			Debug("request end")
	}
}
