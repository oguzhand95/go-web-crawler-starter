package controller

import "github.com/oguzhand95/go-web-crawler-starter/src/internal/configuration"

type Controller interface {
	GetName() string
	Run(controllerConfiguration *configuration.ControllerConfiguration)
}
