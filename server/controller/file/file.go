package file

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/presenter/req"
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

func (ctrl *Controller) HandleUpload(ctx echo.Context) error {
	header, err := ctx.FormFile("target")
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	req := req.File{
		Header: header,
	}

	resp, err := ctrl.svc.File.UploadFile(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(resp, nil, http.StatusCreated))
}
