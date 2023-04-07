package websocket

import (
	"Scharsch-Bot/conf"
	"context"
)

var (
	config = conf.GetConf()
)

func (s *Handler) processEvent(ctx *context.Context, e *Event) {
	if s.authenticated == false && e.Event != Auth {
		return
	}
	switch e.Event {
	case SendPlayers:
		s.sendPlayers(ctx, e)
	case KickPlayer:
		s.kickPlayer(ctx, e)
	case BanPlayer:
		s.banPlayer(ctx, e)
	case UnbanPlayer:
		s.unbanPlayer(ctx, e)
	case PlayerJoined:
		s.playerJoined(ctx, e)
	case PlayerLeft:
		s.playerLeft(ctx, e)
	case Players:
		s.players(ctx, e)
	case ChatMessage:
		s.chatMessage(ctx, e)
	case PlayerDeath:
		s.playerDeath(ctx, e)
	case PlayerAdvancement:
		s.playerAdvancement(ctx, e)
	case Auth:
		s.auth(ctx, e)
	}
}

func (s *Handler) sendPlayers(ctx *context.Context, e *Event) {

}

func (s *Handler) kickPlayer(ctx *context.Context, e *Event) {

}

func (s *Handler) banPlayer(ctx *context.Context, e *Event) {

}

func (s *Handler) unbanPlayer(ctx *context.Context, e *Event) {

}

func (s *Handler) playerJoined(ctx *context.Context, e *Event) {

}

func (s *Handler) playerLeft(ctx *context.Context, e *Event) {

}

func (s *Handler) players(ctx *context.Context, e *Event) {

}

func (s *Handler) chatMessage(ctx *context.Context, e *Event) {

}

func (s *Handler) playerDeath(ctx *context.Context, e *Event) {

}

func (s *Handler) playerAdvancement(ctx *context.Context, e *Event) {

}

func (s *Handler) auth(ctx *context.Context, e *Event) {
	if e.Data.Password == config.SRV.API.Password && e.Data.User == config.SRV.API.User {
		s.authenticated = true
		s.send <- &Event{
			Event: AuthSuccess,
		}
	} else {
		s.authenticated = false
		s.send <- &Event{
			Event: AuthFailed,
		}
	}
}
