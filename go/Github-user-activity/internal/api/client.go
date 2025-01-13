package api

import (
	"Github-user-activity/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// FetchUserActivity fetches recent activity for a GitHub user.
func FetchUserActivity(username string) ([]models.Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch activity: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error fetching activity: " + resp.Status)
	}

	var events []models.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return events, nil
}
