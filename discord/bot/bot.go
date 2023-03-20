package bot

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/console"
	"Scharsch-Bot/discord/commands"
	"flag"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config              = conf.Config
	GuildID             = flag.String("guild", config.Discord.ServerID, "Guild ID")
	Session             *discordgo.Session
	commandRegistration = make([]*discordgo.ApplicationCommand, len(commands.Commands))
)

func init() {
	var BotToken = flag.String("token", config.Discord.Token, "Discord Bot Token")
	Session, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatal("Invalid Bot Configuration:", err)
	}

	if err := Session.Open(); err != nil {
		log.Fatal("Cannot open connection to discord:", err)
	}
}

func Registration() {
	log.Println("Registering commands...")
	Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			} else {
				log.Printf("No handler for %v", i.ApplicationCommandData().Name)
			}

		case discordgo.InteractionMessageComponent:
			if h, ok := commands.Handlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			} else {
				log.Printf("No handler for %v", i.MessageComponentData().CustomID)
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
