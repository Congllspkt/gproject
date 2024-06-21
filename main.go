package main

// import "gproject/baitap"
import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	r.Run()
}


/*
git add .
git commit -m 'init first 2'
git push


*/