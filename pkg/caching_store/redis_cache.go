package cachingstore

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

// NewCacheService returns a new cache service
func NewCacheService() CartTripCacheService {
	redisListener := os.Getenv("REDIS_LISTENER_PORT")
	return &RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr:     redisListener,
			Password: "",
			DB:       0,
		}),
	}
}

// RedisCache caching service with redis
type RedisCache struct {
	Client *redis.Client
}

func (rc *RedisCache) isCacheActive() (bool, error) {
	_, err := rc.Client.Ping().Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

//Get returns a value mapped to key
func (rc *RedisCache) Get(key string) (int, error) {
	if len(key) == 0 {
		return 0, fmt.Errorf("Get: Key's empty")
	}

	value, valueErr := rc.Client.Get(key).Int()
	if valueErr == redis.Nil {
		log.Println(fmt.Errorf("Get:key (%s) does not exist;", key))
		return 0, nil
	}

	return value, valueErr
}

//Set saves a key,value into redis
func (rc *RedisCache) Set(key string, value int) error {
	if key == "" {
		return fmt.Errorf("Set:Invalid key")
	}

	if value == 0 {
		return fmt.Errorf("Set:Invalid value(%v) for key (%s)", value, key)
	}

	if rc.Client == nil {
		return fmt.Errorf("Set:Invalid client reference")
	}

	setValueErr := rc.Client.Set(key, value, 0).Err()
	if setValueErr != nil {
		return setValueErr
	}

	return nil
}

//Del Deletes a key,value from redis
func (rc *RedisCache) Del(key string) error {
	if key == "" {
		return fmt.Errorf("Set:Invalid key")
	}

	if rc.Client == nil {
		return fmt.Errorf("Set:Invalid client reference")
	}

	setValueErr := rc.Client.Del(key).Err()
	if setValueErr != nil {
		return setValueErr
	}

	return nil
}

//ClearCache clears redis cache
func (rc *RedisCache) ClearCache() error {
	clearErr := rc.Client.FlushAll().Err()
	if clearErr != nil {
		return clearErr
	}
	log.Printf("cache cleared")
	return nil
}
