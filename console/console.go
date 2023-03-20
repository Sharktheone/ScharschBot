package console

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/discordMember"
	"Scharsch-Bot/pterodactyl"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var config = conf.GetConf()

//goland:noinspection GoUnusedParameter
func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, server := range conf.Config.Pterodactyl.Servers {
		if server.Console.Reverse {
			for _, neededChannelID := range server.Console.ChannelID {
				if neededChannelID == m.ChannelID {
					command := strings.SplitAfter(m.Message.Content, server.Console.ReversePrefix)
					if command[0] == server.Console.ReversePrefix {
						if discordMember.HasRole(m.Member, config.Whitelist.Roles.ServerRoleID) {
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
	for _, server := range config.Pterodactyl.Servers {
		if server.Chat.Reverse {
			for _, neededChannelID := range server.Chat.ChannelID {
				if neededChannelID == m.ChannelID && m.Author.ID != s.State.User.ID {
					if discordMember.HasRole(m.Member, config.Whitelist.Roles.ServerRoleID) {
						message := fmt.Sprintf(" %v: %v", m.Author.Username, m.Message.Content)
						command := fmt.Sprintf(config.Pterodactyl.ChatCommand, message)
						pterodactyl.SendCommand(command, server.ServerID)
					}
				}

			}
		}
	}
}
