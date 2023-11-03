package thread

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/consts"
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

func (ctrl *Controller) HandleNewThread(ctx echo.Context) error {
	req, err := payload.GetThread(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	if err := ctrl.svc.Thread.NewThread(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(nil, nil, http.StatusCreated))
}

func (ctrl *Controller) HandleSelectThreadsByBoardSlug(ctx echo.Context) error {
	slug, err := payload.GetSlug(ctx.Param("board_slug"))
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	offset, err := payload.GetOffsetByPage(ctx.QueryParam("page"), consts.ThreadItemsPerPage)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	resp, err := ctrl.svc.Thread.SelectByBoardSlug(ctx.Request().Context(), slug, offset)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(resp, nil, http.StatusOK))
}
