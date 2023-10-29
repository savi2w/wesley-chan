package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/model"
)

type Comment struct {
	cli *sqlx.DB
}

func (c *Comment) InsertComment(ctx context.Context, com *model.Comment) error {
	query := `INSERT INTO db_wesley_chan.tb_comment (thread_id, file_id, text_content) VALUES (?, ?, ?);`

	_, err := c.cli.ExecContext(ctx, query, com.ThreadID, com.FileID, com.TextContent)
	if err != nil {
		return err
	}

	return nil
}