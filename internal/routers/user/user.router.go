package user
import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	var groupName = "/user"
	
	userRouterPublic := Router.Group(groupName) 
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	userRouterPrivate := Router.Group(groupName) 
	{
		userRouterPrivate.GET("/get_info")
	}
}