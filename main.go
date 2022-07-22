package main

import (
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/commands"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

func main() {

	bot := commands.Registration()
	defer func(bot *discordgo.Session) {
		err := bot.Close()
		if err != nil {

		}
	}(bot)

	mongodb.Connect()
	defer mongodb.Disconnect()
	defer mongodb.Cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
