package configuration

import (
	"flag"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ControllerConfiguration struct {
	Controllers *string `yaml:"controllers" json:"controllers"`
}

func NewControllerConfiguration() *ControllerConfiguration {
	controllers := flag.String("controllers", "", "comma separated controllers list")

	return &ControllerConfiguration{
		Controllers: controllers,
	}
}

func (cc ControllerConfiguration) Validate() error {
	return validation.ValidateStruct(&cc,
		validation.Field(&cc.Controllers, validation.Required),
	)
}