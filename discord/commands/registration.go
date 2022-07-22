package commands

import (
	"flag"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config  = conf.GetConf()
	bot     *discordgo.Session
	GuildID = flag.String("guild", config.Discord.ServerID, "Guild ID")
)

func init() {
	var BotToken string
	flag.StringVar(&BotToken, "token", config.Discord.Token, "Bot Token")
	flag.Parse()
	var err error
	bot, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Invalid Bot Configuration:", err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatal("Cannot open connection:", err)
	}

}

func Registration() *discordgo.Session {
	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	log.Println("Adding Commands")
	commandRegistration := make([]*discordgo.ApplicationCommand, len(Commands))
	for i, rawCommand := range Commands {
		command, err := bot.ApplicationCommandCreate(bot.State.User.ID, *GuildID, rawCommand)
		if err != nil {
			log.Fatalf("Failed to create %v: %v", rawCommand.Name, err)
		}
		commandRegistration[i] = command
	}
	return bot

}
