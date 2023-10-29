package router

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/server/controller"
	"github.com/savi2w/wesley-chan/server/middleware"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	board := root.Group("/board")
	board.POST("", ctrl.BoardController.HandleNewBoard, middleware.RequireAdminKey(cfg))
	board.GET("", ctrl.BoardController.HandleGetAll)
	board.GET("/:slug", ctrl.ThreadController.HandleGetThreadsByBoardSlug)

	comment := root.Group("/comment")
	comment.POST("", ctrl.CommentController.HandleNewComment)

	file := root.Group("/file")
	file.POST("", ctrl.FileController.HandleUpload)

	thread := root.Group("/thread")
	thread.POST("", ctrl.ThreadController.HandleNewThread)
}
