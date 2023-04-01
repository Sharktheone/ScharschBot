package websocket

import (
	"Scharsch-Bot/pterodactyl"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/gin-gonic/gin"
	"net/http"
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
	conn   *websocket.Conn
	server *pterodactyl.Server
}

var (
	servers  = make(map[*websocket.Conn]*server)
	upgrader = websocket.Upgrader{}
)

func Handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to upgrade connection: %v", err),
		})
		return
	}
	defer conn.Close()
	s, err := pterodactyl.GetServer(c.Param("serverID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to get server: %v", err),
		})
		return
	}
	servers[conn] = &server{
		conn:   conn,
		server: s,
	}
}
