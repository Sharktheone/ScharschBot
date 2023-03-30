package websocket

import (
	"github.com/fasthttp/websocket"
	"github.com/gin-gonic/gin"
)

// Bot => Server: SendPlayers
// Bot => Server: KickPlayer
// Bot => Server: BanPlayer
// Bot => Server: UnbanPlayer
// Bot => Server: SendCommand
// Bot => Server: SendChatMessage
// Server => Bot: PlayerJoined
// Server => Bot: PlayerLeft
// Server => Bot: Players
// Server => Bot: ChatMessage
// Server => Bot: PlayerDeath
// Server => Bot: PlayerAdvancement
// Server => Bot: Auth

const (
	SendPlayers       = "sendPlayers"
	KickPlayer        = "lickPlayer"
	BanPlayer         = "banPlayer"
	UnbanPlayer       = "unbanPlayer"
	SendCommand       = "sendCommand"
	SendChatMessage   = "sendChatMessage"
	PlayerJoined      = "playerJoined"
	PlayerLeft        = "playerLeft"
	Players           = "players"
	ChatMessage       = "chatMessage"
	PlayerDeath       = "playerDeath"
	PlayerAdvancement = "playerAdvancement"
	Auth              = "auth"
)

type server struct {
}

var (
	servers map[*websocket.Conn]*server
)

func init() {

}

func Handler(c *gin.Context) {

}
