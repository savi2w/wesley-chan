package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/model"
)

type Thread struct {
	cli *sqlx.DB
}

func (c *Thread) InsertThread(ctx context.Context, thr *model.Thread) error {
	query := `INSERT INTO db_wesley_chan.tb_thread (file_id, board_id, subject, text_content) VALUES (?, ?, ?, ?);`

	_, err := c.cli.ExecContext(ctx, query, thr.FileID, thr.BoardID, thr.Subject, thr.TextContent)
	if err != nil {
		return err
	}

	return nil
}
