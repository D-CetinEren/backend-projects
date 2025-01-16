package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/cache"
	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/models"
)

// FetchUserActivityWithCache fetches GitHub user activity with caching logic
func FetchUserActivityWithCache(username string, maxPages int, ttl time.Duration) ([]models.Event, error) {
	// Load the cache
	cacheData, err := cache.LoadCache()
	if err != nil {
		return nil, fmt.Errorf("failed to load cache: %v", err)
	}

	// Check if the user data is in the cache and valid
	if item, found := cacheData[username]; found && cache.IsValid(item, ttl) {
		var events []models.Event
		if err := json.Unmarshal(item.Data, &events); err == nil {
			fmt.Println("Loaded data from cache.")
			return events, nil
		}
	}

	// Fetch data from GitHub API if not found in cache or cache expired
	events, err := FetchUserActivity(username, maxPages) // Assuming this function exists to fetch events from the API
	if err != nil {
		return nil, err
	}

	// Save fetched data to the cache
	cacheData[username] = cache.CacheItem{
		Data:      toJSON(events),
		Timestamp: time.Now(),
	}
	if err := cache.SaveCache(cacheData); err != nil {
		fmt.Printf("Warning: failed to save cache: %v\n", err)
	}

	return events, nil
}

// Helper function to convert data to JSON
func toJSON(data interface{}) []byte {
	jsonData, _ := json.Marshal(data)
	return jsonData
}

// FetchUserActivity fetches user activity from GitHub API
func FetchUserActivity(username string, maxPages int) ([]models.Event, error) {
	if maxPages <= 0 {
		return nil, errors.New("maxPages must be greater than 0")
	}

	baseURL := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	var allEvents []models.Event

	for page := 1; page <= maxPages; page++ {
		// Construct URL with pagination
		url := fmt.Sprintf("%s?page=%d", baseURL, page)

		// Make HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch events from GitHub API: %v", err)
		}
		defer resp.Body.Close()

		// Check for HTTP response status
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("GitHub API returned status: %s", resp.Status)
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		// Unmarshal JSON data into Event slice
		var events []models.Event
		if err := json.Unmarshal(body, &events); err != nil {
			return nil, fmt.Errorf("failed to parse response: %v", err)
		}

		// Add events to allEvents
		allEvents = append(allEvents, events...)

		// If less than a full page of events is returned, stop pagination
		if len(events) < 30 {
			break
		}
	}

	return allEvents, nil
}
