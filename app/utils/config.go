package utils

import (
	"io/ioutil"
	"tcglocator/model"

	"gopkg.in/yaml.v3"
)

var config model.Config = model.Config{}

func LoadConfig(url string) error {
	if yfile, err := ioutil.ReadFile("config/settings.yml"); err != nil {
		return err
	} else {
		if err2 := yaml.Unmarshal(yfile, &config); err2 != nil {
			return err2
		}
	}
	return nil
}

func GetConfig() model.Config {
	return config
}
