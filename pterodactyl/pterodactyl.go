package pterodactyl

import (
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type ServerState struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	ID     string `json:"id"`
}
type ServerStat struct {
	Name    string  `json:"name"`
	Ram     int     `json:"memory_bytes"`
	RamMax  int     `json:"memory_limit_bytes"`
	Cpu     float64 `json:"cpu_absolute"`
	Network struct {
		Rx int `json:"rx_bytes"`
		Tx int `json:"tx_bytes"`
	}
	Status string `json:"status"`
	Disk   int    `json:"disk_bytes"`
	Uptime int    `json:"uptime"`
}

var (
	config        = conf.GetConf()
	panelUrl      = config.Pterodactyl.PanelURL
	apiKey        = fmt.Sprintf("Bearer %s", config.Pterodactyl.APIKey)
	websocketAuth = false
	ServerStates  []*ServerState
	ServerStats   []*ServerStat
)

//goland:noinspection GoUnusedConst
const (
	AuthSuccess   = "auth success"
	Status        = "status"
	ConsoleOutput = "console output"
	Stats         = "stats"
	TokenExpiring = "token expiring"
	TokenExpired  = "token expired"
)

func SendCommand(command string, serverID string) (successful bool) {

	var (
		url         = fmt.Sprintf("%s/api/client/servers/%s/command", panelUrl, serverID)
		payloadJson = fmt.Sprintf("{\"command\": \"%v\"}", strings.ReplaceAll(command, `"`, `\"`))
		payload     = strings.NewReader(payloadJson)
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}

func Start(serverID string) (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"start\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}

func Stop(serverID string) (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"stop\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}

//goland:noinspection GoUnusedExportedFunction
func Kill(serverID string) (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"kill\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}

func Restart(serverID string) (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"restart\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}
func getWebsocket(serverID string) (socket *websocket.Conn, successful bool) {
	var (
		url = fmt.Sprintf("%s/api/client/servers/%s/websocket", panelUrl, serverID)
	)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	resSuccessful := false
	if res != nil {
		resSuccessful = res.StatusCode == 200
	} else {
		log.Println("res is empty, maybe wrong url, or server is offline or invalid certificate")
		return nil, false
	}
	var response struct {
		Data struct {
			Token  string `json:"token"`
			Socket string `json:"socket"`
		} `json:"data"`
	}
	if resSuccessful {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&response)
		if err != nil {
			log.Printf("Failed to decode response: %v", err)
		}
	}
	if !resSuccessful {
		log.Printf("Failed to get websocket info: %v", res.Status)
		return nil, false
	}
	conn, resp, err := websocket.DefaultDialer.Dial(response.Data.Socket, nil)
	if err != nil {
		if resp != nil {
			log.Printf("Pterodactyl: Failed to connect to websocket: %v status: %v", err, resp.Status)
		} else {
			log.Printf("Pterodactyl: Failed to connect to websocket: %v", err)
		}
		return nil, false
	}
	auth := fmt.Sprintf(`{"event":"auth","args":["%v"]}`, response.Data.Token)
	err = conn.WriteMessage(websocket.TextMessage, []byte(auth))
	if err != nil {
		log.Printf("Failed to send auth: %v", err)
		return
	}
	websocketAuth = true
	return conn, true
}

func Websocket(serverID string, event string, callback func([]string, string), callbackLines int, callbackTime time.Duration, sendOnlyNew bool, doStats bool) {
	var (
		websocketConn *websocket.Conn
		serverConf    = conf.GetServerConf(serverID, "")
	)
	if !websocketAuth {
		var successful bool
		websocketConn, successful = getWebsocket(serverID)
		if !successful {
			return
		}
		err := websocketConn.WriteMessage(websocket.TextMessage, []byte(`{"event":"send logs","args":[null]}`))
		if err != nil {
			log.Printf("Failed to send log request: %v", err)
			return
		}
	}
	var result struct {
		Event string   `json:"event"`
		Args  []string `json:"args"`
	}
	go func() {
		var (
			lines        = 0
			timer        *time.Timer
			doneCallback = true
			output       []string
		)
		for {
			err := websocketConn.ReadJSON(&result)
			if err != nil {
				log.Printf("Failed to read message: %v", err)
				return
			}
			if result.Event == "token expiring" || result.Event == "token expired" || result.Event == "jwt error" {
				newConn, successful := getWebsocket(serverID)
				if !successful {
					var failed = 0
					for successful == false && failed <= 5 {
						newConn, successful = getWebsocket(serverID)
						failed++
					}
				} else {
					websocketConn = newConn
				}
			}
			var prevState = ""
			if result.Event == "status" {
				for _, Server := range ServerStates {
					if Server.Name == serverConf.ServerName {
						prevState = Server.Status
					}
				}
				for i, Server := range ServerStates {
					if Server.Name == serverConf.ServerName && doStats {
						if len(ServerStates) > i+1 {
							ServerStates = append(ServerStates[:i], ServerStates[i+1:]...)
						}
					}
				}
				ServerStates = append(ServerStates, &ServerState{
					Name:   serverConf.ServerName,
					Status: result.Args[0],
					ID:     serverConf.ServerID,
				})
			}
			if result.Event == "stats" && doStats {
				for i, Server := range ServerStats {
					if Server.Name == serverConf.ServerName {
						ServerStats = append(ServerStats[:i], ServerStats[i+1:]...)

					}
				}
				decoder := json.NewDecoder(strings.NewReader(result.Args[0]))
				var newStats ServerStat
				err = decoder.Decode(&newStats)
				if err != nil {
					log.Printf("Failed to decode stats: %v", err)
				}
				newStats.Name = serverConf.ServerName
				ServerStats = append(ServerStats, &newStats)
			}
			if result.Event == event {
				if doneCallback && callbackTime != 0 {
					timer = time.NewTimer(callbackTime)
					doneCallback = false
					go func() {
						<-timer.C
						callback(output, serverID)
						doneCallback = true
					}()
				}
				if callbackLines == 0 && callbackTime == 0 {
					if sendOnlyNew && result.Event == "status" {
						if prevState != result.Args[0] || prevState == "" {
							if callback != nil {
								go callback(result.Args, serverID)
							}
						}
					} else {
						if callback != nil {
							go callback(result.Args, serverID)
						}
					}
				} else {
					if lines == callbackLines {
						callback(output, serverID)
						timer.Stop()
						doneCallback = true
						lines = 0
						output = nil
					} else {
						if len(result.Args) > 0 {
							regex := regexp.MustCompile(config.Pterodactyl.RegexRemoveAnsi)
							output = append(output, regex.ReplaceAllString(result.Args[0], ""))
							lines++
						}
					}
				}

			}
			result.Args = nil
			result.Event = ""
		}
	}()
}
