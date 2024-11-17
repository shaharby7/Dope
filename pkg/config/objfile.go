package config

import (
	"fmt"
	"os"

	"github.com/shaharby7/Dope/pkg/utils"
	t "github.com/shaharby7/Dope/types"
	"gopkg.in/yaml.v3"
)

type sDopeObjectFile t.DopeObjectFile[any]

// func (c *sDopeObjectFile) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	return nil;
// }

func (d *sDopeObjectFile) UnmarshalYAML(value *yaml.Node) error {
	var h struct {
		Type t.DOPE_OBJECTS `validate:"required" yaml:"type"`
	}
	if err := value.Decode(&h); err != nil {
		return utils.FailedBecause("infer type of obj", err)
	}
	switch h.Type { // TODO: a bit les repetitiveness :)
	case t.DOPE_OBJECT_PROJECT:
		var res t.DopeObjectFile[t.ProjectConfig]
		if err := value.Decode(&res); err != nil {
			return utils.FailedBecause(fmt.Sprintf("parse values for Project %s", d.Name), err)
		}
		res.Values.Name = res.Name
		res.Values.Description = res.Description
		d.Values = res.Values
		d.Name = res.Name
		d.Description = res.Description
		d.Binding = res.Binding
		d.Api = res.Api
	case t.DOPE_OBJECT_APP:
		var res t.DopeObjectFile[t.AppConfig]
		if err := value.Decode(&res); err != nil {
			return utils.FailedBecause(fmt.Sprintf("parse values for App %s", d.Name), err)
		}
		res.Values.Name = res.Name
		res.Values.Description = res.Description
		d.Values = res.Values
		d.Name = res.Name
		d.Description = res.Description
		d.Binding = res.Binding
		d.Api = res.Api
	case t.DOPE_OBJECT_APP_ENV:
		var res t.DopeObjectFile[t.AppEnvConfig]
		if err := value.Decode(&res); err != nil {
			return utils.FailedBecause(fmt.Sprintf("parse values for AppEnv %s", d.Name), err)
		}
		res.Values.Name = res.Name
		res.Values.Description = res.Description
		d.Values = res.Values
		d.Name = res.Name
		d.Description = res.Description
		d.Binding = res.Binding
		d.Api = res.Api
	case t.DOPE_OBJECT_ENV:
		var res t.DopeObjectFile[t.EnvConfig]
		if err := value.Decode(&res); err != nil {
			return utils.FailedBecause(fmt.Sprintf("parse values for Env %s", d.Name), err)
		}
		res.Values.Name = res.Name
		res.Values.Description = res.Description
		d.Values = res.Values
		d.Name = res.Name
		d.Description = res.Description
		d.Binding = res.Binding
		d.Api = res.Api
	}
	d.Type = h.Type
	return nil
}

func readDopeObjFile(path string) (t.DOPE_OBJECTS, *sDopeObjectFile, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", nil, utils.FailedBecause(fmt.Sprintf("read file %s", path), err)
	}
	var obj *sDopeObjectFile
	err = yaml.Unmarshal(content, &obj)
	if err != nil {
		return "", nil, utils.FailedBecause(fmt.Sprintf("unmarshal file %s", path), err)
	}
	return obj.Type, obj, nil
}
