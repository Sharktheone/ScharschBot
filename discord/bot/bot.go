package bot

import (
	"flag"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/console"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/commands"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config              = conf.GetConf()
	GuildID             = flag.String("guild", config.Discord.ServerID, "Guild ID")
	Session             *discordgo.Session
	commandRegistration = make([]*discordgo.ApplicationCommand, len(commands.Commands))
)

func init() {
	var BotToken string
	flag.StringVar(&BotToken, "token", config.Discord.Token, "Bot Token")
	flag.Parse()
	var err error
	Session, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Invalid Bot Configuration:", err)
	}

	err = Session.Open()
	if err != nil {
		log.Fatal("Cannot open connection:", err)
	}
}

func Registration() {
	log.Println("Registering commands...")
	Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	for i, rawCommand := range commands.Commands {
		command, err := Session.ApplicationCommandCreate(Session.State.User.ID, *GuildID, rawCommand)
		if err != nil {
			log.Fatalf("Failed to create %v: %v", rawCommand.Name, err)
		}
		commandRegistration[i] = command
	}
	Session.AddHandler(console.Handler)
	Session.AddHandler(console.ChatHandler)
	log.Println("Commands registered")

}

func RemoveCommands() {
	for _, command := range commandRegistration {
		err := Session.ApplicationCommandDelete(Session.State.User.ID, *GuildID, command.ID)
		if err != nil {
			log.Printf("Failed to delete %v: %v", command.Name, err)
		}
	}
}
