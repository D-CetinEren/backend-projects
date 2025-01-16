package cache

import (
	"encoding/json"
	"os"
	"time"
)

type CacheItem struct {
	Data      json.RawMessage
	Timestamp time.Time
}

// LoadCache loads the cache from a file.
func LoadCache() (map[string]CacheItem, error) {
	file, err := os.Open("cache.json")
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]CacheItem), nil
		}
		return nil, err
	}
	defer file.Close()

	var cacheData map[string]CacheItem
	if err := json.NewDecoder(file).Decode(&cacheData); err != nil {
		return nil, err
	}
	return cacheData, nil
}

// SaveCache saves the cache to a file.
func SaveCache(cacheData map[string]CacheItem) error {
	file, err := os.Create("cache.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(cacheData)
}

// IsValid checks if a cache item is still valid based on the TTL.
func IsValid(item CacheItem, ttl time.Duration) bool {
	return time.Since(item.Timestamp) <= ttl
}
