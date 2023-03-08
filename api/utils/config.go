package utils

import (
	"fmt"
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

		for _, filename := range config.SitesConfig {
			if sfile, err := ioutil.ReadFile("config/sites/" + filename); err != nil {
				return err
			} else {
				site := model.Site{}
				if err2 := yaml.Unmarshal(sfile, &site); err2 != nil {
					return err2
				}
				config.Sites = append(config.Sites, &site)
			}
		}

		config.Locators = map[string]*model.Locator{}
	}

	Debug(config)
	return nil
}

func GetConfig() model.Config {
	return config
}

func Debug(v ...any) {
	if config.Debug {
		for _, a := range v {
			fmt.Print(a)
		}
		fmt.Println()
	}
}
