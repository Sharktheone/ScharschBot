package websocket

import (
	"Scharsch-Bot/pterodactyl"
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	AuthSuccess       = "authSuccess"
	AuthFailed        = "authFailed"
)

type Handler struct {
	conn          *websocket.Conn
	server        *pterodactyl.Server
	uuid          uuid.UUID
	send          chan *Event
	receive       chan *Event
	ctx           context.Context
	authenticated bool
}

var (
	upgrader = websocket.Upgrader{}
)

func ServerWS(c *gin.Context) {
	s, err := pterodactyl.GetServer(c.Param("serverID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid server id %v: %v", c.Param("serverID"), err),
		})
		return
	}
	handler, err := getWSHandler(s, c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to get server: %v", err),
		})
		return
	}
	defer handler.conn.Close()

	go handler.handleInbound()
	go handler.handleOutbound()

}

func getWSHandler(s *pterodactyl.Server, w http.ResponseWriter, r *http.Request) (*Handler, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade connection: %v", err)
	}
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid: %v", err)
	}
	return &Handler{
		conn:   conn,
		server: s,
		uuid:   u,
	}, nil
}
