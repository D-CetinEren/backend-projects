package formatter

import "Github-user-activity/internal/models"

// FormatEvent formats an event into a user-friendly string.
func FormatEvent(event models.Event) string {
	switch event.Type {
	case "PushEvent":
		return "- Pushed to repository '" + event.Repo.Name + "'"
	case "IssuesEvent":
		return "- Opened an issue in '" + event.Repo.Name + "'"
	case "WatchEvent":
		return "- Starred repository '" + event.Repo.Name + "'"
	default:
		return "- " + event.Type + " in '" + event.Repo.Name + "'"
	}
}
