package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/config"
)

type Repo struct {
	File *File
	cli  *sqlx.DB
}

func New(cfg *config.Config) (*Repo, error) {
	target := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.MySQLConfig.Username,
		cfg.MySQLConfig.Password,
		cfg.MySQLConfig.Host,
		cfg.MySQLConfig.Port,
		cfg.MySQLConfig.Database,
	)

	cli, err := sqlx.Connect("mysql", target)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is alive
	if err := cli.Ping(); err != nil {
		return nil, err
	}

	return &Repo{
		File: &File{cli: cli},
		cli:  cli,
	}, nil
}
