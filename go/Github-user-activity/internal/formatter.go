package internal

import "fmt"

func FormatEvent(event Event) string {
	switch event.Type {
	case "PushEvent":
		return fmt.Sprintf("- Pushed to repository '%s'", event.Repo.Name)
	case "IssuesEvent":
		return fmt.Sprintf("- Worked on issues in repository '%s'", event.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("- Starred repository '%s'", event.Repo.Name)
	default:
		return fmt.Sprintf("- %s in repository '%s'", event.Type, event.Repo.Name)
	}
}
