package playersrv

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/whitelist/whitelist"
	"fmt"
	"log"
)

var (
	config = conf.GetConf()
)

func CheckAccount(Name string) ([]string, []string) {
	owner := whitelist.GetOwner(Name, nil)
	if config.Whitelist.KickUnWhitelisted {
		if !owner.Whitelisted {
			command := fmt.Sprintf(config.Whitelist.KickCommand, Name)
			for _, listedServer := range config.Whitelist.Servers {
				for _, server := range config.Pterodactyl.Servers {
					if server.ServerName == listedServer {
						if err := pterodactyl.SendCommand(command, server.ServerID); err != nil {
							log.Printf("Failed to send command to server %v: %v", server.ServerID, err)
						}
					}
				}
			}
		}
	}
	return owner.Players, owner.BannedPlayers
}
