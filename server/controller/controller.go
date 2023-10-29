package controller

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/server/controller/comment"
	"github.com/savi2w/wesley-chan/server/controller/file"
	"github.com/savi2w/wesley-chan/server/controller/health"
	"github.com/savi2w/wesley-chan/server/controller/thread"
	"github.com/savi2w/wesley-chan/service"
)

type Controller struct {
	CommentController *comment.Controller
	FileController    *file.Controller
	HealthController  *health.Controller
	ThreadController  *thread.Controller
}

func New(svc *service.Service, log *zerolog.Logger) *Controller {
	return &Controller{
		CommentController: comment.New(log, svc),
		FileController:    file.New(log, svc),
		HealthController:  health.New(),
		ThreadController:  thread.New(log, svc),
	}
}
