package controller

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/server/controller/board"
	"github.com/savi2w/wesley-chan/server/controller/comment"
	"github.com/savi2w/wesley-chan/server/controller/file"
	"github.com/savi2w/wesley-chan/server/controller/health"
	"github.com/savi2w/wesley-chan/server/controller/thread"
	"github.com/savi2w/wesley-chan/service"
	"github.com/savi2w/wesley-chan/util/resutil"
)

type Controller struct {
	BoardController   *board.Controller
	CommentController *comment.Controller
	FileController    *file.Controller
	HealthController  *health.Controller
	ThreadController  *thread.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	resutil := resutil.New(logger)

	return &Controller{
		BoardController:   board.New(logger, resutil, svc),
		CommentController: comment.New(logger, resutil, svc),
		FileController:    file.New(logger, resutil, svc),
		HealthController:  health.New(),
		ThreadController:  thread.New(logger, resutil, svc),
	}
}
