package utils

import (
	"fmt"
	"io/ioutil"
	"log"
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
		log.Println(config)
	}
	return nil
}

func GetConfig() model.Config {
	return config
}

func GetCollections() map[string]*model.Collection {
	return config.Collections
}

func Debug(v ...any) {
	if config.Debug {
		for _, a := range v {
			fmt.Print(a)
		}
		fmt.Println()
	}
}
