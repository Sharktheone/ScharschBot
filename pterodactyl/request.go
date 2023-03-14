package pterodactyl

import (
	"net/http"
	url2 "net/url"
	"strings"
)

func request(path string, method string, payload *strings.Reader) (*http.Response, error) {
	var (
		url, err = url2.JoinPath(panelUrl, path)
	)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", apiKey)
	return http.DefaultClient.Do(req)
}
