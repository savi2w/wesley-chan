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

func New(svc *service.Service, resutil *resutil.ResUtil, log *zerolog.Logger) *Controller {
	return &Controller{
		BoardController:   board.New(log, resutil, svc),
		CommentController: comment.New(log, resutil, svc),
		FileController:    file.New(log, resutil, svc),
		HealthController:  health.New(),
		ThreadController:  thread.New(log, resutil, svc),
	}
}
