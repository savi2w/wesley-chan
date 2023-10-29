package main

import (
	"log"
	"os"
	"time"

	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/logger"
	"github.com/savi2w/wesley-chan/repo"
	"github.com/savi2w/wesley-chan/server"
	"github.com/savi2w/wesley-chan/server/controller"
	"github.com/savi2w/wesley-chan/service"
)

func main() {
	cfg := config.Get()
	logger := logger.New(cfg)

	repo, err := repo.New(cfg)
	if err != nil {
		end(err, "failed to initialize the repo manager")
	}

	svc := service.New(cfg, logger, repo)
	ctrl := controller.New(svc, logger)

	if err := server.New(cfg, logger, ctrl).Start(); err != nil {
		end(err, "failed to start server")
	}
}

func end(err error, message string) {
	log.Fatalf("%+v: %+v", message, err)
	time.Sleep(time.Millisecond * 50)

	os.Exit(1)
}
