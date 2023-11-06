package global

import (
	"interface/model/common/config"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Config config.Config

func GetConfig(patn string) {
	configBytes, err := os.ReadFile(patn)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = yaml.Unmarshal(configBytes, &Config)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
