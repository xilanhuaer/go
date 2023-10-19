package utils

import (
	"interface/model/common/config"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYaml(path string) config.Config {
	configYaml, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error:%v", err)
	}
	var config config.Config
	if err = yaml.Unmarshal(configYaml, &config); err != nil {
		log.Fatalf("error:%v", err)
	}
	return config
}
