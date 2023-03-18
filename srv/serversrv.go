package srv

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/pterodactyl"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math"
	"strings"
	"time"
)

func channelStats() {
	for _, server := range config.Pterodactyl.Servers {
		var stat pterodactyl.ServerStatus
		for _, serverStat := range pterodactyl.ServerStats {
			if serverStat.State == server.ServerName {
				stat = *serverStat
			}
		}
		if server.ChannelInfo.Enabled {
			info := server.ChannelInfo.Format
			// {{onlineNumber}} players are Online | Server online for {{uptime}} | {{ram}} RAM | {{cpu}}% CPU | Server is {{state}} | Network in {{networkIn}} | Network out {{networkOut}}
			info = strings.ReplaceAll(info, "{{onlineNumber}}", fmt.Sprintf("%v", len(OnlinePlayers)))
			//strings.ReplaceAll(info, "{{uptime}}", stats.Uptime)
			info = strings.ReplaceAll(info, "{{ram}}", convertSize(stat.Ram))
			info = strings.ReplaceAll(info, "{{cpu}}", fmt.Sprintf("%.2f", stat.Cpu))
			info = strings.ReplaceAll(info, "{{state}}", stat.State)
			info = strings.ReplaceAll(info, "{{networkIn}}", convertSize(stat.Network.Rx))
			info = strings.ReplaceAll(info, "{{networkOut}}", convertSize(stat.Network.Tx))
			info = strings.ReplaceAll(info, "{{disk}}", convertSize(stat.Disk))
			info = strings.ReplaceAll(info, "{{uptime}}", convertTime(stat.Uptime))
			for _, channelID := range server.ChannelInfo.ChannelID {
				_, err := Session.ChannelEditComplex(channelID, &discordgo.ChannelEdit{
					Topic: info,
				})
				if err != nil {
					log.Printf("Failed to edit channel topic: %v (channelID %v)", err, channelID)
				}
			}
		}
	}
}
func serverStarting(server *conf.Server) {
	if server.StateMessages.StartEnabled {
		for _, channelID := range server.StateMessages.ChannelID {
			_, err := Session.ChannelMessageSend(channelID, server.StateMessages.Start)
			if err != nil {
				log.Printf("Failed to send server start message to discord: %v, (channelID %v)", err, channelID)
			}
		}
	}
}
func serverStopping(server *conf.Server) {
	if server.StateMessages.StopEnabled {
		for _, channelID := range server.StateMessages.ChannelID {
			_, err := Session.ChannelMessageSend(channelID, server.StateMessages.Stop)
			if err != nil {
				log.Printf("Failed to send server stop message to discord: %v (channelID: %v)", err, channelID)
			}
		}
	}
}
func serverOnline(server *conf.Server) {
	if server.StateMessages.OnlineEnabled {
		for _, channelID := range server.StateMessages.ChannelID {
			_, err := Session.ChannelMessageSend(channelID, server.StateMessages.Online)
			if err != nil {
				log.Printf("Failed to send server online message to discord: %v (channelID: %v)", err, channelID)
			}
		}
	}
}
func serverOffline(server *conf.Server) {
	if server.StateMessages.OfflineEnabled {
		for _, channelID := range server.StateMessages.ChannelID {
			_, err := Session.ChannelMessageSend(channelID, server.StateMessages.Offline)
			if err != nil {
				log.Printf("Failed to send server offline message to discord: %v (channelID: %v)", err, channelID)
			}
		}
	}
}

func handlePower(power []string, serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	if power == nil {
		return
	} else if power[0] == "starting" {
		serverStarting(&serverConf)
	} else if power[0] == "stopping" {
		serverStopping(&serverConf)
	} else if power[0] == "running" {
		serverOnline(&serverConf)
	} else if power[0] == "offline" {
		serverOffline(&serverConf)
	} else {
		log.Printf("Unknown power state: %v", power)
		return
	}
}
func HandlePower(status string, server *conf.Server) {
	if status == "" {
		return
	}
	switch status {
	case pterodactyl.PowerStatusStarting:
		serverStarting(server)
	case pterodactyl.PowerStatusStopping:
		serverStopping(server)
	case pterodactyl.PowerStatusRunning:
		serverOnline(server)
	case pterodactyl.PowerStatusOffline:
		serverOffline(server)
	default:
		log.Printf("Unknown power state: %v", status)
	}
}
func convertSize(bytes int) string {
	floatBytes := float64(bytes)
	if bytes < 1024 {
		return fmt.Sprintf("%.2f B", floatBytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.2f KB", floatBytes/1024)
	} else if bytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", floatBytes/(1024*1024))
	} else if bytes < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f GB", floatBytes/(1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2f TB", floatBytes/(1024*1024*1024*1024))
	}
}

func convertTime(milliseconds int) string {
	duration := time.Duration(milliseconds * int(time.Millisecond))
	if duration.Hours() >= 24 {
		days := math.Floor(duration.Hours() / 24)
		return fmt.Sprintf("%vd %vh %vm", days, math.Floor(duration.Hours()-(days*24)), math.Floor(duration.Minutes()-(math.Floor(duration.Hours())*60)))
	} else {
		return fmt.Sprintf("%v", duration)
	}
}
