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

func (s *Project_Service) CreateProject (ctx context.Context , req dto.Create_Project_dto) error {
	title := strings.ToLower(strings.ReplaceAll(req.Title," ","-"))
	
	exists, err := s.Repo.Check_Is_Exists(ctx,title)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("project with this title already exists")
	}
	project := models.Project{
		Id: uuid.New(),
		Title: req.Title,
		
		Description: req.Description,
		LiveURl: req.LiveURL,
		GithubURL: req.GithubURL,
		ImageLink: req.LiveURL,
		CreatedAt:   time.Now().Unix(),
		TechStack: req.TechStack,
	}

	return   s.Repo.Set_Project(ctx, title, project)

}
func (s *Project_Service) GetProjects(ctx context.Context) (map[string]interface{}, error) {
	return s.Repo.Get_All_Projects(ctx)
}
