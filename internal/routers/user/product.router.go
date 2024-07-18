package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	var groupName = "/product"

	productRouterPublic := Router.Group(groupName)

	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
}
