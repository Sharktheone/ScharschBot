package srv

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"log"
)

func ConsoleSrv(console []string, serverID string) {
	if console == nil {
		return
	}
	serverConf := conf.GetServerConf(serverID, "")
	var message string
	for _, line := range console {
		message += fmt.Sprintf("\n%v", line)
	}
	for _, channelID := range serverConf.Console.ChannelID {
		_, err := Session.ChannelMessageSend(channelID, fmt.Sprintf("```%v```", message))
		if err != nil {
			log.Printf("Failed to send console to discord: %v (ChannelID: %v)", err, channelID)
		}
	}
}
