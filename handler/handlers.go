package handler

import (
	"net/http"

	"github.com/Coding-Brownies/shorty/shortener"
	"github.com/Coding-Brownies/shorty/store"
	"github.com/gin-gonic/gin"
)

// Data models: type request struct {} and type response struct {}
// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {

	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	err := store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error creating the key"})
		return
	}

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {

	shortUrl := c.Param("shortUrl")
	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "url not found"})
	}
	c.Redirect(302, initialUrl)
}
