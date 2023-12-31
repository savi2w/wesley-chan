package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/model"
)

type File struct {
	cli *sqlx.DB
}

func (f *File) InsertFile(ctx context.Context, file *model.File) error {
	query := `INSERT INTO db_wesley_chan.tb_file (file_id, original_file_name) VALUES (?, ?);`

	_, err := f.cli.ExecContext(ctx, query, file.ID, file.OriginalFileName)
	if err != nil {
		return err
	}

	return nil
}
