package thread

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/errors"
	"github.com/savi2w/wesley-chan/payload"
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

func (ctrl *Controller) HandleNewThread(ctx echo.Context) error {
	req, err := payload.GetThread(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err))
	}

	if err := ctrl.svc.Thread.NewThread(ctx.Request().Context(), req); err != nil {
		ctrl.logger.Err(err).Msg(err.Error())

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, nil)
}
