package srv

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"github.com/bwmarrin/discordgo"
	"log"
	"math"
	"strings"
	"time"
)

func channelStats() {
	for _, server := range config.Pterodactyl.Servers {
		var stat pterodactyl.ServerStat
		for _, serverStat := range pterodactyl.ServerStats {
			if serverStat.Name == server.ServerName {
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
			info = strings.ReplaceAll(info, "{{state}}", stat.Status)
			info = strings.ReplaceAll(info, "{{networkIn}}", convertSize(stat.Network.Rx))
			info = strings.ReplaceAll(info, "{{networkOut}}", convertSize(stat.Network.Tx))
			info = strings.ReplaceAll(info, "{{disk}}", convertSize(stat.Disk))
			info = strings.ReplaceAll(info, "{{uptime}}", convertTime(stat.Uptime))
			_, err := Session.ChannelEditComplex(server.ChannelInfo.ChannelID, &discordgo.ChannelEdit{
				Topic: info,
			})
			if err != nil {
				log.Printf("Failed to edit channel topic: %v", err)
			}
		}
	}
}
func serverStarting(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	if serverConf.StateMessages.StartEnabled {
		_, err := Session.ChannelMessageSend(serverConf.StateMessages.ChannelID, serverConf.StateMessages.Start)
		if err != nil {
			log.Printf("Failed to send server start message to discord: %v", err)
		}
	}
}
func serverStopping(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	if serverConf.StateMessages.StopEnabled {
		_, err := Session.ChannelMessageSend(serverConf.StateMessages.ChannelID, serverConf.StateMessages.Stop)
		if err != nil {
			log.Printf("Failed to send server stop message to discord: %v", err)
		}
	}
}
func serverOnline(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	if serverConf.StateMessages.OnlineEnabled {
		_, err := Session.ChannelMessageSend(serverConf.StateMessages.ChannelID, serverConf.StateMessages.Online)
		if err != nil {
			log.Printf("Failed to send server online message to discord: %v", err)
		}
	}
}
func serverOffline(serverID string) {
	serverConf := conf.GetServerConf(serverID, "")
	if serverConf.StateMessages.OfflineEnabled {
		_, err := Session.ChannelMessageSend(serverConf.StateMessages.ChannelID, serverConf.StateMessages.Offline)
		if err != nil {
			log.Printf("Failed to send server offline message to discord: %v", err)
		}
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
