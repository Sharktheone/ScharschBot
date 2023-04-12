package conf

import (
	"Scharsch-Bot/config"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	confPath = flag.String("configPath", "config.yml", "Path to the config file (default: config/config.yml)")
	Config   Format
)

func init() {
	flag.Parse()
	GetConf()
}
func GetConf() Format {
	ymlConf, err := os.ReadFile(*confPath)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := config.StandardConf.ReadFile("config.yml")
			if err != nil {
				log.Fatalf("Failed to get default config: %v", err)
			}
			if err := os.WriteFile(*confPath, f, 0644); err != nil {
				log.Fatalf("Failed to write default config: %v", err)
			}
			fmt.Printf("No config found, created default config at %s", *confPath)
			os.Exit(0)
		}
		log.Fatalf("Failed to get config: %v", err)
	}
	if err := yaml.Unmarshal(ymlConf, &Config); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return Config
}
