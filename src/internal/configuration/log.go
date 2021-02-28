package configuration

import (
	"flag"
	validation "github.com/go-ozzo/ozzo-validation"
)

type logConfiguration struct {
	Verbose *bool `yaml:"verbose" json:"verbose"`
}

func NewLogConfiguration() *logConfiguration {
	verbose := flag.Bool("verbose", false, "verbose logging")

	return &logConfiguration{
		Verbose: verbose,
	}
}

func (lc logConfiguration) Validate() error {
	return validation.ValidateStruct(&lc,
		validation.Field(&lc.Verbose, validation.NilOrNotEmpty),
	)
}

