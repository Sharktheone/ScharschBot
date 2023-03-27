package api

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/srv/api/handlers"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	config   = conf.GetConf()
	user     = flag.String("apiUser", config.SRV.API.User, "Username for the API")
	password = flag.String("apiPassword", config.SRV.API.Password, "Password for the API")
	port     = flag.Int("apiPort", config.SRV.API.Port, "Port for the API")
)

func init() {
	flag.Parse()
}

func Start() {
	log.Printf("Starting http server on port %v", *port)
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Recovery(), gin.BasicAuth(gin.Accounts{
		*user: *password,
	}))

	r.POST("/", handlers.PlayerSRVEventHandler)

	if err := r.Run(fmt.Sprintf(":%v", *port)); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
	log.Println("Started http server")
}
