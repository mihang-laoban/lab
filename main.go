package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/this", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "that",
		})
	})
	r.GET("/", func(c *gin.Context) {
		count, err := c.Writer.Write([]byte("Hello world"))
		log.Println(count, err)
	})
	if err := r.Run(":5000"); err != nil {
		log.Println(err)
	}
}