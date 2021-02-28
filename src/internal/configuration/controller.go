package configuration

import (
	"flag"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ControllerConfiguration struct {
	Controllers  *string `yaml:"controllers" json:"controllers"`
	PUBGUsername *string `yaml:"pubgUsername" json:"pubgUsername"`
}

func NewControllerConfiguration() *ControllerConfiguration {
	controllers := flag.String("controllers", "", "comma separated controllers list")
	pubgUsername := flag.String("pubg-username", "", "PUBG username")

	return &ControllerConfiguration{
		Controllers:  controllers,
		PUBGUsername: pubgUsername,
	}
}

func (cc ControllerConfiguration) Validate() error {
	return validation.ValidateStruct(&cc,
		validation.Field(&cc.Controllers, validation.Required),
		validation.Field(&cc.PUBGUsername, validation.Required),
	)
}
