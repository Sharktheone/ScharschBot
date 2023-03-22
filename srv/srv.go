package srv

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/pterodactyl/listeners"
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start() {
	log.Printf("Starting http server on port %v", port)
	//TODO: replace with gin or fiber
	http.HandleFunc("/", playerSRVEventHandler)
	addr := fmt.Sprintf(":%v", port)
	var err error
	go func() { err = http.ListenAndServe(addr, nil) }()
	if err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
	log.Println("Started http server")

	for _, server := range config.Pterodactyl.Servers {
		go func(server conf.Server) {
			ctx := context.Background()
			s := pterodactyl.New(&ctx, &server)
			if server.Console.Enabled {
				s.AddConsoleListener(func(server *conf.Server, console chan string) {
					listeners.ConsoleListener(ctx, server, console, nil)
				})
			}
			if server.StateMessages.Enabled {
				s.AddListener(func(ctx *context.Context, server *conf.Server, data chan *pterodactyl.ChanData) {
					listeners.StatusListener(*ctx, server, data)
				}, server.ServerID+"_stateMessages")
			}
			if server.ChannelInfo.Enabled {
				s.AddListener(func(ctx *context.Context, server *conf.Server, data chan *pterodactyl.ChanData) {
					listeners.StatsListener(*ctx, server, data)
				}, server.ServerID+"_channelInfo")
			}
			if err := s.Listen(); err != nil {
				log.Printf("Error while listening to server %v: %v", server.ServerID, err)
			}

		}(server)
	}
}
