package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type Coin struct {
	Face string `form:"coin" json:"coin" binding:"required"`
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	pw   = "postgres_password"
	db   = "postgres"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{":3000"})

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, pw, db)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	backend := router.Group("/backend")
	{
		backend.POST("/increment", func(c *gin.Context) {
			var coin Coin
			if c.BindJSON(&coin) == nil {
				c.JSON(http.StatusOK, gin.H{
					"message": coin.Face,
				})
			}
		})
	}
	router.Run(":8080")
}
