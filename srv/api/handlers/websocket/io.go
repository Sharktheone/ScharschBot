package websocket

import (
	"Scharsch-Bot/srv/playersrv"
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
			go h.handleEvents(data)
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

func (h *Handler) handleEvents(data *types.WebsocketEvent) {
	for {
		select {
		case <-h.ctx.Done():
			return
		default:
			h.receive <- data
			pSRV, err := playersrv.DecodePlayer(data, h.server)
			if err != nil {
				h.send <- &types.WebsocketEvent{
					Event: Error,
					Data: types.WebsocketEventData{
						Error: err.Error(),
					},
				}
			}
			if pSRV != nil {
				pSRV.SwitchEvents()
			}

		}
	}
}
