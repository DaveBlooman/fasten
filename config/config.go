package config

import (
	"io/ioutil"

	"github.com/DaveBlooman/fasten/appmeta"
	yaml "gopkg.in/yaml.v1"
)

func LoadFastenFile() (appmeta.AppStack, error) {
	var appStack appmeta.AppStack

	fastenFile, err := ioutil.ReadFile("fasten.yaml")
	if err != nil {
		return appStack, err
	}

	err = yaml.Unmarshal([]byte(fastenFile), &appStack)
	if err != nil {
		return appStack, err
	}

	return appStack, nil
}
