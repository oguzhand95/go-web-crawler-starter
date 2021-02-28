package configuration

import "flag"

type Configuration struct {
	LogConfiguration *logConfiguration
}

type Configurable interface {
	SetFlags()
}

func NewConfiguration() (*Configuration, error) {
	configuration := &Configuration{
		LogConfiguration: nil,
	}

	logConfiguration := NewLogConfiguration()

	flag.Parse()

	configuration.LogConfiguration = logConfiguration

	return configuration, nil
}