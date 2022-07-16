package main

import (
	"flag"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/commands"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

var (
	err     error
	config  = conf.GetConf()
	bot     *discordgo.Session
	GuildID = flag.String("guild", config.Discord.ServerID, "Guild ID")
)

func init() {
	var BotToken string
	flag.StringVar(&BotToken, "token", config.Discord.Token, "Bot Token")
	flag.Parse()

	bot, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Invalid Bot Configuration:", err)
	}

}

func main() {
	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	err := bot.Open()
	if err != nil {
		log.Fatal("Cannot open connection:", err)
	}
	log.Println("Adding Commands")
	commandRegistration := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, rawCommand := range commands.Commands {
		command, err := bot.ApplicationCommandCreate(bot.State.User.ID, *GuildID, rawCommand)
		if err != nil {
			log.Fatalf("Failed to create %v: %v", rawCommand.Name, err)
		}
		commandRegistration[i] = command
	}
	defer func(bot *discordgo.Session) {
		err := bot.Close()
		if err != nil {

		}
	}(bot)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
