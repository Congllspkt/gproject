package server


import (
	"net/http"
	"github.com/gin-gonic/gin"
  )

func Init() {
	r := gin.Default()
	r.GET("/ping", getContext)
	r.Run()
}


func getContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	  "message": "pong",
	})
  }
