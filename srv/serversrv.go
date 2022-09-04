package srv

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"log"
)

func startSrv() {
	// TODO
}

func ConsoleSrv(console []string, serverID string) {
	if console == nil {
		return
	}
	serverConf := conf.GetServerConf(serverID, "")
	var message string
	for _, line := range console {
		message += fmt.Sprintf("\n%v", line)
	}

	_, err := Session.ChannelMessageSend(serverConf.Console.ChannelID, fmt.Sprintf("```%v```", string(message)))
	if err != nil {
		log.Printf("Failed to send console to discord: %v", err)
	}
}
func channelStats() {
	// TODO
}
func serverStarting(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	_, err := Session.ChannelMessageSend(serverConf.Console.ChannelID, serverConf.StartMessage)
	if err != nil {
		log.Printf("Failed to send server start message to discord: %v", err)
	}
}
func serverStopping(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	_, err := Session.ChannelMessageSend(serverConf.Console.ChannelID, serverConf.StopMessage)
	if err != nil {
		log.Printf("Failed to send server stop message to discord: %v", err)
	}
}
func serverOnline(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	_, err := Session.ChannelMessageSend(serverConf.Console.ChannelID, serverConf.OnlineMessage)
	if err != nil {
		log.Printf("Failed to send server online message to discord: %v", err)
	}
}
func serverOffline(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	_, err := Session.ChannelMessageSend(serverConf.Console.ChannelID, serverConf.OfflineMessage)
	if err != nil {
		log.Printf("Failed to send server offline message to discord: %v", err)
	}
}

func handlePower(power []string, serverID string) {
	if power == nil {
		return
	} else if power[0] == "starting" {
		serverStarting(serverID)
	} else if power[0] == "stopping" {
		serverStopping(serverID)
	} else if power[0] == "running" {
		serverOnline(serverID)
	} else if power[0] == "offline" {
		serverOffline(serverID)
	} else {
		log.Printf("Unknown power state: %v", power)
		return
	}

}
