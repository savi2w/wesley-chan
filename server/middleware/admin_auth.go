package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savi2w/wesley-chan/config"
)

func RequireAdminKey(cfg *config.Config) echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, ctx echo.Context) (bool, error) {
		return key == cfg.InternalConfig.AdminKey, nil
	})
}
