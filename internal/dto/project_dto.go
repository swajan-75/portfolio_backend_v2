package dto

type Create_Project_dto struct {
	Title       string   `json:"title" binding:"required,min=3"`
	Description string   `json:"description" binding:"required"`
	TechStack   []string `json:"tech_stack" binding:"required"`
	LiveURL     string   `json:"live_url"`
	GithubURL   string   `json:"github_url"`
	ImageURL    string   `json:"image_url" binding:"required"`
}

type ProjectResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"tech_stack"`
	LiveURL     string   `json:"live_url"`
	GithubURL   string   `json:"github_url"`
	ImageURL    string   `json:"image_url"`
}