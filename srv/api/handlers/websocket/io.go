package websocket

import (
	"github.com/Sharktheone/ScharschBot/types"
	"github.com/fasthttp/websocket"
	"log"
)

func (h *Handler) handleInbound() {
	for {
		var data *types.WebsocketEvent
		if err := h.conn.ReadJSON(&data); err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived, websocket.CloseServiceRestart) {
				h.cancel()
				if err := h.conn.Close(); err != nil {
					log.Printf("Failed to close websocket connection: %v", err)
				}
				return
			}
			log.Printf("Failed to read websocket message: %v", err)
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
	select {
	case h.receive <- data:
		break
	case <-h.ctx.Done():
		return
	default:
		break
	}
	pSRV, err := h.DecodePlayer(data)
	if err != nil {
		h.send <- &types.WebsocketEvent{
			Event: Error,
			Data: types.WebsocketEventData{
				Error: err.Error(),
			},
		}
	}
	if pSRV != nil {
		pSRV.processEvent()
	}
}
