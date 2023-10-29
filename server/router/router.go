package router

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/server/controller"
)

func Register(svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("wc")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	comment := root.Group("/comment")
	comment.POST("", ctrl.CommentController.HandleNewComment)

	file := root.Group("/file")
	file.POST("", ctrl.FileController.HandleUpload)

	thread := root.Group("/thread")
	thread.POST("", ctrl.ThreadController.HandleNewThread)
}
