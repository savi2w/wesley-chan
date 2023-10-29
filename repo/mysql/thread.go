package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/consts"
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

func (t *Thread) SelectByBoardSlug(ctx context.Context, slug string, offset int64) (result []model.Thread, err error) {
	query := `
		SELECT thr.thread_id, thr.board_id, thr.file_id, thr.subject, thr.text_content, thr.created_at, thr.updated_at, thr.deleted_at
		FROM db_wesley_chan.tb_thread thr
		INNER JOIN db_wesley_chan.tb_board brd
		WHERE thr.deleted_at IS NULL
		AND brd.deleted_at IS NULL
		AND brd.slug = ?
		ORDER BY thr.updated_at DESC
		LIMIT ?
		OFFSET ?;`

	if err := t.cli.SelectContext(ctx, &result, query, slug, consts.ThreadItemsPerPage, offset); err != nil {
		return nil, err
	}

	return result, nil
}
