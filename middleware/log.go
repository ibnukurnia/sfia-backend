package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		zap.L().Info("request", zap.String("path", c.Request.URL.Path), zap.String("method", c.Request.Method))

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error("Error", zap.String("error", e))
			}
		}
	}
}
