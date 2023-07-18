package main

import (
	"fmt"
	"net/http"

	"github.com/Coding-Brownies/shorty/handler"
	"github.com/Coding-Brownies/shorty/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// json response
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Shorty",
		})
	})

	// In more complex apps, endpoints should live in a separate file, but for the sake of simplicity and since they are just 2 endpoints, we will be having them in the main.go file
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
