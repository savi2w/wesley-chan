package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/repo"
)

type FileService struct {
	// Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewFileService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *FileService {
	return &FileService{
		// Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s *FileService) UploadFile() error {
	return nil
}
