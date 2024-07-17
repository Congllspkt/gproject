package routers

import (
	"fmt"
	"gproject/internal/controllers"
	"gproject/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Start ----> AA")
		c.Next()
		fmt.Println("End ----> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Start ----> BB")
		c.Next()
		fmt.Println("End ----> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Start ----> CC")
	c.Next()
	fmt.Println("End ----> CC")
}

func NewRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC)

	v1 := r.Group("v1/2024")
	{
		v1.GET("/ping", controllers.NewPongController().Pong)        //  curl http://localhost:8082/v1/2024/ping
		v1.GET("/user", controllers.NewUserController().GetuserById) //  curl http://localhost:8082/v1/2024/ping
		// v1.PUT("/ping", Pong)
	}

	// v2 := r.Group("v2/2024")
	// {
	// curl http://localhost:8082/v2/2024/ping
	// v2.GET("/ping/:name", Pong)
	// v2.GET("/ping", Pong)
	// v2.PUT("/ping", Pong)
	// }

	return r

}
