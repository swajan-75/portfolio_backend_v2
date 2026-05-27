package dto

type Create_Project_dto struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	TechStack   []string `json:"tech_stack" binding:"required"`
	LiveURL     string   `json:"live_url"`
	GithubURL   string   `json:"github_url"`
	Category    string   `json:"category" binding:"required"`
	ImageLink   string   `json:"image_link"`
	Rank        int      `json:"rank"`
}

type Update_Project_dto struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"tech_stack"`
	LiveURL     string   `json:"live_url"`
	GithubURL   string   `json:"github_url"`
	Category    string   `json:"category"`
	ImageLink   string   `json:"image_link"`
	Rank        int      `json:"rank"`
}

type ProjectResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"tech_stack"`
	LiveURL     string   `json:"live_url"`
	GithubURL   string   `json:"github_url"`
	ImageURL    string   `json:"image_url"`
	Rank        int      `json:"rank"`
}