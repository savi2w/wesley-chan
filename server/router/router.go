package router

import (
	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/server/controller"
	"github.com/savi2w/wesley-chan/server/middleware"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("", middleware.RequireClientKey(cfg))
	root.GET("/health", ctrl.HealthController.HealthCheck)

	file := root.Group("/file")
	file.POST("/upload", ctrl.FileController.HandleUpload)

	board := root.Group("/board")
	board.POST("", ctrl.BoardController.HandleNewBoard, middleware.RequireAdminKey(cfg))
	board.GET("", ctrl.BoardController.HandleSelect)

	thread := board.Group("/:board_slug/thread")
	thread.POST("", ctrl.ThreadController.HandleNewThread)
	thread.GET("", ctrl.ThreadController.HandleSelectThreadsByBoardSlug)

	comment := thread.Group("/:thread_id/comment")
	comment.POST("", ctrl.CommentController.HandleNewComment)
	comment.GET("", ctrl.CommentController.HandleSelectCommentsByThreadID)
}
