package srv

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/bot"
	"Scharsch-Bot/discord/discordMember"
	"Scharsch-Bot/discord/embed/srvEmbed"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	config        = conf.GetConf()
	eventJson     types.EventJson
	Session       = bot.Session
	OnlinePlayers []string
	port          = config.SRV.API.Port
	APIUser       = config.SRV.API.User
	APIPassword   = config.SRV.API.Password
)

func playerSRVEventHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, _ := r.BasicAuth()
	if user == "" || pass == "" {
		w.WriteHeader(http.StatusUnauthorized)
	} else if user != APIUser || pass != APIPassword {
		w.WriteHeader(http.StatusForbidden)
	} else if user == APIUser && pass == APIPassword {
		var (
			serverConf conf.Server
			found      = false
		)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&eventJson)
		if err != nil {
			log.Printf("Failed to decode json: %v", err)
		}

		for _, server := range config.Pterodactyl.Servers {
			if server.ServerName == eventJson.Server {
				serverConf = server
				found = true
			}
		}
		if !found {
			log.Printf("Server %v not found in Pterodactyl server config", eventJson.Server)
			return
		}
		userID, onWhitelist := whitelist.GetOwner(eventJson.Name)
		var (
			FooterIcon string
			username   string
		)
		if onWhitelist {
			member, successful := discordMember.GetUserProfile(userID, Session)

			if successful {
				w.WriteHeader(http.StatusNoContent)
				FooterIcon = member.User.AvatarURL("40")
				username = member.User.String()
			} else {
				w.WriteHeader(http.StatusOK)
				FooterIcon = config.Discord.EmbedErrorIcon
			}
		} else {
			w.WriteHeader(http.StatusResetContent)
			FooterIcon = config.Discord.EmbedErrorIcon
		}
		checkAccount(strings.ToLower(eventJson.Name))
		switch eventJson.Type {
		case "chat":
			if serverConf.Chat.Embed {
				messageEmbed := srvEmbed.Chat(eventJson, serverConf, FooterIcon, username, Session)
				for _, channelID := range serverConf.Chat.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
					}
				}
			} else {
				for _, channelID := range serverConf.Chat.ChannelID {
					_, err = Session.ChannelMessageSend(channelID, fmt.Sprintf("%v%v %v", eventJson.Name, serverConf.Chat.Prefix, eventJson.Value))
					if err != nil {
						log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
					}
				}
			}

		case "death":
			if serverConf.SRV.Events.Death {
				messageEmbed := srvEmbed.PlayerDeath(eventJson, serverConf, FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Death embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		case "advancement":
			if serverConf.SRV.Events.Advancement {
				messageEmbed := srvEmbed.PlayerAdvancement(eventJson, serverConf, FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Advancement embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		case "join":
			if serverConf.SRV.Events.Join {
				OnlinePlayers = append(OnlinePlayers, strings.ToLower(eventJson.Name))
				messageEmbed := srvEmbed.PlayerJoin(serverConf, strings.ToLower(eventJson.Name), FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Join embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		case "quit":
			if serverConf.SRV.Events.Quit {
				for i, player := range OnlinePlayers {
					if player == strings.ToLower(eventJson.Name) {
						OnlinePlayers = append(OnlinePlayers[:i], OnlinePlayers[i+1:]...)
						break
					}
				}
				messageEmbed := srvEmbed.PlayerQuit(serverConf, strings.ToLower(eventJson.Name), FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Quit embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		}

		eventJson.Type = ""
	}
}

func checkAccount(Name string) (accounts []string, bannedAccounts []string) {
	userID, onWhitelist := whitelist.GetOwner(Name)
	if config.Whitelist.KickUnWhitelisted {
		if !onWhitelist {
			command := fmt.Sprintf(config.Whitelist.KickCommand, eventJson.Name)
			for _, listedServer := range config.Whitelist.Servers {
				for _, server := range config.Pterodactyl.Servers {
					if server.ServerName == listedServer {
						pterodactyl.SendCommand(command, server.ServerID)
					}
				}
			}
		}
	}
	ListedAccounts := whitelist.ListedAccountsOf(userID, false)
	Banned := whitelist.CheckBans(userID)
	return ListedAccounts, Banned
}
