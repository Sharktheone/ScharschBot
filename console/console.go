package console

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var config = conf.GetConf()

//goland:noinspection GoUnusedParameter
func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	allowed := false
	for _, server := range config.Pterodactyl.Servers {
		if server.Console.Reverse {
			for _, neededChannelID := range server.Console.ChannelID {
				if neededChannelID == m.ChannelID {
					command := strings.SplitAfter(m.Message.Content, server.Console.ReversePrefix)
					if command[0] == server.Console.ReversePrefix {
						for _, role := range m.Member.Roles {
							for _, neededRole := range config.Whitelist.Roles.ServerRoleID {
								if role == neededRole {
									allowed = true
									break
								}
							}
						}
						if allowed {
							var commandString string
							for _, element := range command[1:] {
								commandString += element
							}
							log.Printf("%v is sending command to server %v: %v", m.Author.ID, server.ServerName, commandString)
							pterodactyl.SendCommand(commandString, server.ServerID)
						}
					}
					break
				}
			}
		}

	}
}

func ChatHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	allowed := false
	for _, server := range config.Pterodactyl.Servers {
		if server.Chat.Reverse {
			for _, neededChannelID := range server.Chat.ChannelID {
				if neededChannelID == m.ChannelID && m.Author.ID != s.State.User.ID {
					for _, role := range m.Member.Roles {
						for _, neededRole := range config.Whitelist.Roles.ServerRoleID {
							if role == neededRole {
								allowed = true
								break
							}
						}
					}
					if allowed {
						message := fmt.Sprintf(" %v: %v", m.Author.Username, m.Message.Content)
						command := fmt.Sprintf(config.Pterodactyl.ChatCommand, message)
						pterodactyl.SendCommand(command, server.ServerID)
					}
				}

			}
		}
	}
}
