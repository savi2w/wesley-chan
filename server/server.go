package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
	"github.com/savi2w/wesley-chan/server/controller"
	"github.com/savi2w/wesley-chan/server/middleware"
	"github.com/savi2w/wesley-chan/server/router"
)

var (
	instance *Server
	once     sync.Once
)

type Server struct {
	cfg    *config.Config
	svr    *echoadapter.EchoLambda
	logger *zerolog.Logger
	ctrl   *controller.Controller
}

func New(cfg *config.Config, logger *zerolog.Logger, ctrl *controller.Controller) *Server {
	once.Do(func() {
		svr := echo.New()

		svr.HideBanner = true
		svr.HidePort = true

		middleware.SetMiddlewares(svr, cfg)
		router.Register(cfg, svr, ctrl)

		instance = &Server{
			cfg:    cfg,
			svr:    echoadapter.New(svr),
			logger: logger,
			ctrl:   ctrl,
		}
	})

	return instance
}

func (s *Server) Start() error {
	s.logger.Info().Msg("starting server")

	if err := s.svr.Echo.Start(fmt.Sprintf(":%d", s.cfg.InternalConfig.ServerPort)); err != nil {
		return err
	}

	return nil
}

func (s *Server) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	return s.svr.ProxyWithContext(ctx, req)
}
