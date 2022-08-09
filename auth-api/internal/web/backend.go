package web

import "github.com/SebasGA19/spAInews/auth-api/internal/controller"

type Backend struct {
	Controller *controller.Controller
}

func NewBackend(c *controller.Controller) *Backend {
	return &Backend{
		Controller: c,
	}
}
