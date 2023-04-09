package websocket

import (
	"Scharsch-Bot/types"
	"github.com/fasthttp/websocket"
)

func (h *Handler) handleInbound() {
	for {
		var data *types.WebsocketEvent
		if err := h.conn.ReadJSON(&data); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived, websocket.CloseServiceRestart) {
				return
			}
		} else {
			h.receive <- data
		}
		select {
		case <-h.ctx.Done():
			return
		default:
			continue
		}
	}
}

func (h *Handler) handleOutbound() {
	for {
		select {
		case data := <-h.send:
			if err := h.conn.WriteJSON(data); err != nil {
				return
			}
		case <-h.ctx.Done():
			return
		}
	}
}
