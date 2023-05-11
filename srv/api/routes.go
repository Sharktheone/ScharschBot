package api

import (
	"fmt"
	"github.com/Sharktheone/ScharschBot/conf"
	"github.com/Sharktheone/ScharschBot/flags"
	"github.com/Sharktheone/ScharschBot/srv/api/handlers/websocket"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	config   = conf.Config
	user     = flags.StringWithFallback("apiUser", &config.SRV.API.User)
	password = flags.StringWithFallback("apiPassword", &config.SRV.API.Password)
	port     = flags.IntWithFallback("apiPort", &config.SRV.API.Port)
)

func Start() {
	log.Printf("Starting http server on port %v", *port)
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	//r.Use(gin.Recovery(), gin.BasicAuth(gin.Accounts{
	//	*user: *password,
	//}))
	// TODO: Enable it again, when BasicAuth in Plugins
	r.GET("/:serverID/ws", websocket.ServerWS)

	if err := r.Run(fmt.Sprintf(":%v", *port)); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
	log.Println("Started http server")
}
