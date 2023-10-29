package service

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/req"
	"github.com/savi2w/wesley-chan/presenter/res"
	"github.com/savi2w/wesley-chan/repo"
	"github.com/savi2w/wesley-chan/util/stringutil"
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

func (s *FileService) UploadFile(ctx context.Context, r *req.File) (resp *res.File, err error) {
	src, err := r.Header.Open()
	if err != nil {
		return nil, err
	}

	defer src.Close()

	v4, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	fileID := strings.Join([]string{v4.String(), stringutil.GetFileExt(r.Header.Filename)}, ".")

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(s.Config.AWSConfig.Region),
	})

	if err != nil {
		return nil, err
	}

	uploader := s3.New(session)

	_, err = uploader.PutObject(&s3.PutObjectInput{
		Body:   src,
		Bucket: aws.String(s.Config.AWSConfig.FileBucketName),
		Key:    aws.String(fileID),
	})

	if err != nil {
		return nil, err
	}

	err = s.RepoManager.MySQL.File.InsertFile(ctx, &model.File{
		ID: fileID,
	})

	if err != nil {
		return nil, err
	}

	return &res.File{
		ID:  fileID,
		URL: stringutil.GetFileURL(s.Config, fileID),
	}, nil
}
