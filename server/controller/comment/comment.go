package comment

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

func (ctrl *Controller) HandleNewComment(ctx echo.Context) error {
	req, err := payload.GetComment(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	if err := ctrl.svc.Comment.NewComment(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(nil, nil, http.StatusCreated))
}

func (ctrl *Controller) HandleSelectCommentsByThreadID(ctx echo.Context) error {
	thrID, err := payload.GetThreadID(ctx.Param("thread_id"))
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	offset, err := payload.GetOffsetByPage(ctx.QueryParam("page"), consts.CommentItemsPerPage)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	resp, err := ctrl.svc.Comment.SelectByThreadID(ctx.Request().Context(), thrID, offset)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(resp, nil, http.StatusOK))
}
