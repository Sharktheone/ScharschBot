package websocket

import (
	"Scharsch-Bot/discord/session"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/types"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
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
	ReportPlayer      = "reportPlayer"
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
	Error             = "error"
)

type Handler struct {
	conn          *websocket.Conn
	server        *pterodactyl.Server
	uuid          uuid.UUID
	send          chan *types.WebsocketEvent
	receive       chan *types.WebsocketEvent
	ctx           context.Context
	cancel        context.CancelFunc
	authenticated bool
}

type PSRVEvent struct {
	h           *Handler
	e           *types.WebsocketEvent
	userID      *string
	onWhitelist *bool
	footerIcon  *string
	username    *string
	member      *discordgo.Member
	session     *session.Session
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
	handler, err := getWSHandler(s, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to get server: %v", err),
		})
		return
	}

	go handler.handleInbound()
	go handler.handleOutbound()
}

func getWSHandler(s *pterodactyl.Server, c *gin.Context) (*Handler, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to upgrade connection: %v", err)
	}
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Handler{
		conn:    conn,
		server:  s,
		uuid:    u,
		ctx:     ctx,
		cancel:  cancel,
		send:    make(chan *types.WebsocketEvent),
		receive: make(chan *types.WebsocketEvent),
	}, nil
}
