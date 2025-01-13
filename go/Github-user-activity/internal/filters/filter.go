package filters

import "Github-user-activity/internal/models"

// FilterEventsByType filters events based on the given type.
func FilterEventsByType(events []models.Event, eventType string) []models.Event {
	if eventType == "" {
		return events
	}

	var filtered []models.Event
	for _, event := range events {
		if event.Type == formatEventType(eventType) {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

// formatEventType converts user-friendly type names to GitHub API event types.
func formatEventType(eventType string) string {
	switch eventType {
	case "push":
		return "PushEvent"
	case "issue":
		return "IssuesEvent"
	case "star":
		return "WatchEvent"
	default:
		return eventType
	}
}
