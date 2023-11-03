package board

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/payload"
	"github.com/savi2w/wesley-chan/service"
	"github.com/savi2w/wesley-chan/util/resutil"
)

type Controller struct {
	logger  *zerolog.Logger
	resutil *resutil.ResUtil
	svc     *service.Service
}

func New(logger *zerolog.Logger, resutil *resutil.ResUtil, svc *service.Service) *Controller {
	return &Controller{
		logger:  logger,
		resutil: resutil,
		svc:     svc,
	}
}

func (ctrl *Controller) HandleNewBoard(ctx echo.Context) error {
	req, err := payload.GetBoard(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	if err := ctrl.svc.Board.NewBoard(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(nil, nil, http.StatusCreated))
}

func (ctrl *Controller) HandleSelect(ctx echo.Context) error {
	resp, err := ctrl.svc.Board.Select(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(resp, nil, http.StatusOK))
}
