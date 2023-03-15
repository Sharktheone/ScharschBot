package pterodactyl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
)

type eventType struct {
	Event string   `json:"event"`
	Args  []string `json:"args"`
}

func (s *Server) connectWS() error {
	res, err := request(fmt.Sprintf("/api/client/servers/%s/websocket", s.server.ServerID), "GET", nil)
	if err != nil {
		return fmt.Errorf("could not connect to websocket: %w", err)
	}
	if res != nil {
		if res.StatusCode == 200 {
			var socketInfo struct {
				Data struct {
					Token  string `json:"token"`
					Socket string `json:"socket"`
				} `json:"data"`
			}
			if err := json.NewDecoder(res.Body).Decode(&socketInfo); err != nil {
				return fmt.Errorf("failed to decode pterodactyl websocket information for server %v: %s",
					s.server.ServerName, err)
			}
			var auth bytes.Buffer
			if err := json.NewEncoder(&auth).Encode(eventType{
				Event: "auth",
				Args:  []string{socketInfo.Data.Token},
			}); err != nil {
				return fmt.Errorf("failed to encode pterodactyl websocket auth for server %v: %s", s.server.ServerName, err)
			}
			if !s.connected {
				s.socket, _, err = websocket.DefaultDialer.Dial(socketInfo.Data.Socket, nil)
				if err != nil {
					return fmt.Errorf("failed to connect to pterodactyl websocket for server %v: %s", s.server.ServerName, err)
				}
			}
			if err := s.socket.WriteMessage(websocket.TextMessage, auth.Bytes()); err != nil {
				return fmt.Errorf("failed to send auth to pterodactyl websocket for server %v: %s", s.server.ServerName, err)
			}
			var (
				event eventType
			)
			if err := s.socket.ReadJSON(event); err != nil {
				fmt.Printf("failed to read websocket message: %s", err)
				return err
			}
			if event.Event == WebsocketAuthSuccess {
				return nil
			} else {
				return fmt.Errorf("failed to authenticate to pterodactyl websocket for server %v: %s", s.server.ServerName, err)
			}

		} else {
			return fmt.Errorf("could not connect to websocket: %v", res.Status)
		}
	} else {
		return fmt.Errorf("cannot reach pterodactyl instance with panel url %v", panelUrl)
	}
}

func (s *Server) Listen() error {
	if !s.connected {
		if err := s.connectWS(); err != nil {
			return err
		} else {
			s.connected = true
		}
	}
	for {
		var (
			event eventType
		)
		if err := s.socket.ReadJSON(&event); err != nil {
			fmt.Printf("failed to read websocket message: %s", err)
			continue
		}
		if event.Event == WebsocketTokenExpired || event.Event == WebsocketTokenExpiring {
			if event.Event == WebsocketTokenExpired {
				s.connected = false
			}
			if err := s.connectWS(); err != nil {
				var tries int
				for tries < 5 && err != nil {
					if err := s.connectWS(); err != nil {
						tries++
					}
				}
				if err != nil {
					return err
				}
			}
			continue
		}
		s.setStats(&event)
	}
}

func (s *Server) setStats(data *eventType) {

}
