package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/saurav11sarkar/ticket/internal/config"
)

const maxImageFileSize int64 = 5 * 1024 * 1024

type imageUploadAPI interface {
	Upload(ctx context.Context, file any, uploadParams uploader.UploadParams) (*uploader.UploadResult, error)
}

type CloudinaryUploader struct {
	uploadAPI imageUploadAPI
	folder    string
}

func NewCloudinaryUploader(cfg *config.Config) (*CloudinaryUploader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cloudinary config is required")
	}
	if strings.TrimSpace(cfg.CloudinaryCloudName) == "" ||
		strings.TrimSpace(cfg.CloudinaryAPIKey) == "" ||
		strings.TrimSpace(cfg.CloudinaryAPISecret) == "" {
		return nil, fmt.Errorf("cloudinary cloud name, API key, and API secret are required")
	}

	cld, err := cloudinary.NewFromParams(
		cfg.CloudinaryCloudName,
		cfg.CloudinaryAPIKey,
		cfg.CloudinaryAPISecret,
	)
	if err != nil {
		return nil, fmt.Errorf("initialize cloudinary: %w", err)
	}

	return &CloudinaryUploader{
		uploadAPI: &cld.Upload,
		folder:    "uploadImage",
	}, nil
}

func (u *CloudinaryUploader) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	if u == nil || u.uploadAPI == nil {
		return "", fmt.Errorf("cloudinary uploader is not initialized")
	}
	if fileHeader == nil {
		return "", fmt.Errorf("image file is required")
	}

	if fileHeader.Size > maxImageFileSize {
		return "", fmt.Errorf("image exceeds the maximum size of %d bytes", maxImageFileSize)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer file.Close()
	contentType := fileHeader.Header.Get("Content-Type")
	switch contentType {
	case "image/jpeg", "image/png", "image/gif", "image/webp":
	default:
		return "", fmt.Errorf("unsupported image type: %s", contentType)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := u.uploadAPI.Upload(
		ctx,
		file,
		uploader.UploadParams{
			Folder:       u.folder,
			ResourceType: "image",
		},
	)
	if err != nil {
		return "", fmt.Errorf("upload image: %w", err)
	}

	return result.SecureURL, nil
}
