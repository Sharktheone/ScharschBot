package pterodactyl

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"net/http"
	"strings"
)

var (
	config   = conf.GetConf()
	serverID = config.Pterodactyl.ServerID
	panelUrl = config.Pterodactyl.PanelURL
	apiKey   = fmt.Sprintf("Bearer %s", config.Pterodactyl.APIKey)
)

func SendCommand(command string) (successful bool) {
	var (
		url = fmt.Sprintf("%s/api/client/servers/%s/command", panelUrl, serverID)

		payloadJson = fmt.Sprintf("{\n\t\"command\": \"%v\"\n}", command)
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

func Start() (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"start\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}

func Stop() (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"stop\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}
func Kill() (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"kill\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}
func Restart() (successful bool) {
	var (
		url     = fmt.Sprintf("%s/api/client/servers/%s/power", panelUrl, serverID)
		payload = strings.NewReader("{\n\t\"signal\": \"restart\"\n}")
	)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)
	resSuccessful := res.StatusCode == 204

	return resSuccessful
}
