// Handles API communication with GitHub to fetch user activity
package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Event represents a GitHub user activity event
type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

// RateLimitInfo holds GitHub API rate limit details
type RateLimitInfo struct {
	Limit     int
	Remaining int
	Reset     int64
}

// FetchUserActivity retrieves recent activity for a given GitHub username
func FetchUserActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the rate limit has been exceeded
	rateLimit := parseRateLimitHeaders(resp)
	if resp.StatusCode == 403 && rateLimit.Remaining == 0 {
		return nil, fmt.Errorf(
			"rate limit exceeded. Limit: %d, Reset at: %d",
			rateLimit.Limit, rateLimit.Reset,
		)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("GitHub user '%s' not found", username)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if rateLimit.Remaining < 5 {
		fmt.Printf("Warning: Only %d API requests remaining. Reset at %d.\n", rateLimit.Remaining, rateLimit.Reset)
	}

	return events, nil
}

// parseRateLimitHeaders extracts rate limit details from HTTP headers
func parseRateLimitHeaders(resp *http.Response) RateLimitInfo {
	limit, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Limit"))
	remaining, _ := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
	reset, _ := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)

	return RateLimitInfo{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}
}
