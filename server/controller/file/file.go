package file

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/service"
)

type Controller struct {
	svc *service.Service
}

func New(svc *service.Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (ctrl *Controller) HandleUploadFile(ctx echo.Context) error {
	return nil
}
