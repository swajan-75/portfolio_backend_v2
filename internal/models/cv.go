package models

type CV struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	PublicID  string `json:"public_id"`
	Active    bool   `json:"active"`
	CreatedAt int64  `json:"created_at"`
}