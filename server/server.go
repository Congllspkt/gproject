package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

func TestHttp() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello Cong")
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye Cong")
	})

	http.ListenAndServe(":4554", nil)
}

func InitRateLimit() {
	limiter := ratelimit.New(100, ratelimit.Per(time.Minute))

	r := gin.Default()
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