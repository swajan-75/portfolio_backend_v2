package service

import (
	"context"
	"fmt"
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
	url, publicID, err := s.StorageRepo.UploadCV(ctx, file, header) // ← get publicID too
	if err != nil {
		return err
	}

	cv := models.CV{
		Name:      name,
		URL:       url,
		PublicID:  publicID, // ← store it
		Active:    false,
		CreatedAt: time.Now().Unix(),
	}

	return s.Repo.Save_CV(ctx, cv)
}

func (s *CV_Service) DeleteCV(ctx context.Context, id string) error {
	// get the CV first to retrieve public_id
	cvs, err := s.Repo.Get_All_CVs(ctx)
	if err != nil {
		return err
	}

	cv, ok := cvs[id].(map[string]interface{})
	if !ok {
		return fmt.Errorf("cv not found")
	}

	// delete from cloudinary
	if publicID, ok := cv["public_id"].(string); ok && publicID != "" {
		if err := s.StorageRepo.DeleteFile(ctx, publicID, "raw"); err != nil {
			fmt.Println("Warning: failed to delete from Cloudinary:", err)
			// don't return error — still delete from Firebase
		}
	}

	// delete from firebase
	return s.Repo.Delete_CV(ctx, id)
}

func (s *CV_Service) GetAllCVs(ctx context.Context) (map[string]interface{}, error) {
	return s.Repo.Get_All_CVs(ctx)
}

func (s *CV_Service) SetActiveCV(ctx context.Context, id string) error {
	return s.Repo.Set_Active_CV(ctx, id)
}


func (s *CV_Service) GetActiveCV(ctx context.Context) (map[string]interface{}, error) {
	return s.Repo.Get_Active_CV(ctx)
}