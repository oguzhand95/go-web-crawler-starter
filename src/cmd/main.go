package main

import (
	"github.com/golang/glog"
	controllerInterface "github.com/oguzhand95/go-web-crawler-starter/src/controller"
	"github.com/oguzhand95/go-web-crawler-starter/src/controller/pubg.op.gg"
	"github.com/oguzhand95/go-web-crawler-starter/src/internal/configuration"
	"strings"
)

func main() {
	controllerMap := make(map[string]controllerInterface.Controller)

	appConfiguration, err := configuration.NewConfiguration()

	if err != nil {
		glog.Fatalf("Failed to retrieve configuration:\n%s", err.Error())
	}

	if appConfiguration != nil {
		glog.Infof("controllers: %s", *appConfiguration.ControllerConfiguration.Controllers)
	}

	registerControllers(controllerMap)

	if appConfiguration == nil {
		glog.Fatal("Application configuration is null")
	}

	listOfControllers := getControllerListFromConfiguration(appConfiguration.ControllerConfiguration)

	err = runControllers(controllerMap, listOfControllers)

	if err != nil {
		glog.Fatalf("failed during execution of controllers:%s\n", err.Error())
	}
}

func registerControllers(controllerMap map[string]controllerInterface.Controller) {
	var pubgOpGgController controllerInterface.Controller = controller.NewPubgOpGgController()

	controllerMap[pubgOpGgController.GetName()] = pubgOpGgController
}

func runControllers(controllerMap map[string]controllerInterface.Controller, controllersToRun []string) error {
	for _, controllerName := range controllersToRun {
		if val, ok := controllerMap[controllerName]; ok {
			glog.Infof("running controller %q", val.GetName())
			go val.Run()
		} else {
			glog.Warningf("controller %q could not found", controllerName)
		}
	}

	return nil
}

func getControllerListFromConfiguration(controllerConfiguration *configuration.ControllerConfiguration) []string {
	return strings.Split(*controllerConfiguration.Controllers, ",")
}