package websocket

func (s *server) listen() {
	for {
		var data Event
		if err := s.conn.ReadJSON(&data); err == nil {
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

func (s *server) sendLoop() {
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
