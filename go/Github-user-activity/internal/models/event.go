package models

// Event represents a GitHub event structure.
type Event struct {
	Type string
	Repo struct {
		Name string
	}
}
