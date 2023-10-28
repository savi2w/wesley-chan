package server

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/server/controller"
	"github.com/savi2w/wesley-chan/server/router"
)

var (
	instance *Server
	once     sync.Once
)

type Server struct {
	cfg  *config.Config
	svr  *echo.Echo
	log  echo.Logger
	ctrl *controller.Controller
}

func New(cfg *config.Config) *Server {
	once.Do(func() {
		svr := echo.New()

		instance = &Server{
			cfg:  cfg,
			svr:  svr,
			log:  svr.Logger,
			ctrl: controller.New(),
		}
	})

	return instance
}

func (s *Server) Start() error {
	// Set middlewares
	// Set routers
	router.Register(s.svr, s.ctrl)

	if err := s.svr.Start(fmt.Sprintf(":%d", s.cfg.InternalConfig.Port)); err != nil {
		return err
	}

	return nil
}
