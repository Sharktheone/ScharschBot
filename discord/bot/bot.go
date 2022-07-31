package bot

import (
	"flag"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/commands"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config  = conf.GetConf()
	bot     *discordgo.Session
	GuildID = flag.String("guild", config.Discord.ServerID, "Guild ID")
	Session *discordgo.Session
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
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}

		case discordgo.InteractionMessageComponent:
			if h, ok := commands.Handlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	log.Println("Adding Commands")
	commandRegistration := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, rawCommand := range commands.Commands {
		command, err := bot.ApplicationCommandCreate(bot.State.User.ID, *GuildID, rawCommand)
		if err != nil {
			log.Fatalf("Failed to create %v: %v", rawCommand.Name, err)
		}
		commandRegistration[i] = command
	}
	Session = bot
	return bot

}
