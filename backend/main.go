package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	backend := router.Group("/backend")
	{
		backend.POST("/increment", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			return
		})
	}
	router.Run(":8080")
}
