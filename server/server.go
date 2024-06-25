package server
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

func Init() {
	r := gin.Default()
	limiter := ratelimit.New(1)
	r.GET("/ping", rateLimitMiddleware(limiter), getContext)
	r.Run()
}

func rateLimitMiddleware(limiter ratelimit.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter.Take()
		c.Next()
	}
}

func getContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}