package store_test

import (
	"testing"

	"github.com/Coding-Brownies/shorty/store"
	"github.com/stretchr/testify/assert"
)

var testStoreService = &store.StorageService{}

func init() {
	testStoreService = store.InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService != nil) // assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	store.SaveUrlMapping(shortURL, initialLink)

	// Retrieve initial URL
	retrievedUrl, _ := store.RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
