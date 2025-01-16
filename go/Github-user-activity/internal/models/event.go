package models

type Event struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Repo  Repo   `json:"repo"`
	Actor Actor  `json:"actor"`
}

type Repo struct {
	Name string `json:"name"`
}

type Actor struct {
	Login string `json:"login"`
}
