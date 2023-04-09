package playersrv

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/bot"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strings"
)

var (
	s      = bot.Session
	config = conf.GetConf()
)

type PlayerSrv struct {
	event       *types.WebsocketEvent
	userID      string
	onWhitelist bool
	footerIcon  string
	username    string
	member      *discordgo.Member
	server      *pterodactyl.Server
}

func Decode(eventJson *types.EventJson) (error, int, *PlayerSrv) {
	server, err := pterodactyl.GetServerByName(eventJson.Server)
	if err != nil {
		return fmt.Errorf("failed to get server: %v", err), http.StatusNotFound, nil
	}
	var (
		statusCode int
		errMsg     error
		pSrv       = &PlayerSrv{
			event:  &types.WebsocketEvent{},
			server: server,
		}
	)
	if userID, onWhitelist := whitelist.GetOwner(eventJson.Name); onWhitelist {
		pSrv.onWhitelist = true
		pSrv.userID = userID
		if member, err := s.GetUserProfile(userID); err != nil {
			statusCode = http.StatusOK
			errMsg = fmt.Errorf("failed to get user profile: %v", err)
			pSrv.footerIcon = config.Discord.EmbedErrorIcon
			pSrv.username = fmt.Sprintf("%v (User not on whitelist)", eventJson.Name)
			pSrv.member = member
		} else {
			statusCode = http.StatusNoContent
			pSrv.footerIcon = member.User.AvatarURL("40")
			pSrv.username = member.User.String()
			pSrv.member = member
		}
	} else {
		statusCode = http.StatusAccepted
		pSrv.onWhitelist = false
		pSrv.footerIcon = config.Discord.EmbedErrorIcon
	}
	checkAccount(strings.ToLower(eventJson.Name))

	return errMsg, statusCode, pSrv

}

func checkAccount(Name string) ([]string, []string) {
	userID, onWhitelist := whitelist.GetOwner(Name)
	if config.Whitelist.KickUnWhitelisted {
		if !onWhitelist {
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
	return whitelist.ListedAccountsOf(userID, false), whitelist.CheckBans(userID)
}
