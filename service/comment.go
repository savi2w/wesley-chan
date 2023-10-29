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

type CommentService struct {
	Config      *config.Config
	Logger      *zerolog.Logger
	RepoManager *repo.RepoManager
}

func NewCommentService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *CommentService {
	return &CommentService{
		Config:      cfg,
		Logger:      logger,
		RepoManager: repo,
	}
}

func (s *CommentService) NewComment(ctx context.Context, r *req.Comment) (resp *res.Comment, err error) {
	comment := &model.Comment{
		ThreadID:    r.ThreadID,
		FileID:      r.FileID,
		TextContent: r.TextContent,
	}

	if err := s.RepoManager.MySQL.Comment.InsertComment(ctx, comment); err != nil {
		return nil, err
	}

	return &res.Comment{
		ThreadID:    r.ThreadID,
		FileID:      r.FileID,
		TextContent: r.TextContent,
	}, nil
}
