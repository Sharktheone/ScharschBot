package main

import (
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/bot"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/embed"
	"github.com/Sharktheone/Scharsch-bot-discord/srv"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/checkroles"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
)

var config = conf.GetConf()

func main() {
	mongodb.Connect()

	dcBot := bot.Session
	if config.Whitelist.Enabled {
		bot.Registration()
		checkroles.CheckRoles()
		rolesCron := cron.New()
		err := rolesCron.AddFunc("0 */10 * * * *", checkroles.CheckRoles)
		if err != nil {
			log.Fatalf("Error adding cron job: %v", err)
		}
		rolesCron.Start()

		consoleCron := cron.New()
		err = consoleCron.AddFunc("0 * * * * *", checkroles.CheckRoles)
		if err != nil {
			log.Fatalf("Error adding cron job: %v", err)
		}
		consoleCron.Start()
	}
	embed.BotAvatarURL = dcBot.State.User.AvatarURL("40")
	srv.Start()

	defer mongodb.Disconnect()
	defer mongodb.Cancel()
	defer func(bot *discordgo.Session) {
		err := bot.Close()
		if err != nil {

		}

	}(dcBot)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	defer bot.RemoveCommands()

}
