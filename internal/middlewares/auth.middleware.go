package middlewares

import (
	"fmt"
	"gproject/internal/responses"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {

		fmt.Println("Start Authorization")

		// token := c.GetHeader("Authorization")
		if false {
			responses.ErrorResponse(c, responses.ERR_INVALID_TOKEN, "")
			c.Abort()
			return
		}
		c.Next()
		fmt.Println("End Authorization")

	}
}