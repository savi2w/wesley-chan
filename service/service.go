package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/repo"
)

type Service struct {
	Board   *BoardService
	Comment *CommentService
	File    *FileService
	Thread  *ThreadService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		Board:   NewBoardService(cfg, logger, repo),
		Comment: NewCommentService(cfg, logger, repo),
		File:    NewFileService(cfg, logger, repo),
		Thread:  NewThreadService(cfg, logger, repo),
	}
}
