package conf

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	confPath = flag.String("configPath", "config/config.yml", "Path to the config file")
	Config   Format
)

func init() {
	flag.Parse()
	GetConf()
}
func GetConf() Format {
	ymlConf, err := os.ReadFile(*confPath)
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}
	if err := yaml.Unmarshal(ymlConf, &Config); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return Config
}
func GetServerConf(serverID string, serverName string) (serverConf Server) {
	if serverID != "" && serverName != "" {
		for _, server := range GetConf().Pterodactyl.Servers {
			if server.ServerID == serverID && server.ServerName == serverName {
				return server
			}
		}
	} else if serverID != "" {
		for _, server := range GetConf().Pterodactyl.Servers {
			if server.ServerID == serverID {
				return server
			}
		}
	} else if serverName != "" {
		for _, server := range GetConf().Pterodactyl.Servers {
			if server.ServerName == serverName {
				return server
			}
		}
	} else {
		return
	}
	return
}
