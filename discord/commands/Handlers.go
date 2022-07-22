package commands

import (
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist"
	"github.com/bwmarrin/discordgo"
	"log"
)

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
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
			log.Printf("Failed execute command whitelistadd: %v", err)
		}
		whitelist.Add(name, i.Member.User.ID)

	},
}
