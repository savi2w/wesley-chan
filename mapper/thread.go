package mapper

import (
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/res"
	"github.com/savi2w/wesley-chan/util/stringutil"
)

func ThreadModelToRes(cfg *config.Config, threads []model.Thread) (resp []res.Thread) {
	for _, thread := range threads {
		var file *res.File

		if thread.FileID != nil {
			file = &res.File{
				ID:  *thread.FileID,
				URL: stringutil.GetFileURL(cfg, *thread.FileID),
			}
		}

		resp = append(resp, res.Thread{
			ID:          thread.ID,
			BoardID:     thread.BoardID,
			File:        file,
			Subject:     thread.Subject,
			TextContent: thread.TextContent,
			CreatedAt:   thread.CreatedAt,
			UpdatedAt:   thread.UpdatedAt,
			DeletedAt:   thread.DeletedAt,
		})
	}

	return resp
}
