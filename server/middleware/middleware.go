package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/consts"
)

func SetMiddlewares(e *echo.Echo, cfg *config.Config) {
	if !cfg.InternalConfig.RunningLocal {
		e.Pre(middleware.HTTPSRedirect())

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"wesley-chan.dev"},
		}))
	}

	e.Pre(middleware.BodyLimit(consts.BodyLimit))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.Secure())

	e.Use(middleware.Recover())

	e.Use(middleware.ContextTimeout(consts.Timeout))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: consts.Timeout,
	}))
}
