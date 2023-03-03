package srv

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/bot"
	"Scharsch-Bot/discord/discordMember"
	"Scharsch-Bot/discord/embed"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	config        = conf.GetConf()
	eventJson     types.EventJson
	Session       = bot.Session
	OnlinePlayers []string
	port          = config.SRV.API.Port
	APIUser       = config.SRV.API.User
	APIPassword   = config.SRV.API.Password
)

func Start() {
	http.HandleFunc("/", EventHandler)
	log.Printf("Starting http server on port %v", port)
	addr := fmt.Sprintf(":%v", port)
	var err error
	go func() { err = http.ListenAndServe(addr, nil) }()
	if err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}
	log.Println("Started http server")

	for _, server := range config.Pterodactyl.Servers {
		var doStats = true
		if server.Console.Enabled {
			maxTime := server.Console.MaxTimeInSeconds * int(time.Second)
			// TODO remove go-routine - fixed maybe
			//goland:noinspection GoBoolExpressions
			go pterodactyl.Websocket(server.ServerID, pterodactyl.ConsoleOutput, ConsoleSrv, server.Console.MessageLines, time.Duration(maxTime), false, doStats)
			doStats = false
		}
		if server.StateMessages.Enabled {
			// TODO remove go-routine - fixed maybe
			go pterodactyl.Websocket(server.ServerID, pterodactyl.Status, handlePower, 0, 0, true, doStats)
			doStats = false
		}
		if server.ChannelInfo.Enabled {
			// TODO remove go-routine - fixed maybe
			go pterodactyl.Websocket(server.ServerID, pterodactyl.Stats, nil, 0, 0, true, doStats)
			doStats = false

		}

	}
	channelCron := cron.New()
	err = channelCron.AddFunc("0 * * * * *", channelStats)
	if err != nil {
		log.Fatalf("Error adding ChannelCron job: %v", err)
	}
	channelCron.Start()

}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, _ := r.BasicAuth()
	if user == "" || pass == "" {
		w.WriteHeader(http.StatusUnauthorized)
	} else if user != APIUser || pass != APIPassword {
		w.WriteHeader(http.StatusForbidden)
	} else if user == APIUser && pass == APIPassword {
		var (
			serverConf conf.Server
			found      = false
		)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&eventJson)
		if err != nil {
			log.Printf("Failed to decode json: %v", err)
		}

		for _, server := range config.Pterodactyl.Servers {
			if server.ServerName == eventJson.Server {
				serverConf = conf.Server(server)
				found = true
			}
		}
		if !found {
			log.Printf("Server %v not found in Pterodactyl server config", eventJson.Server)
			return
		}
		userID, onWhitelist := whitelist.GetOwner(eventJson.Name)
		var (
			FooterIcon string
			username   string
		)
		if onWhitelist {
			member, successful := discordMember.GetUserProfile(userID, Session)

			if successful {
				w.WriteHeader(http.StatusNoContent)
				FooterIcon = member.User.AvatarURL("40")
				username = member.User.String()
			} else {
				w.WriteHeader(http.StatusOK)
				FooterIcon = config.Discord.EmbedErrorIcon
			}
		} else {
			w.WriteHeader(http.StatusResetContent)
			FooterIcon = config.Discord.EmbedErrorIcon
		}
		checkAccount(strings.ToLower(eventJson.Name))
		switch eventJson.Type {
		case "chat":
			if serverConf.Chat.Embed {
				messageEmbed := embed.Chat(eventJson, serverConf, FooterIcon, username, Session)
				for _, channelID := range serverConf.Chat.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
					}
				}
			} else {
				for _, channelID := range serverConf.Chat.ChannelID {
					_, err = Session.ChannelMessageSend(channelID, fmt.Sprintf("%v%v %v", eventJson.Name, serverConf.Chat.Prefix, eventJson.Value))
					if err != nil {
						log.Printf("Failed to send Chat (embed): %v (channelID: %v)", err, channelID)
					}
				}
			}

		case "death":
			if serverConf.SRV.Events.Death {
				messageEmbed := embed.PlayerDeath(eventJson, serverConf, FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Death embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		case "advancement":
			if serverConf.SRV.Events.Advancement {
				messageEmbed := embed.PlayerAdvancement(eventJson, serverConf, FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Advancement embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		case "join":
			if serverConf.SRV.Events.Join {
				OnlinePlayers = append(OnlinePlayers, strings.ToLower(eventJson.Name))
				messageEmbed := embed.PlayerJoin(serverConf, strings.ToLower(eventJson.Name), FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Join embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		case "quit":
			if serverConf.SRV.Events.Quit {
				for i, player := range OnlinePlayers {
					if player == strings.ToLower(eventJson.Name) {
						OnlinePlayers = append(OnlinePlayers[:i], OnlinePlayers[i+1:]...)
						break
					}
				}
				messageEmbed := embed.PlayerQuit(serverConf, strings.ToLower(eventJson.Name), FooterIcon, username, Session)
				for _, channelID := range serverConf.SRV.ChannelID {
					_, err = Session.ChannelMessageSendEmbed(channelID, &messageEmbed)
					if err != nil {
						log.Printf("Failed to send Quit embed: %v (channelID: %v)", err, channelID)
					}
				}
			}
		}

		eventJson.Type = ""
	}
}

func checkAccount(Name string) (accounts []string, bannedAccounts []string) {
	userID, onWhitelist := whitelist.GetOwner(Name)
	if config.Whitelist.KickUnWhitelisted {
		if !onWhitelist {
			command := fmt.Sprintf(config.Whitelist.KickCommand, eventJson.Name)
			for _, listedServer := range config.Whitelist.Servers {
				for _, server := range config.Pterodactyl.Servers {
					if server.ServerName == listedServer {
						pterodactyl.SendCommand(command, server.ServerID)
					}
				}
			}
		}
	}
	ListedAccounts := whitelist.ListedAccountsOf(userID, false)
	Banned := whitelist.CheckBans(userID)
	return ListedAccounts, Banned
}
