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

func (s *BoardService) NewBoard(ctx context.Context, r *req.Board) error {
	board := &model.Board{
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
	}

	return s.RepoManager.MySQL.Board.InsertBoard(ctx, board)
}

func (s *BoardService) Select(ctx context.Context) (resp []res.Board, err error) {
	boards, err := s.RepoManager.MySQL.Board.Select(ctx)
	if err != nil {
		return nil, err
	}

	for _, board := range boards {
		resp = append(resp, res.Board{
			ID:          board.ID,
			Name:        board.Name,
			Slug:        board.Slug,
			Description: board.Description,
			CreatedAt:   board.CreatedAt,
			UpdatedAt:   board.UpdatedAt,
			DeletedAt:   board.DeletedAt,
		})
	}

	return resp, nil
}
