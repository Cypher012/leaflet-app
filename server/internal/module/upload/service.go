package upload

import (
	"context"
	"fmt"
	"server/internal/platform/storage"

	"github.com/google/uuid"
)

type UploadService struct {
	storage *storage.Storage
}

func NewUploadService(storage *storage.Storage) *UploadService {
	return &UploadService{storage: storage}
}

func (s *UploadService) GetPresignedURL(ctx context.Context, uploadType string, contentType string) (uploadURL string, publicURL string, err error) {
	if !AllowedTypes[uploadType] {
		return "", "", fmt.Errorf("invalid upload type: %s", uploadType)
	}

	key := fmt.Sprintf("%ss/%s", uploadType, uuid.New().String())
	uploadURL, err = s.storage.PresignUpload(ctx, key, contentType)
	if err != nil {
		return "", "", err
	}

	return uploadURL, s.storage.PublicURL(key), nil
}
