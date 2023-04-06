package websocket

import "context"

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
	case SendCommand:
		s.sendCommand(ctx, e)
	case SendChatMessage:
		s.sendChatMessage(ctx, e)
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
	case AuthSuccess:
		s.authSuccess(ctx, e)
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

func (s *Handler) sendCommand(ctx *context.Context, e *Event) {

}

func (s *Handler) sendChatMessage(ctx *context.Context, e *Event) {

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

}

func (s *Handler) authSuccess(ctx *context.Context, e *Event) {

}
