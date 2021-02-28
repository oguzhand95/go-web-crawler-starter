package main

import (
	"github.com/golang/glog"
	"github.com/oguzhand95/go-web-crawler-starter/src/internal/configuration"
)

func main() {
	configuration, err := configuration.NewConfiguration()

	if err != nil {
		glog.Fatalf("Failed to retrieve configuration:\n%s", err.Error())
	}

	glog.Infof("controllers: %s", configuration.ControllerConfiguration.Controllers)
}