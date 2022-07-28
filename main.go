package main

import (
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/bot"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/checkroles"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
)

func main() {

	dcBot := bot.Registration()
	defer func(bot *discordgo.Session) {
		err := bot.Close()
		if err != nil {

		}
	}(dcBot)

	mongodb.Connect()
	defer mongodb.Disconnect()
	defer mongodb.Cancel()

	checkroles.CheckRoles()
	c := cron.New()
	err := c.AddFunc("0 */10 * * * *", checkroles.CheckRoles)
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}
	c.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
