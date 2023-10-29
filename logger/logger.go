package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/savi2w/wesley-chan/config"
)

func New(cfg *config.Config) *zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Str("service", cfg.InternalConfig.ServiceName).Timestamp().Logger()

	return &logger
}
