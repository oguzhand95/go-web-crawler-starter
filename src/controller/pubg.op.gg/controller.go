package pubg_op_gg

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/golang/glog"
	"github.com/oguzhand95/go-web-crawler-starter/src/internal/configuration"
	"strconv"
)

const (
	baseUrl          = "https://pubg.op.gg/"
	playerProfileUrl = baseUrl + "user/%s"
	controllerName   = "pubg.op.gg"
)

type PubgOpGgController struct {
	Crawler *colly.Collector
}

func NewPubgOpGgController(crawler *colly.Collector) *PubgOpGgController {
	return &PubgOpGgController{
		Crawler: crawler,
	}
}

func (p *PubgOpGgController) GetName() string {
	return controllerName
}

func (p *PubgOpGgController) Run(controllerConfiguration *configuration.ControllerConfiguration) {
	glog.Infof("[%s] Controller running", controllerName)

	p.Crawler.OnHTML("li.matches-item", func(e *colly.HTMLElement) {
		e.ForEach("div.matches-item__summary", func(index int, e *colly.HTMLElement) {
			e.ForEach(".matches-item__column--damage", func(index int, e *colly.HTMLElement) {
				damage, err := strconv.Atoi(e.ChildText(".matches-item__value"))

				if err != nil {
					glog.Errorf("[%s] could not parse damage from element as an integer:\n%s", controllerName, err.Error())
				}

				glog.Infof("[%s] damage found %d", controllerName, damage)
			})
		})
	})

	p.Crawler.OnRequest(func(r *colly.Request) {
		glog.Infof("[%s] Visiting %s", controllerName, r.URL)
	})

	err := p.Crawler.Visit(fmt.Sprintf(playerProfileUrl, *controllerConfiguration.PUBGUsername))

	if err != nil {
		glog.Errorf("[%s] crawler could not visit:\n%s", controllerName, err.Error())
	}
}
