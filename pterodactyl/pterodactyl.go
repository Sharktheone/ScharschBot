package pterodactyl

import (
	"Scharsch-Bot/conf"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	url2 "net/url"
	"strings"
)

//goland:noinspection GoUnusedConst
const (
	WebsocketAuthSuccess   = "auth success"
	WebsocketStatus        = "status"
	WebsocketConsoleOutput = "console output"
	WebsocketStats         = "stats"
	WebsocketTokenExpiring = "token expiring"
	WebsocketTokenExpired  = "token expired"

	PowerSignalStart   = "start"
	PowerSignalStop    = "stop"
	PowerSignalKill    = "kill"
	PowerSignalRestart = "restart"
)

type Server struct {
	server    *conf.Server
	listeners []func(server *conf.Server, data chan string)
	data      chan string
	socket    *websocket.Conn
}

func New(server *conf.Server) *Server {
	return &Server{
		server:    server,
		listeners: []func(serverID *conf.Server, data chan string){},
		data:      make(chan string),
	}
}

func (s *Server) SendCommand() {

}

func (s *Server) Start() error {
	return s.Power(PowerSignalStart)
}

func (s *Server) Stop() error {
	return s.Power(PowerSignalStop)
}

func (s *Server) Kill() error {
	return s.Power(PowerSignalKill)
}

func (s *Server) Restart() error {
	return s.Power(PowerSignalRestart)
}

func (s *Server) Power(signal string) error {
	var (
		url, _  = url2.JoinPath(panelUrl, fmt.Sprintf("/api/client/servers/%s/power", s.server.ServerID))
		payload = strings.NewReader(fmt.Sprintf(`{"signal": "%s"}`, signal))
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 204 {
		return fmt.Errorf("could not send power signal to server %s. Failed with %s", s.server.ServerID, res.Status)
	}
	return nil
}

func (s *Server) getWebsocket() {

}

func (s *Server) Websocket() { // start the websocket

}

func (s *Server) AddListener(listener func(server *conf.Server, data chan string)) {
	s.listeners = append(s.listeners, listener)
	go listener(s.server, s.data)
}
