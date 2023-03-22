package listeners

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/srv/serversrv"
	"context"
)

func StatusListener(ctx context.Context, server *conf.Server, data chan *pterodactyl.ChanData) {
	for {
		select {
		case d := <-data:
			if d.Event == pterodactyl.WebsocketStatus {
				serversrv.HandlePower(d.Data.State, server)
			} else {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}
