package main

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

var (
	bot               *discordgo.Session
	DefaultPermission = true
	GuildID           = flag.String("guild", "796764100532109312", "Guild ID")
	commands          = []*discordgo.ApplicationCommand{
		{
			Name:              "whitelistadd",
			Description:       "Add your account to the Whitelist",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the account to add",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"whitelistadd": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			name := optionMap["name"].StringValue()
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Adding " + name + " to Whitelist",
				},
			})
			if err != nil {
				return
			}
			whitelistadd(name)

		},
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	var BotToken string
	flag.StringVar(&BotToken, "token", os.Getenv("BOT_TOKEN"), "Bot Token")
	flag.Parse()

	bot, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Invalid Bot Configuration:", err)
	}
}

func main() {

	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	err := bot.Open()
	if err != nil {
		log.Fatal("Cannot open connection:", err)
	}
	log.Println("Adding Commands")
	commandRegistrator := make([]*discordgo.ApplicationCommand, len(commands))
	for i, rawCommand := range commands {
		command, err := bot.ApplicationCommandCreate(bot.State.User.ID, *GuildID, rawCommand)
		if err != nil {
			log.Fatalf("Failed to create %v: %v", rawCommand.Name, err)
		}
		commandRegistrator[i] = command
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

func whitelistadd(username string) {
	log.Println("*Add " + username + " to whitelist")
}
