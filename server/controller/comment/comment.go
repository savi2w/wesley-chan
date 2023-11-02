package comment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/consts"
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

func (ctrl *Controller) HandleNewComment(ctx echo.Context) error {
	req, err := payload.GetComment(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err))
	}

	if err := ctrl.svc.Comment.NewComment(ctx.Request().Context(), req); err != nil {
		ctrl.logger.Err(err).Msg(err.Error())

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (ctrl *Controller) HandleSelectCommentsByThreadID(ctx echo.Context) error {
	thrID, err := payload.GetThreadID(ctx.Param("thread_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err))
	}

	offset, err := payload.GetOffsetByPage(ctx.QueryParam("page"), consts.CommentItemsPerPage)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err))
	}

	resp, err := ctrl.svc.Comment.SelectByThreadID(ctx.Request().Context(), thrID, offset)
	if err != nil {
		ctrl.logger.Err(err).Msg(err.Error())

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, resp)
}
