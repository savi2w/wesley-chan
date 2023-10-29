package router

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/server/controller"
)

func Register(svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	file := root.Group("/file")
	file.GET("/test", ctrl.FileController.HandleUploadFile)
}
