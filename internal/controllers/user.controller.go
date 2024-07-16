package controllers

import (
	"gproject/internal/responses"
	"gproject/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetuserById(c *gin.Context) {
	// responses.SuccessResponse(c, 20001, []string{"aa","bb","cc"})
	responses.ErrorResponse(c, 20003, "no")
}