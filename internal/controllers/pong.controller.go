package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (uc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", " .") // curl http://localhost:8082/v2/2024/ping/name
	uid := c.Query("uid")                // curl http://localhost:8082/v2/2024/ping/?uid=1234
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping kkk  " + name,
		"uid":     uid,
		"users":   []string{"xxx", "kkk"},
	})
}