package controller

import (
	"github.com/savi2w/wesley-chan/server/controller/file"
	"github.com/savi2w/wesley-chan/server/controller/health"
	"github.com/savi2w/wesley-chan/service"
)

type Controller struct {
	FileController   *file.Controller
	HealthController *health.Controller
}

func New(svc *service.Service) *Controller {
	return &Controller{
		FileController:   file.New(svc),
		HealthController: health.New(),
	}
}
