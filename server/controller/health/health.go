package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func New() *HealthController {
	return &HealthController{}
}

func (ctrl *HealthController) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}
