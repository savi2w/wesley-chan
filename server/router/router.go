package router

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/server/controller"
	"github.com/savi2w/wesley-chan/server/middleware"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("wc")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	board := root.Group("/board", middleware.RequireAdminKey(cfg))
	board.POST("", ctrl.BoardController.HandleNewBoard)

	comment := root.Group("/comment")
	comment.POST("", ctrl.CommentController.HandleNewComment)

	file := root.Group("/file")
	file.POST("", ctrl.FileController.HandleUpload)

	thread := root.Group("/thread")
	thread.POST("", ctrl.ThreadController.HandleNewThread)
}
