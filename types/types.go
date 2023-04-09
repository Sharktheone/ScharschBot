package types

type EventJson struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Type   string `json:"type"`
	Server string `json:"server"`
}

type WebsocketEvent struct {
	Event string `json:"event"`
	Data  struct {
		Players      []string `json:"players"`
		Player       string   `json:"player"`
		Reason       string   `json:"reason"`
		Command      string   `json:"command"`
		Message      string   `json:"message"`
		DeathMessage string   `json:"death_message"`
		Advancement  string   `json:"advancement"`
		Password     string   `json:"password"`
		User         string   `json:"user"`
	}
}
