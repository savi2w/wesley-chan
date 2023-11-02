package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/req"
	"github.com/savi2w/wesley-chan/presenter/res"
	"github.com/savi2w/wesley-chan/repo"
	"github.com/savi2w/wesley-chan/util/stringutil"
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

func (s *CommentService) NewComment(ctx context.Context, r *req.Comment) error {
	comment := &model.Comment{
		ThreadID:    r.ThreadID,
		FileID:      r.FileID,
		TextContent: r.TextContent,
	}

	return s.RepoManager.MySQL.Comment.InsertComment(ctx, comment)
}

func (s *CommentService) SelectByThreadID(ctx context.Context, thrID string, offset int64) (resp []res.Comment, err error) {
	comments, err := s.RepoManager.MySQL.Comment.SelectByThreadID(ctx, thrID, offset)
	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		var file *res.File

		if comment.FileID != nil {
			file = &res.File{
				ID:  *comment.FileID,
				URL: stringutil.GetFileURL(s.Config, *comment.FileID),
			}
		}

		resp = append(resp, res.Comment{
			ID:          comment.ID,
			ThreadID:    comment.ThreadID,
			File:        file,
			TextContent: comment.TextContent,
			CreatedAt:   comment.CreatedAt,
			UpdatedAt:   comment.UpdatedAt,
			DeletedAt:   comment.DeletedAt,
		})
	}

	return resp, nil
}
