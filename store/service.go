package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	// ctx          = context.Background()
)

// Note that in a real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the
// values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMS whenever the cache is full

const CacheDuration = 6 * time.Hour

// Initializing the store service (in this case our Redis client) and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/

func SaveUrlMapping(shortUrl string, originalUrl string) error {
	return storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/

func RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		return "", err
	}
	return result, nil

}
