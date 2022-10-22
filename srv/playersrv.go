package srv

import (
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/bot"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/embed"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/whitelist"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	config    = conf.GetConf()
	eventJson struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Type   string `json:"type"`
		Server string `json:"server"`
	}
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
			go pterodactyl.Websocket(server.ServerID, pterodactyl.ConsoleOutput, ConsoleSrv, server.Console.MessageLines, time.Duration(maxTime), false, doStats)
			doStats = false
		}
		if server.StateMessages {
			go pterodactyl.Websocket(server.ServerID, pterodactyl.Status, handlePower, 0, 0, true, doStats)
			doStats = false
		}
		if server.ChannelInfo.Enabled {
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
		userID, _ := whitelist.GetOwner(eventJson.Name)
		var (
			FooterIcon string
			username   string
		)
		member, successful := bot.GetUserProfile(userID)

		if successful {
			w.WriteHeader(http.StatusNoContent)
			FooterIcon = member.User.AvatarURL("40")
			username = member.User.String()
		} else {
			w.WriteHeader(http.StatusNotFound)
			FooterIcon = config.Discord.EmbedErrorIcon
		}
		switch eventJson.Type {
		case "chat":
			if serverConf.Chat.Embed {
				accounts, bannedAccounts := checkAccount(strings.ToLower(eventJson.Name))
				messageEmbed := embed.Chat(eventJson.Name, accounts, bannedAccounts, fmt.Sprintf("%v %v", serverConf.Chat.Prefix, eventJson.Value), serverConf.Chat.EmbedFooter, serverConf.Chat.EmbedOneLine, serverConf.Chat.FooterIcon, FooterIcon, username)
				_, err = Session.ChannelMessageSendEmbed(serverConf.Chat.ChannelID, &messageEmbed)
			} else {
				_, err = Session.ChannelMessageSend(serverConf.Chat.ChannelID, fmt.Sprintf("%v%v %v", eventJson.Name, serverConf.Chat.Prefix, eventJson.Value))
			}

			if err != nil {
				log.Printf("Failed to send Chat (embed): %v", err)
			}
		case "death":
			accounts, bannedAccounts := checkAccount(strings.ToLower(eventJson.Name))
			messageEmbed := embed.PlayerDeath(eventJson.Name, accounts, bannedAccounts, eventJson.Value, serverConf.SRV.Footer, serverConf.SRV.OneLine, serverConf.SRV.FooterIcon, FooterIcon, username)
			_, err = Session.ChannelMessageSendEmbed(serverConf.SRV.ChannelID, &messageEmbed)
			if err != nil {
				log.Printf("Failed to send Death embed: %v", err)
			}
		case "advancement":
			accounts, bannedAccounts := checkAccount(strings.ToLower(eventJson.Name))
			messageEmbed := embed.PlayerAdvancement(eventJson.Name, accounts, bannedAccounts, eventJson.Value, serverConf.SRV.Footer, serverConf.SRV.OneLine, serverConf.SRV.FooterIcon, FooterIcon, username)
			_, err = Session.ChannelMessageSendEmbed(serverConf.SRV.ChannelID, &messageEmbed)
			if err != nil {
				log.Printf("Failed to send Advancement embed: %v", err)
			}
		case "join":
			OnlinePlayers = append(OnlinePlayers, strings.ToLower(eventJson.Name))
			accounts, bannedAccounts := checkAccount(strings.ToLower(eventJson.Name))
			messageEmbed := embed.PlayerJoin(strings.ToLower(eventJson.Name), accounts, bannedAccounts, serverConf.SRV.Footer, serverConf.SRV.OneLine, serverConf.SRV.FooterIcon, FooterIcon, username)
			_, err = Session.ChannelMessageSendEmbed(serverConf.SRV.ChannelID, &messageEmbed)
			if err != nil {
				log.Printf("Failed to send Join embed: %v", err)
			}
		case "quit":
			for i, player := range OnlinePlayers {
				if player == strings.ToLower(eventJson.Name) {
					OnlinePlayers = append(OnlinePlayers[:i], OnlinePlayers[i+1:]...)
					break
				}
			}
			accounts, bannedAccounts := checkAccount(strings.ToLower(eventJson.Name))
			messageEmbed := embed.PlayerQuit(eventJson.Name, accounts, bannedAccounts, serverConf.SRV.Footer, serverConf.SRV.OneLine, serverConf.SRV.FooterIcon, FooterIcon, username)
			_, err = Session.ChannelMessageSendEmbed(serverConf.SRV.ChannelID, &messageEmbed)
			if err != nil {
				log.Printf("Failed to send Quit embed: %v", err)
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
	ListedAccounts := whitelist.ListedAccountsOf(userID)
	Banned := whitelist.CheckBans(userID)
	return ListedAccounts, Banned
}
