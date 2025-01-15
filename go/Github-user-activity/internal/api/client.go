package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/D-CetinEren/backend-projects/go/Github-user-activity/internal/models"
)

// FetchUserActivity fetches recent activity for a GitHub user with pagination.
func FetchUserActivity(username string, maxPages int) ([]models.Event, error) {
	var allEvents []models.Event

	for page := 1; page <= maxPages; page++ {
		url := fmt.Sprintf("https://api.github.com/users/%s/events?page=%d", username, page)
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

		// Stop fetching if there are no more events
		if len(events) == 0 {
			break
		}

		allEvents = append(allEvents, events...)
	}

	return allEvents, nil
}
