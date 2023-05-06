package listeners

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/srv/serversrv"
	"context"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func StatsListener(ctx context.Context, server *conf.Server, stats chan *pterodactyl.ChanData) {
	var (
		status *pterodactyl.ServerStatus
	)
	c := cron.New()
	if err := c.AddFunc(fmt.Sprintf("@every %v", server.ChannelInfo.Interval), func() {
		if status != nil {
			serversrv.ChannelStats(status, server)
		}
	}); err != nil {
		log.Panicf("Failed to add cron job: %v for server %v ", err, server.ServerID)
	}
	c.Start()
	for {
		select {
		case stat := <-stats:
			if stat.Event == pterodactyl.WebsocketStats {
				status = stat.Data
			}
		case <-ctx.Done():
			c.Stop()
			return
		}
	}
}
