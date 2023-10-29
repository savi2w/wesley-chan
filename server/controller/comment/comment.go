package comment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/errors"
	"github.com/savi2w/wesley-chan/service"
	"github.com/savi2w/wesley-chan/valid"
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

func (ctrl *Controller) HandleNewComment(ctx echo.Context) error {
	req, err := valid.GetComment(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err))
	}

	resp, err := ctrl.svc.Comment.NewComment(ctx.Request().Context(), req)
	if err != nil {
		ctrl.logger.Err(err)

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, resp)
}
