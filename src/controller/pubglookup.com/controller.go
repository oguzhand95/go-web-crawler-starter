package pubglookup_com

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/golang/glog"
	"github.com/oguzhand95/go-web-crawler-starter/src/internal/configuration"
	"strings"
)

const (
	baseUrl          = "https://www.pubglookup.com/"
	playerProfileUrl = baseUrl + "players/%s/%s" // %s=platform %s=pubg-username
	controllerName   = "pubglookup.com"
	platform         = "steam"
)

type PubgLookupController struct {
	Crawler *colly.Collector
}

func NewPubgLookupController(crawler *colly.Collector) *PubgLookupController {
	return &PubgLookupController{
		Crawler: crawler,
	}
}

func (p *PubgLookupController) GetName() string {
	return controllerName
}

func (p *PubgLookupController) Run(controllerConfiguration *configuration.ControllerConfiguration) {
	glog.Infof("[%s] Controller running", controllerName)

	p.Crawler.OnHTML(".header-body", func(e *colly.HTMLElement) {
		e.ForEach("h2", func(i int, e *colly.HTMLElement) {
			text := strings.Split(strings.Trim(e.Text, " "), " ")[0]

			glog.Infof("[%s] Found profile name: %s", controllerName, text)
		})
	})

	p.Crawler.OnRequest(func(r *colly.Request) {
		glog.Infof("[%s] Visiting %s", controllerName, r.URL)
	})

	err := p.Crawler.Visit(fmt.Sprintf(playerProfileUrl, platform, *controllerConfiguration.PUBGUsername))

	if err != nil {
		glog.Errorf("[%s] crawler could not visit:\n%s", controllerName, err.Error())
	}
}
