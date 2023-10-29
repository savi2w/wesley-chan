package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/model"
)

type Board struct {
	cli *sqlx.DB
}

func (b *Board) InsertBoard(ctx context.Context, board *model.Board) error {
	query := `INSERT INTO db_wesley_chan.tb_board (name, slug, description) VALUES (?, ?, ?);`

	_, err := b.cli.ExecContext(ctx, query, board.Name, board.Slug, board.Description)
	if err != nil {
		return err
	}

	return nil
}
