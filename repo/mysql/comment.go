package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/model"
)

type Comment struct {
	cli *sqlx.DB
}

func (c *Comment) InsertComment(ctx context.Context, cmt *model.Comment) error {
	query := `INSERT INTO db_wesley_chan.tb_comment (thread_id, file_id, text_content) VALUES (?, ?, ?);`

	_, err := c.cli.ExecContext(ctx, query, cmt.ThreadID, cmt.FileID, cmt.TextContent)
	if err != nil {
		return err
	}

	return nil
}
