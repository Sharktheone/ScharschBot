package listeners

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/srv"
	"context"
)

func StatusListener(ctx context.Context, server *conf.Server, data chan *pterodactyl.ChanData) {
	for {
		select {
		case d := <-data:
			if d.Event == pterodactyl.WebsocketStatus {
				srv.HandlePower(d.Data.State, server)
			} else {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}
