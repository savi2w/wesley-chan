package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/repo"
)

type Service struct {
	File *FileService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		File: NewFileService(cfg, logger, repo),
	}
}
