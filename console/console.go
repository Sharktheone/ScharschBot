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

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	allowed := false
	for _, server := range config.Pterodactyl.Servers {
		if server.Console.Reverse {
			if server.Console.ChannelID == m.ChannelID {
				command := strings.SplitAfter(m.Message.Content, server.Console.ReversePrefix)
				if command[0] == server.Console.ReversePrefix {
					for _, role := range m.Member.Roles {
						if role == server.Console.ReverseRoleID {
							allowed = true
						}
						break
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

func ChatHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	allowed := false
	for _, server := range config.Pterodactyl.Servers {
		if server.Chat.Reverse {
			if server.Console.ChannelID == m.ChannelID && !m.Author.Bot {
				for _, role := range m.Member.Roles {
					if role == server.Console.ReverseRoleID {
						allowed = true
					}
					break
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
