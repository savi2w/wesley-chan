package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/mapper"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/req"
	"github.com/savi2w/wesley-chan/presenter/res"
	"github.com/savi2w/wesley-chan/repo"
)

type ThreadService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewThreadService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *ThreadService {
	return &ThreadService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s *ThreadService) NewThread(ctx context.Context, r *req.Thread) error {
	thread := &model.Thread{
		BoardID:     r.BoardID,
		FileID:      r.FileID,
		TextContent: r.TextContent,
		Subject:     r.Subject,
	}

	return s.RepoManager.MySQL.Thread.InsertThread(ctx, thread)
}

func (s *ThreadService) SelectByBoardSlug(ctx context.Context, slug string, offset int64) (resp []res.Thread, err error) {
	threads, err := s.RepoManager.MySQL.Thread.SelectByBoardSlug(ctx, slug, offset)
	if err != nil {
		return nil, err
	}

	return mapper.ThreadModelToRes(s.Config, threads), nil
}
