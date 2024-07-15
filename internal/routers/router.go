package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("v1/2024")
	{
		v1.GET("/ping", Pong) //  curl http://localhost:8082/v1/2024/ping
		v1.PUT("/ping", Pong)
	}

	v2 := r.Group("v2/2024")
	{
		// curl http://localhost:8082/v2/2024/ping
		// v2.GET("/ping/:name", Pong)
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
	}

	return r

}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", " .") // curl http://localhost:8082/v2/2024/ping/name
	uid := c.Query("uid")                // curl http://localhost:8082/v2/2024/ping/?uid=1234
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping kkk  " + name,
		"uid":     uid,
		"users":   []string{"xxx", "kkk"},
	})
}
