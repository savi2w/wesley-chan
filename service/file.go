package service

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/repo"
)

type FileService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewFileService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *FileService {
	return &FileService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s *FileService) UploadImage(ctx context.Context, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// isImage, err := imageutil.IsImage(src)
	// if err != nil {
	// 	return err
	// }

	// if !isImage {
	// 	return errors.New("unknown file type")
	// }

	v4, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	fileKey := v4.String() + ".jpg"

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(s.Config.AWSConfig.Region),
	})

	if err != nil {
		return err
	}

	uploader := s3.New(session)

	_, err = uploader.PutObject(&s3.PutObjectInput{
		Body:   src,
		Bucket: aws.String(s.Config.AWSConfig.FileBucketName),
		Key:    aws.String(fileKey),
	})

	if err != nil {
		return err
	}

	err = s.RepoManager.MySQL.File.InsertFile(ctx, &model.File{
		Key: fileKey,
	})

	if err != nil {
		return err
	}

	return nil
}
