package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Storage_repo struct{}

func New_Storage_repo() *Storage_repo {
	return &Storage_repo{}
}

func (s *Storage_repo) UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader, folder string, resourceType string) (string, error) {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return "", fmt.Errorf("failed to init cloudinary: %w", err)
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:       folder,
		ResourceType: resourceType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return resp.SecureURL, nil
}

func (s *Storage_repo) UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	return s.UploadFile(ctx, file, header, "projects", "image")
}

func (s *Storage_repo) UploadCV(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return "", fmt.Errorf("failed to init cloudinary: %w", err)
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:       "cvs",
		ResourceType: "raw",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload CV: %w", err)
	}

	// force browser to download instead of opening
	url := strings.Replace(resp.SecureURL, "/upload/", "/upload/fl_attachment/", 1)
	return url, nil
}