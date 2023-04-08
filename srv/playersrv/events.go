package playersrv

import (
	"Scharsch-Bot/discord/embed/srvEmbed"
	"Scharsch-Bot/srv/api/handlers/websocket"
	"fmt"
	"log"
	"strings"
)

func (p *PlayerSrv) SwitchEvents() {
	switch p.eventJson.Type {
	case "chat":
		if p.server.Config.Chat.Embed {
			messageEmbed := srvEmbed.Chat(*p.eventJson, *p.server.Config, p.footerIcon, p.username, s)
			for _, channelID := range p.server.Config.Chat.ChannelID {
				_, err := s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
				}
			}
		} else {
			for _, channelID := range p.server.Config.Chat.ChannelID {
				_, err := s.ChannelMessageSend(channelID, fmt.Sprintf("%v%v %v", p.eventJson.Name, p.server.Config.Chat.Prefix, p.eventJson.Value))
				if err != nil {
					log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
				}
			}
		}

	case "death":
		if p.server.Config.SRV.Events.Death {
			messageEmbed := srvEmbed.PlayerDeath(*p.eventJson, *p.server.Config, p.footerIcon, p.username, s)
			for _, channelID := range p.server.Config.SRV.ChannelID {
				_, err := s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Death embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	case "advancement":
		if p.server.Config.SRV.Events.Advancement {
			// Temporary conversion to websocket.Event
			e := websocket.Event{}
			e.Data.Player = p.eventJson.Name
			e.Data.Advancement = p.eventJson.Value

			messageEmbed := srvEmbed.PlayerAdvancement(&e, p.server.Config, &p.footerIcon, &p.username, s)
			for _, channelID := range p.server.Config.SRV.ChannelID {
				_, err := s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Advancement embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	case "join":
		if p.server.Config.SRV.Events.PlayerJoinLeft {
			p.server.OnlinePlayers.Mu.Lock()
			defer p.server.OnlinePlayers.Mu.Unlock()
			name := strings.ToLower(p.eventJson.Name)
			*p.server.OnlinePlayers.Players = append(*p.server.OnlinePlayers.Players, name)
			messageEmbed := srvEmbed.PlayerJoin(*p.server.Config, strings.ToLower(p.eventJson.Name), p.footerIcon, p.username, s)
			for _, channelID := range p.server.Config.SRV.ChannelID {
				_, err := s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Join embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	case "quit":
		if p.server.Config.SRV.Events.PlayerJoinLeft {
			p.server.OnlinePlayers.Mu.Lock()
			defer p.server.OnlinePlayers.Mu.Unlock()
			for i, player := range *p.server.OnlinePlayers.Players {
				if player == strings.ToLower(p.eventJson.Name) {
					players := *p.server.OnlinePlayers.Players
					*p.server.OnlinePlayers.Players = append(players[:i], players[i+1:]...)
					break
				}
			}
			messageEmbed := srvEmbed.PlayerQuit(*p.server.Config, strings.ToLower(p.eventJson.Name), p.footerIcon, p.username, s)
			for _, channelID := range p.server.Config.SRV.ChannelID {
				_, err := s.ChannelMessageSendEmbed(channelID, &messageEmbed)
				if err != nil {
					log.Printf("Failed to send Quit embed: %v (channelID: %v)", err, channelID)
				}
			}
		}
	}
}
