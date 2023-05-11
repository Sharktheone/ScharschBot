package pprof

import (
	"fmt"
	"github.com/Sharktheone/ScharschBot/flags"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var (
	enabled = flags.Bool("pprof")
	port    = flags.Int("pprof-port")
)

func Start() {
	if *enabled {
		go func() {
			log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
		}()
	}
}
