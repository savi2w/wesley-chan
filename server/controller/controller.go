package controller

import "github.com/savi2w/wesley-chan/server/controller/health"

type Controller struct {
	// Add controllers here
	HealthController *health.HealthController
}

func New() *Controller {
	return &Controller{
		HealthController: health.New(),
	}
}
