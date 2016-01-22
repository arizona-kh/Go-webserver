package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	webServer()
}

func webServer() {
	r := gin.Default()
	r.LoadHTMLGlob("resources/*.html")
	r.GET("/", codingPage)
	r.GET("/index", codingPage)
	r.Run(":8080")
}

func codingPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "GoBox",
	})
}
