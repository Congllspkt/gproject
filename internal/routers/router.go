package routers

import (
	"fmt"
	"gproject/internal/initialize/global"
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

	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	manageRouter := RouterGroupApp.Manage
	userRouter := RouterGroupApp.User

	mainGroup := r.Group("/v1/2024")
	{
		mainGroup.GET("checkStatus")
	}

	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}

	{
		manageRouter.InitUserRouter(mainGroup)
		manageRouter.InitAdminRouter(mainGroup)
	}

	return r

}
