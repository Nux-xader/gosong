package main

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func make_resp(c *gin.Context, resp_data, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": resp_data, "message": msg})
}

func err_handler(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		make_resp(c, "The simple API for getting lyrics of song, write in Go", map[int]int{})
	})

	r.GET("/search", func(c *gin.Context) {
		title := strings.Replace(c.Query("title"), " ", "+", -1)
		doc, err := goquery.NewDocument("https://www.google.com/search?q=lirik+lagu+" + title)
		err_handler(err)

		make_resp(c, doc.Find(".hwc").First().Text(), map[int]int{})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}
