package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (ctrl *Controller) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
