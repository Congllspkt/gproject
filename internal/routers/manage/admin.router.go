package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	var groupName = "/admin"
	adminRouterPublic := Router.Group(groupName)

	{
		adminRouterPublic.POST("/login")
	}

	adminRouterPrivate := Router.Group(groupName)
	{
		adminRouterPrivate.POST("/active_user")
	}
}
