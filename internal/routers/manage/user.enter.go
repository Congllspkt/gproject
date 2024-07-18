package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	var groupName = "/admin/user"

	userRouterPrivate := Router.Group(groupName)
	{
		userRouterPrivate.POST("/active_user")
	}
}
