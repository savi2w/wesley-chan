package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
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
		end(logger, err, "failed to initialize the repo manager")
	}

	svc := service.New(cfg, logger, repo)
	ctrl := controller.New(svc, logger)

	svr := server.New(cfg, logger, ctrl)

	if cfg.InternalConfig.RunningLocal {
		if err := svr.Start(); err != nil {
			end(logger, err, "failed to start server")
		}
	} else {
		lambda.Start(svr.Handler)
	}
}

func end(logger *zerolog.Logger, err error, message string) {
	logger.Fatal().Err(err).Msgf("%+v: %+v", message, err)
	time.Sleep(time.Millisecond * 50)

	os.Exit(1)
}
