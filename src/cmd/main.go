package main

import (
	"context"
	"github.com/gocolly/colly/v2"
	"github.com/golang/glog"
	controllerInterface "github.com/oguzhand95/go-web-crawler-starter/src/controller"
	"github.com/oguzhand95/go-web-crawler-starter/src/controller/pubg.op.gg"
	"github.com/oguzhand95/go-web-crawler-starter/src/internal/configuration"
	"strings"
	"sync"
	"time"
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

	crawler := createCrawler()

	registerControllers(controllerMap, crawler)

	if appConfiguration == nil {
		glog.Fatal("Application configuration is null")
	}

	listOfControllers := getControllerListFromConfiguration(appConfiguration.ControllerConfiguration)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	waitGroup := &sync.WaitGroup{}

	err = runControllers(ctx, waitGroup, appConfiguration.ControllerConfiguration, controllerMap, listOfControllers)

	if err != nil {
		glog.Fatalf("failed during execution of controllers:%s\n", err.Error())
	}

	waitGroup.Wait()
}

func registerControllers(controllerMap map[string]controllerInterface.Controller, crawler *colly.Collector) {
	var pubgOpGgController controllerInterface.Controller = controller.NewPubgOpGgController(crawler)

	controllerMap[pubgOpGgController.GetName()] = pubgOpGgController
}

func runControllers(ctx context.Context, waitGroup *sync.WaitGroup,
	controllerConfiguration *configuration.ControllerConfiguration,
	controllerMap map[string]controllerInterface.Controller, controllersToRun []string) error {
	for _, controllerName := range controllersToRun {
		if val, ok := controllerMap[controllerName]; ok {
			glog.Infof("running controller %q", val.GetName())
			waitGroup.Add(1)

			go func(ctx context.Context, waitGroup *sync.WaitGroup) {
				val.Run(controllerConfiguration)
				waitGroup.Done()
				<-ctx.Done()
			}(ctx, waitGroup)
		} else {
			glog.Warningf("controller %q could not found", controllerName)
		}
	}

	return nil
}

func getControllerListFromConfiguration(controllerConfiguration *configuration.ControllerConfiguration) []string {
	return strings.Split(*controllerConfiguration.Controllers, ",")
}

func createCrawler() *colly.Collector {
	return colly.NewCollector()
}
