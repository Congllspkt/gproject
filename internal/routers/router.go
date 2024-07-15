package routers

import (
	"gproject/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("v1/2024")
	{
		v1.GET("/ping", controllers.NewPongController().Pong) //  curl http://localhost:8082/v1/2024/ping
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