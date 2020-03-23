package controller

import "../service"

// Controller properties
type Controller struct {
	Service service.Service
}

// New controller
func New(service *service.Service) *Controller {
	return &Controller{*service}
}
