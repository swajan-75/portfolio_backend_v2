package service

import (
	"context"
	"errors"
	"portfolio_backend_go/internal/dto"
	"portfolio_backend_go/internal/models"
	"portfolio_backend_go/internal/repository"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Project_Service struct {
	Repo *repository.Project_repo
}

func New_Project_Service(repo *repository.Project_repo) *Project_Service {
	return &Project_Service{Repo: repo}
}

func titleToSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func (s *Project_Service) CreateProject(ctx context.Context, req dto.Create_Project_dto) error {
	slug := titleToSlug(req.Title)

	exists, err := s.Repo.Check_Is_Exists(ctx, slug)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("project with this title already exists")
	}

	project := models.Project{
		Id:          uuid.New(),
		Title:       req.Title,
		Category:    req.Category,
		Description: req.Description,
		LiveURl:     req.LiveURL,
		GithubURL:   req.GithubURL,
		ImageLink:   req.ImageLink,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		TechStack:   req.TechStack,
	}

	return s.Repo.Set_Project(ctx, slug, project)
}

func (s *Project_Service) GetProjects(ctx context.Context) (map[string]interface{}, error) {
	return s.Repo.Get_All_Projects(ctx)
}

func (s *Project_Service) GetProjectBySlug(ctx context.Context, slug string) (map[string]interface{}, error) {
	project, err := s.Repo.Get_Project_By_Slug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New("project not found")
	}
	return project, nil
}

func (s *Project_Service) UpdateProject(ctx context.Context, slug string, req dto.Update_Project_dto) error {
	exists, err := s.Repo.Check_Is_Exists(ctx, slug)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("project not found")
	}

	updates := map[string]interface{}{
		"updated_at": time.Now().Unix(),
	}

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.LiveURL != "" {
		updates["live_url"] = req.LiveURL
	}
	if req.GithubURL != "" {
		updates["github_url"] = req.GithubURL
	}
	if req.ImageLink != "" {
		updates["image_link"] = req.ImageLink
	}
	if len(req.TechStack) > 0 {
		updates["tech_stack"] = req.TechStack
	}

	return s.Repo.Update_Project(ctx, slug, updates)
}

func (s *Project_Service) DeleteProject(ctx context.Context, slug string) error {
	exists, err := s.Repo.Check_Is_Exists(ctx, slug)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("project not found")
	}
	return s.Repo.Delete_Project(ctx, slug)
}