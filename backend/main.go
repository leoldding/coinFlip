package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Coin string `form:"coin" json:"coin" binding:"required"`
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{":3000"})

	backend := router.Group("/backend")
	{
		backend.POST("/increment", func(c *gin.Context) {
			var res Result
			if c.BindJSON(&res) == nil {
				c.JSON(http.StatusOK, gin.H{
					"message": res.Coin,
				})
			}
		})
	}
	router.Run(":8080")
}
