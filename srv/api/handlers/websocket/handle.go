package websocket

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/embed/srvEmbed"
	"Scharsch-Bot/types"
	"context"
)

var (
	config = conf.GetConf()
)

func (p *PSRVEvent) processEvent() {
	ctx := &p.h.ctx
	e := p.e
	if p.h.authenticated == false && e.Event != Auth {
		return
	}
	switch e.Event {
	case SendPlayers:
		p.sendPlayers(ctx, e)
	case KickPlayer:
		p.kickPlayer(ctx, e)
	case BanPlayer:
		p.banPlayer(ctx, e)
	case UnbanPlayer:
		p.unbanPlayer(ctx, e)
	case PlayerJoined:
		p.playerJoined(ctx, e)
	case PlayerLeft:
		p.playerLeft(ctx, e)
	case Players:
		p.players(ctx, e)
	case ChatMessage:
		p.chatMessage(ctx, e)
	case PlayerDeath:
		p.playerDeath(ctx, e)
	case PlayerAdvancement:
		p.playerAdvancement(ctx, e)
	case Auth:
		p.auth(ctx, e)
	}
}

// sendPlayers send total online players to server
func (p *PSRVEvent) sendPlayers(ctx *context.Context, e *types.WebsocketEvent) {

}

// kickPlayer kick player on all servers
func (p *PSRVEvent) kickPlayer(ctx *context.Context, e *types.WebsocketEvent) {

}

// banPlayer ban player on all servers
func (p *PSRVEvent) banPlayer(ctx *context.Context, e *types.WebsocketEvent) {

}

func (p *PSRVEvent) unbanPlayer(ctx *context.Context, e *types.WebsocketEvent) {

}

func (p *PSRVEvent) playerJoined(ctx *context.Context, e *types.WebsocketEvent) {
	if p.h.server.Config.SRV.Events.PlayerJoinLeft {
		*p.h.server.OnlinePlayers.Players = append(*p.h.server.OnlinePlayers.Players, e.Data.Player)
	}
}

func (p *PSRVEvent) playerLeft(ctx *context.Context, e *types.WebsocketEvent) {
	if p.h.server.Config.SRV.Events.PlayerJoinLeft {
		*p.h.server.OnlinePlayers.Players = append(*p.h.server.OnlinePlayers.Players, e.Data.Player)
	}
}

func (p *PSRVEvent) players(ctx *context.Context, e *types.WebsocketEvent) {
	if p.h.server.Config.SRV.Events.PlayerJoinLeft {
		p.h.server.OnlinePlayers.Mu.Lock()
		defer p.h.server.OnlinePlayers.Mu.Unlock()
		p.h.server.OnlinePlayers.Players = &e.Data.Players
	}
}

func (p *PSRVEvent) chatMessage(ctx *context.Context, e *types.WebsocketEvent) {

}

func (p *PSRVEvent) playerDeath(ctx *context.Context, e *types.WebsocketEvent) {
	if p.h.server.Config.SRV.Events.Death {
		messageEmbed := srvEmbed.PlayerDeath(p.e, p.h.server.Config, p.footerIcon, p.username, p.session)
		p.session.SendEmbeds(p.h.server.Config.SRV.ChannelID, messageEmbed, "Death")
	}

}

func (p *PSRVEvent) playerAdvancement(ctx *context.Context, e *types.WebsocketEvent) {
	if p.h.server.Config.SRV.Events.Advancement {
		messageEmbed := srvEmbed.PlayerAdvancement(p.e, p.h.server.Config, p.footerIcon, p.username, p.session)
		p.session.SendEmbeds(p.h.server.Config.SRV.ChannelID, &messageEmbed, "Advancement")
	}
}

func (p *PSRVEvent) auth(ctx *context.Context, e *types.WebsocketEvent) {
	if e.Data.Password == config.SRV.API.Password && e.Data.User == config.SRV.API.User {
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
