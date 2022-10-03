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
	host = "postgres"
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

	sql := "CREATE TABLE IF NOT EXISTS coin(face VARCHAR(40) PRIMARY KEY, count INTEGER);"
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}

	var counter int

	db.QueryRow("SELECT count(*) FROM coin").Scan(&counter)

	if counter == 0 {
		sql = "INSERT INTO coin(face, count) VALUES ($1, $2), ($3, $4);"
		_, err = db.Exec(sql, "head", 0, "tail", 0)
		if err != nil {
			panic(err)
		}
	}

	backend := router.Group("/backend")
	{
		backend.POST("/increment", func(c *gin.Context) {
			var coin Coin
			var value int
			if c.BindJSON(&coin) == nil {
				db.Exec("UPDATE coin SET count = count + 1 WHERE face = $1", coin.Face)
				db.QueryRow("SELECT count FROM COIN WHERE face = $1", coin.Face).Scan(&value)
				c.JSON(http.StatusOK, gin.H{
					"value": value,
				})
			}
		})

		backend.GET("/reset", func(c *gin.Context) {
			db.Exec("UPDATE coin SET count = 0")
		})

		backend.GET("/load", func(c *gin.Context) {
			var headNum int
			var tailNum int

			db.QueryRow("SELECT count FROM coin WHERE face = $1", "head").Scan(&headNum)
			db.QueryRow("SELECT count FROM coin WHERE face = $1", "tail").Scan(&tailNum)

			c.JSON(http.StatusOK, gin.H{
				"headNum": headNum,
				"tailNum": tailNum,
			})
		})
	}
	router.Run(":8080")
}
