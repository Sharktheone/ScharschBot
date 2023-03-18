package pterodactyl

import (
	"Scharsch-Bot/conf"
	"fmt"
	"github.com/fasthttp/websocket"
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

type ServerStatus struct {
	State   string  `json:"state"`
	Ram     int     `json:"memory_bytes"`
	RamMax  int     `json:"memory_limit_bytes"`
	Cpu     float64 `json:"cpu_absolute"`
	Network struct {
		Rx int `json:"rx_bytes"`
		Tx int `json:"tx_bytes"`
	} `json:"network"`
	Disk   int `json:"disk_bytes"`
	Uptime int `json:"uptime"`
}

type Server struct {
	server    *conf.Server
	data      chan *ServerStatus
	console   chan string
	status    ServerStatus
	socket    *websocket.Conn
	connected bool
}

func New(server *conf.Server) *Server {
	return &Server{
		server: server,
		data:   make(chan *ServerStatus),
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
		payload = strings.NewReader(fmt.Sprintf(`{"signal": "%s"}`, signal))
	)
	res, err := request("POST", fmt.Sprintf("/api/client/servers/%s/power", s.server.ServerID), payload)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		return fmt.Errorf("could not send power signal to server %s. Failed with %s", s.server.ServerID, res.Status)
	}
	return nil
}

func (s *Server) AddListener(listener func(server *conf.Server, data chan *ServerStatus), event string) {
	go listener(s.server, s.data)
}
func (s *Server) AddConsoleListener(listener func(server *conf.Server, console chan string)) {
	go listener(s.server, s.console)
}
