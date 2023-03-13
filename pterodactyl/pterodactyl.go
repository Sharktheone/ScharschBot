package pterodactyl

import (
	"github.com/gorilla/websocket"
)

//goland:noinspection GoUnusedConst
const (
	AuthSuccess   = "auth success"
	Status        = "status"
	ConsoleOutput = "console output"
	Stats         = "stats"
	TokenExpiring = "token expiring"
	TokenExpired  = "token expired"
)

type Server struct {
	serverID  *string
	maxLines  *int
	listeners *[]func(serverID string, data chan string)
	data      chan string
	socket    *websocket.Conn
}

func (s *Server) SendCommand() {

}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Kill() {

}

func (s *Server) Restart() {

}

func (s *Server) getWebsocket() {

}

func (s *Server) Websocket() { // start the websocket

}

func (s *Server) AddListener() {

}
