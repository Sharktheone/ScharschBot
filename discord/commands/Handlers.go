package commands

import (
	"fmt"
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
		alreadyListed, existingAcc := whitelist.Add(name, i.Member.User.ID)
		var message string
		if existingAcc {
			if alreadyListed {
				message = fmt.Sprintf("%v is already on Whitelist", name)
			} else {
				message = fmt.Sprintf("Adding %v to Whitelist", name)
			}
		} else {
			message = fmt.Sprintf("Account %v is not Existing", name)
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: message,
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistadd: %v", err)
		}

	},
	"whitelistremove": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		name := optionMap["name"].StringValue()
		whitelist.Remove(name, i.Member.User.ID, i.Member.Roles)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{

				Content: "Removing " + name + " from Whitelist",
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistremove: %v", err)
		}

	},
	"whitelistwhois": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		name := optionMap["name"].StringValue()
		userID, allowed, found := whitelist.Whois(name, i.Member.User.ID, i.Member.Roles)
		var message string
		if allowed {
			if found {
				message = fmt.Sprintf("Player %v was whitelisted by <@%v>", name, userID)
			} else {
				message = fmt.Sprintf("Player %v was not found on Whitelist", name)
			}
		} else {
			message = "Operation not permitted!"
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{

				Content: message,
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistwhois: %v", err)
		}

	},
}
