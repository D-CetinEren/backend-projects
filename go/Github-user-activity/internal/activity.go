package internal

import "fmt"

func DisplayUserActivity(username string) error {
	events, err := FetchUserActivity(username)
	if err != nil {
		return err
	}

	if len(events) == 0 {
		fmt.Printf("No recent activity found for GitHub user '%s'.\n", username)
		return nil
	}

	fmt.Printf("Recent activity for GitHub user '%s':\n", username)
	for _, event := range events {
		fmt.Println(FormatEvent(event))
	}

	return nil
}
