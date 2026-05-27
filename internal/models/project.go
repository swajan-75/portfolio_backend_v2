package models

import "github.com/google/uuid"

type Project struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	TechStack   []string  `json:"tech_stack"`
	LiveURl     string    `json:"live_url"`
	GithubURL   string    `json:"github_url"`
	ImageLink   string    `json:"image_link"`
	Rank        int       `json:"rank"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
}