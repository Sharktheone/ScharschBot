package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const confPath = "config/config2.yml"

func GetConf() Format {
	var config Format
	ymlConf, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}
	if err := yaml.Unmarshal(ymlConf, &config); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return config
}
