package service

import (
	"context"
	"mime/multipart"
	"portfolio_backend_go/internal/models"
	"portfolio_backend_go/internal/repository"
	"time"
)

type CV_Service struct {
	Repo        *repository.CV_repo
	StorageRepo *repository.Storage_repo
}

func New_CV_Service(repo *repository.CV_repo, storageRepo *repository.Storage_repo) *CV_Service {
	return &CV_Service{Repo: repo, StorageRepo: storageRepo}
}

func (s *CV_Service) UploadCV(ctx context.Context, file multipart.File, header *multipart.FileHeader, name string) error {
	url, err := s.StorageRepo.UploadCV(ctx, file, header)
	if err != nil {
		return err
	}

	cv := models.CV{
		Name:      name,
		URL:       url,
		Active:    false,
		CreatedAt: time.Now().Unix(),
	}

	return s.Repo.Save_CV(ctx, cv)
}

func (s *CV_Service) GetAllCVs(ctx context.Context) (map[string]interface{}, error) {
	return s.Repo.Get_All_CVs(ctx)
}

func (s *CV_Service) SetActiveCV(ctx context.Context, id string) error {
	return s.Repo.Set_Active_CV(ctx, id)
}

func (s *CV_Service) DeleteCV(ctx context.Context, id string) error {
	return s.Repo.Delete_CV(ctx,id)
}

func (s *CV_Service) GetActiveCV(ctx context.Context) (map[string]interface{}, error) {
	return s.Repo.Get_Active_CV(ctx)
}