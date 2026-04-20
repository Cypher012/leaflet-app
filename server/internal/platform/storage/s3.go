package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appCfg "server/internal/platform/config"
)

type R2Credentials struct {
	AccountID string
	AccessKey string
	SecretKey string
	Bucket    string
	PublicURL string
}

type Storage struct {
	client      *s3.Client
	credentials R2Credentials
}

func New(appCfg appCfg.AppConfig) *Storage {
	cred := R2Credentials{
		AccountID: appCfg.R2AccountID,
		AccessKey: appCfg.R2AccessKey,
		SecretKey: appCfg.R2SecretKey,
		Bucket:    appCfg.R2Bucket,
		PublicURL: appCfg.R2PublicURL,
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cred.AccessKey, cred.SecretKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal("Failed to initialize R2 config:", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cred.AccountID))
	})

	return &Storage{client: client, credentials: cred}
}

// PresignUpload returns a presigned URL the client uploads to directly
func (s *Storage) PresignUpload(ctx context.Context, key string, contentType string) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	req, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.credentials.Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		return "", fmt.Errorf("failed to presign upload: %w", err)
	}

	return req.URL, nil
}

// Delete removes a file from R2
func (s *Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.credentials.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}

// PublicURL returns the public URL for a stored object
func (s *Storage) PublicURL(key string) string {
	return fmt.Sprintf("%s/%s", s.credentials.PublicURL, key)
}
