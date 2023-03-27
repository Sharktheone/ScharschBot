package handlers

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/bot"
	"Scharsch-Bot/discord/embed/srvEmbed"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var (
	config = conf.GetConf()
	s      = bot.Session
)

func PlayerSRVEventHandler(c *gin.Context) {
	var eventJson *types.EventJson
	if err := json.NewDecoder(c.Request.Body).Decode(&eventJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to decode json: %v", err),
		})
		return
	}
	server, err := pterodactyl.GetServerByName(eventJson.Server)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Failed to get server: %v", err),
		})
		return
	}

	userID, onWhitelist := whitelist.GetOwner(eventJson.Name)
	var (
		FooterIcon string
		username   string
	)
	if onWhitelist {
		member, err := s.GetUserProfile(userID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": fmt.Sprintf("Failed to get user profile: %v", err),
			})
			FooterIcon = config.Discord.EmbedErrorIcon
		} else {
			c.Status(http.StatusNoContent)
			FooterIcon = member.User.AvatarURL("40")
			username = member.User.String()
		}
	} else {
		c.Status(http.StatusAccepted)
		FooterIcon = config.Discord.EmbedErrorIcon
	}
	checkAccount(strings.ToLower(eventJson.Name))
	switch eventJson.Type {
	case "chat":
		if server.Config.Chat.Embed {
			messageEmbed := srvEmbed.Chat(*eventJson, *server.Config, FooterIcon, username, s)
			for _, channelID := range server.Config.Chat.ChannelID {
				_, err = s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
				}
			}
		} else {
			for _, channelID := range server.Config.Chat.ChannelID {
				_, err = s.ChannelMessageSend(channelID, fmt.Sprintf("%v%v %v", eventJson.Name, server.Config.Chat.Prefix, eventJson.Value))
				if err != nil {
					log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
				}
			}
		}

	case "death":
		if server.Config.SRV.Events.Death {
			messageEmbed := srvEmbed.PlayerDeath(*eventJson, *server.Config, FooterIcon, username, s)
			for _, channelID := range server.Config.SRV.ChannelID {
				_, err = s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Death embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	case "advancement":
		if server.Config.SRV.Events.Advancement {
			messageEmbed := srvEmbed.PlayerAdvancement(*eventJson, *server.Config, FooterIcon, username, s)
			for _, channelID := range server.Config.SRV.ChannelID {
				_, err = s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Advancement embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	case "join":
		if server.Config.SRV.Events.Join {
			server.OnlinePlayers.Mu.Lock()
			defer server.OnlinePlayers.Mu.Unlock()
			name := strings.ToLower(eventJson.Name)
			server.OnlinePlayers.Players = append(server.OnlinePlayers.Players, &name)
			messageEmbed := srvEmbed.PlayerJoin(*server.Config, strings.ToLower(eventJson.Name), FooterIcon, username, s)
			for _, channelID := range server.Config.SRV.ChannelID {
				_, err = s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Join embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	case "quit":
		if server.Config.SRV.Events.Quit {
			server.OnlinePlayers.Mu.Lock()
			defer server.OnlinePlayers.Mu.Unlock()
			for i, player := range server.OnlinePlayers.Players {
				if *player == strings.ToLower(eventJson.Name) {
					server.OnlinePlayers.Players = append(server.OnlinePlayers.Players[:i], server.OnlinePlayers.Players[i+1:]...)
					break
				}
			}
			messageEmbed := srvEmbed.PlayerQuit(*server.Config, strings.ToLower(eventJson.Name), FooterIcon, username, s)
			for _, channelID := range server.Config.SRV.ChannelID {
				_, err = s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Quit embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	}
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
