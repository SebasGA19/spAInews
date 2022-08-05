package web

import "github.com/SebasGA19/spAInews/api/internal/controller"

type Backend struct {
	Controller *controller.Controller
}

func NewBackend(c *controller.Controller) *Backend {
	return &Backend{
		Controller: c,
	}
}
