package listeners

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/srv"
	"context"
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func StatsListener(ctx context.Context, server *conf.Server, stats chan *pterodactyl.ServerStatus) {
	var (
		status *pterodactyl.ServerStatus
	)
	c := cron.New()
	if err := c.AddFunc(fmt.Sprintf("@every %v", server.ChannelInfo.Interval), srv.ChannelStats(status, server)); err != nil {
		log.Panicf("Failed to add cron job: %v for server %v ", err, server.ServerID)
	}
	c.Start()
	for {
		select {
		case stat := <-stats:
			status = stat
		case <-ctx.Done():
			c.Stop()
			return
		}
	}
}
