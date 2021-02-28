package controller

import (
	"github.com/golang/glog"
)

const (
	baseUrl = "https://pubg.op.gg/"
	controllerName = "pubg.op.gg"
)

type PubgOpGgController struct {

}

func NewPubgOpGgController() *PubgOpGgController {
	return &PubgOpGgController{}
}

func (p *PubgOpGgController) GetName() string {
	return controllerName
}

func (p *PubgOpGgController) Run() {
	glog.Info("Controller running: %s", controllerName)
}