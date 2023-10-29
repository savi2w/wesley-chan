package repo

import (
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/repo/mysql"
)

type RepoManager struct {
	MySQL *mysql.Repo
}

func New(cfg *config.Config) (*RepoManager, error) {
	mysql, err := mysql.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RepoManager{
		MySQL: mysql,
	}, nil
}
