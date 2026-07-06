package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware is a higher-order function (Functional Programming pattern).
// It takes our dependencies and returns a Gin Handler function (a closure) that Gin will execute.
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Record the start time
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 2. Pass control to the next middleware or controller (like AddBook)
		c.Next()

		// 3. After the controller finishes, record the end time and log the result
		latency := time.Since(start)

		logger.Info("HTTP Request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
