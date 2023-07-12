package main

import (
	"fmt"

	"github.com/Coding-Brownies/shorty/handler"
	"github.com/Coding-Brownies/shorty/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// json response
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Hey Go URL Shortener !🚀",
			"long_url": "https://it.everli.com/it/spesa/coop/biella",
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
