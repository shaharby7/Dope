package v1

import (
	"reflect"

	"github.com/shaharby7/Dope/pkg/entities/entity"
)

var ClientManifest = &entity.EntityTypeManifest{
	Name:            "Client",
	BindingSettings: nil,
	ValuesType:      reflect.TypeOf(ClientConfig{}),
	CliOptions: &entity.CliOptions{
		Aliases:      []string{"client"},
		PathTemplate: "modules/clients/{{.Client}}",
	},
}

type ClientConfig struct {
	Apps []string `validate:"required"`
}
