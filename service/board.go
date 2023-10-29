package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/req"
	"github.com/savi2w/wesley-chan/presenter/res"
	"github.com/savi2w/wesley-chan/repo"
)

type BoardService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewBoardService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *BoardService {
	return &BoardService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s *BoardService) NewBoard(ctx context.Context, r *req.Board) (resp *res.Board, err error) {
	board := &model.Board{
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
	}

	if err := s.RepoManager.MySQL.Board.InsertBoard(ctx, board); err != nil {
		return nil, err
	}

	return &res.Board{
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
	}, nil
}
