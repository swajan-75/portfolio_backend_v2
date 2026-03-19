package models

type CV struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Active    bool   `json:"active"`
	CreatedAt int64  `json:"created_at"`
}