package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/consts"
	"github.com/savi2w/wesley-chan/model"
)

type Comment struct {
	cli *sqlx.DB
}

func (c *Comment) InsertComment(ctx context.Context, cmt *model.Comment) error {
	tx, err := c.cli.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	comment := `INSERT INTO db_wesley_chan.tb_comment (thread_id, file_id, text_content) VALUES (?, ?, ?);`

	_, err = tx.ExecContext(ctx, comment, cmt.ThreadID, cmt.FileID, cmt.TextContent)
	if err != nil {
		return err
	}

	thread := `UPDATE db_wesley_chan.tb_thread SET updated_at = NOW() WHERE id = ?;`

	_, err = tx.ExecContext(ctx, thread, cmt.ThreadID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (c *Comment) SelectByThreadID(ctx context.Context, thrID string, offset int64) (result []model.Comment, err error) {
	query := `
		SELECT cmt.comment_id, cmt.thread_id, cmt.file_id, cmt.text_content, cmt.created_at, cmt.updated_at, cmt.deleted_at
		FROM db_wesley_chan.tb_comment cmt
		INNER JOIN db_wesley_chan.tb_thread thr
		WHERE cmt.deleted_at IS NULL
		AND thr.deleted_at IS NULL
		AND thr.thread_id = ?
		ORDER BY cmt.created_at ASC
		LIMIT ?
		OFFSET ?;`

	if err := c.cli.SelectContext(ctx, &result, query, thrID, consts.CommentItemsPerPage, offset); err != nil {
		return nil, err
	}

	return result, nil
}
