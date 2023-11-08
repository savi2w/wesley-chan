package mapper

import (
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/res"
	"github.com/savi2w/wesley-chan/util/stringutil"
)

func CommentModelToRes(cfg *config.Config, comments []model.Comment) (resp []res.Comment) {
	for _, comment := range comments {
		var file *res.File

		if comment.FileID != nil {
			file = &res.File{
				ID:  *comment.FileID,
				URL: stringutil.GetFileURL(cfg, *comment.FileID),
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

	return resp
}
