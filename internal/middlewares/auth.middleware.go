package middlewares

import (
	"gproject/internal/responses"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid_token" {
			responses.ErrorResponse(c, responses.ERR_INVALID_TOKEN, "")
			c.Abort()
			return
		}
		c.Next()
	}
}