package initialize

import (
	"gproject/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	return routers.NewRouter()
}