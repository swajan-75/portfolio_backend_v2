package service

import (
	"context"
	"mime/multipart"
	"portfolio_backend_go/internal/repository"
)

type Storage_Service struct {
	Repo *repository.Storage_repo
}

func New_Storage_Service(repo *repository.Storage_repo) *Storage_Service {
	return &Storage_Service{Repo: repo}
}

func (s *Storage_Service) UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	return s.Repo.UploadImage(ctx, file, header)
}