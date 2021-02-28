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

	if err := controllerConfiguration.Validate(); err != nil {
		return nil, err
	}

	configuration.ControllerConfiguration = controllerConfiguration

	return configuration, nil
}