package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savi2w/wesley-chan/config"
)

func RequireClientKey(cfg *config.Config) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-Key",
		Validator: func(key string, ctx echo.Context) (bool, error) {
			return key == cfg.InternalConfig.ClientKey, nil
		},
	})
}
