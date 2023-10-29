package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/model"
)

type File struct {
	cli *sqlx.DB
}

func (f *File) InsertFile(ctx context.Context, file *model.File) (fileID int64, err error) {
	query := ``

	result, err := f.cli.NamedExecContext(ctx, query, file)
	if err != nil {
		return 0, err
	}

	fileID, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return fileID, nil
}
