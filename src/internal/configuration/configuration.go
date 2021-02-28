package configuration

import "flag"

type Configuration struct {
	ControllerConfiguration *ControllerConfiguration `yaml:"controllerConfiguration" json:"controllerConfiguration"`
}

func NewConfiguration() (*Configuration, error) {
	configuration := &Configuration{
		ControllerConfiguration: nil,
	}

	controllerConfiguration := NewControllerConfiguration()

	flag.Parse()

	configuration.ControllerConfiguration = controllerConfiguration

	return configuration, nil
}