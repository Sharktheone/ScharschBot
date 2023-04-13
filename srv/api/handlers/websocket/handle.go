package websocket

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/embed/srvEmbed"
	"Scharsch-Bot/types"
)

var (
	config = conf.GetConf()
)

func (p *PSRVEvent) processEvent() {
	if p.h.authenticated == false && p.e.Event != Auth {
		return
	}
	switch p.e.Event {
	case SendPlayers:
		p.sendPlayers()
	case KickPlayer:
		p.kickPlayer()
	case BanPlayer:
		p.banPlayer()
	case UnbanPlayer:
		p.unbanPlayer()
	case PlayerJoined:
		p.playerJoined()
	case PlayerLeft:
		p.playerLeft()
	case Players:
		p.players()
	case ChatMessage:
		p.chatMessage()
	case PlayerDeath:
		p.playerDeath()
	case PlayerAdvancement:
		p.playerAdvancement()
	case Auth:
		p.auth()
	}
}

// sendPlayers send total online players to server
func (p *PSRVEvent) sendPlayers() {

}

// kickPlayer kick player on all servers
func (p *PSRVEvent) kickPlayer() {

}

// banPlayer ban player on all servers
func (p *PSRVEvent) banPlayer() {

}

func (p *PSRVEvent) unbanPlayer() {

}

func (p *PSRVEvent) playerJoined() {
	if p.h.server.Config.SRV.Events.PlayerJoinLeft {
		p.h.server.OnlinePlayers.Players = append(p.h.server.OnlinePlayers.Players, &p.e.Data.Player)
	}
}

func (p *PSRVEvent) playerLeft() {
	if p.h.server.Config.SRV.Events.PlayerJoinLeft {
		p.h.server.OnlinePlayers.Players = append(p.h.server.OnlinePlayers.Players, &p.e.Data.Player)
	}
}

func (p *PSRVEvent) players() {
	if p.h.server.Config.SRV.Events.PlayerJoinLeft {
		p.h.server.OnlinePlayers.Mu.Lock()
		defer p.h.server.OnlinePlayers.Mu.Unlock()
		var players []*string
		for _, player := range p.e.Data.Players {
			players = append(players, &player)
		}
		p.h.server.OnlinePlayers.Players = players
	}
}

func (p *PSRVEvent) chatMessage() {
	if p.h.server.Config.Chat.Enabled {
		if p.h.server.Config.Chat.Embed {
			messageEmbed := srvEmbed.Chat(p.e, p.h.server.Config, p.footerIcon, p.username, p.session)
			p.session.SendEmbeds(p.h.server.Config.SRV.ChannelID, messageEmbed, "Chat")
		} else {

		}
	}
}

func (p *PSRVEvent) playerDeath() {
	if p.h.server.Config.SRV.Events.Death {
		messageEmbed := srvEmbed.PlayerDeath(p.e, p.h.server.Config, p.footerIcon, p.username, p.session)
		p.session.SendEmbeds(p.h.server.Config.SRV.ChannelID, messageEmbed, "Death")
	}

}

func (p *PSRVEvent) playerAdvancement() {
	if p.h.server.Config.SRV.Events.Advancement {
		messageEmbed := srvEmbed.PlayerAdvancement(p.e, p.h.server.Config, p.footerIcon, p.username, p.session)
		p.session.SendEmbeds(p.h.server.Config.SRV.ChannelID, &messageEmbed, "Advancement")
	}
}

func (p *PSRVEvent) auth() {
	if p.e.Data.Password == config.SRV.API.Password && p.e.Data.User == config.SRV.API.User {
		p.h.authenticated = true
		p.h.send <- &types.WebsocketEvent{
			Event: AuthSuccess,
		}
	} else {
		p.h.authenticated = false
		p.h.send <- &types.WebsocketEvent{
			Event: AuthFailed,
		}
	}
}
