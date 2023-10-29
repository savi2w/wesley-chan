package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/savi2w/wesley-chan/config"
)

type Repo struct {
	Board   *Board
	Comment *Comment
	File    *File
	Thread  *Thread

	cli *sqlx.DB
}

func New(cfg *config.Config) (*Repo, error) {
	target := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
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

	cli.DB.SetConnMaxLifetime(time.Minute * 5)
	cli.DB.SetMaxIdleConns(5)
	cli.DB.SetMaxOpenConns(100)

	// Ping the database to ensure the connection is alive
	if err := cli.Ping(); err != nil {
		return nil, err
	}

	return &Repo{
		Board:   &Board{cli: cli},
		Comment: &Comment{cli: cli},
		File:    &File{cli: cli},
		Thread:  &Thread{cli: cli},

		cli: cli,
	}, nil
}
