package websocket

import "github.com/fasthttp/websocket"

func (s *Handler) handleInbound() {
	for {
		var data Event
		if err := s.conn.ReadJSON(&data); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived, websocket.CloseServiceRestart) {
				return
			}
		} else {
			s.receive <- data
		}
		select {
		case <-s.ctx.Done():
			return
		default:
			continue
		}
	}
}

func (s *Handler) handleOutbound() {
	for {
		select {
		case data := <-s.send:
			if err := s.conn.WriteJSON(data); err != nil {
				return
			}
		case <-s.ctx.Done():
			return
		}
	}
}
