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

func GetServerConf(serverID string, serverName string) (serverConf Server) {
	if serverID != "" && serverName != "" {
		for _, server := range GetConf().Pterodactyl.Servers {
			if server.ServerID == serverID && server.ServerName == serverName {
				return Server(server)
			}
		}
	} else if serverID != "" {
		for _, server := range GetConf().Pterodactyl.Servers {
			if server.ServerID == serverID {
				return Server(server)
			}
		}
	} else if serverName != "" {
		for _, server := range GetConf().Pterodactyl.Servers {
			if server.ServerName == serverName {
				return Server(server)
			}
		}
	} else {
		return
	}
	return
}
