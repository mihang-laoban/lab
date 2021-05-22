package test

import (
	"github.com/gin-gonic/gin"
	"testing"
	"dev-go-base/server"
)

func TestServer(t *testing.T){
	r := server.NewServer()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
