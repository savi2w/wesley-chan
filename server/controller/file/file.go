package file

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/errors"
	"github.com/savi2w/wesley-chan/presenter/req"
	"github.com/savi2w/wesley-chan/service"
)

type Controller struct {
	logger *zerolog.Logger
	svc    *service.Service
}

func New(logger *zerolog.Logger, svc *service.Service) *Controller {
	return &Controller{
		logger: logger,
		svc:    svc,
	}
}

func (ctrl *Controller) HandleUpload(ctx echo.Context) error {
	header, err := ctx.FormFile("target")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err))
	}

	req := req.File{
		Header: header,
	}

	resp, err := ctrl.svc.File.UploadFile(ctx.Request().Context(), &req)
	if err != nil {
		ctrl.logger.Err(err).Msg(err.Error())

		// It's important to not leak any information about the error to the client.
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, resp)
}
