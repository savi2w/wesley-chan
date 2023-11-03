package resutil

import (
	"github.com/rs/zerolog"
)

type ResUtil struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) *ResUtil {
	return &ResUtil{
		logger: logger,
	}
}

type Wrapped struct {
	Success bool        `json:"success"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *ResUtil) Wrap(data interface{}, err error, status int) (code int, resp *Wrapped) {
	resp = &Wrapped{}

	if err != nil {
		switch {
		case status >= 500:
			r.logger.Error().Err(err).Msg(err.Error())
			resp.Message = nil
		case status >= 400:
			message := err.Error()

			r.logger.Warn().Err(err).Msg(message)
			resp.Message = &message
		default:
			r.logger.Info().Err(err).Msg(err.Error())
			resp.Message = nil
		}

		resp.Success = false
		resp.Data = nil

		return status, resp
	}

	resp.Success = true
	resp.Message = nil
	resp.Data = data

	return status, resp
}
