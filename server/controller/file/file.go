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

// TODO: Implement this controller properly
func (ctrl *Controller) HandleUploadFile(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.String(400, "Bad Request")
	}

	if err := ctrl.svc.File.UploadImage(ctx.Request().Context(), file); err != nil {
		ctrl.svc.File.Logger.Err(err)

		return ctx.String(500, "Internal Server Error")
	}

	return ctx.String(201, "Created")
}
