package commands

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"whitelistadd": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		name := strings.ToLower(optionMap["name"].StringValue())
		var message string
		if mongodb.Ready {
			alreadyListed, existingAcc, freeAccount, allowed := whitelist.Add(name, i.Member.User.ID, i.Member.Roles)

			if allowed {
				if freeAccount {
					if existingAcc {
						if alreadyListed {
							message = fmt.Sprintf("%v is already on whitelist", name)
						} else {
							message = fmt.Sprintf("Adding %v to whitelist", name)
						}
					} else {
						message = fmt.Sprintf("Account %v is not existing", name)
					}
				} else {
					message = "You have no free Account anymore"
				}

			} else {
				message = "You are not allowed to add accounts"
			}
		} else {
			message = "Database is not ready, please try again later"
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

		name := strings.ToLower(optionMap["name"].StringValue())
		var message string
		if mongodb.Ready {
			allowed, onWhitelist := whitelist.Remove(name, i.Member.User.ID, i.Member.Roles)

			if allowed {

				if onWhitelist {
					message = fmt.Sprintf("Removing %v from whitelist", name)
				} else {
					message = fmt.Sprintf("%v is not on the whitelist", name)
				}

			} else {
				message = "Operation not permitted!"
			}
		} else {
			message = "Database is not ready, please try again later"
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{

				Content: message,
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
		name := strings.ToLower(optionMap["name"].StringValue())
		var message string
		if mongodb.Ready {
			userID, allowed, found := whitelist.Whois(name, i.Member.User.ID, i.Member.Roles)
			if allowed {
				if found {
					message = fmt.Sprintf("Player %v was whitelisted by <@%v>", name, userID)
				} else {
					message = fmt.Sprintf("Player %v was not found on Whitelist", name)
				}
			} else {
				message = "Operation not permitted!"
			}
		} else {
			message = "Database is not ready, please try again later"
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
	"whitelistuser": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		userID := optionMap["userid"].StringValue()
		var message string
		if mongodb.Ready {
			accounts, allowed, found := whitelist.HasListed(userID, i.Member.User.ID, i.Member.Roles)
			if allowed {
				if found {
					message = fmt.Sprintf("<@%v> has whitelisted players %v", userID, accounts)
				} else {
					message = fmt.Sprintf("UserID %v was not found on Whitelist", userID)
				}
			} else {
				message = "Operation not permitted!"
			}
		} else {
			message = "Database is not ready, please try again later"
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
	"whitelistmyaccounts": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var message string
		if mongodb.Ready {
			accounts, allowed, found := whitelist.HasListed(i.Member.User.ID, i.Member.User.ID, i.Member.Roles)
			if allowed {
				if found {
					message = fmt.Sprintf("You have whitelisted players %v", accounts)
				} else {
					message = "You have no whitelisted players"
				}
			} else {
				message = "Operation not permitted!"
			}
		} else {
			message = "Database is not ready, please try again later"
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
	"whitelistremoveall": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var message string
		allowed, onWhitelist := whitelist.RemoveAll(i.Member.User.ID, i.Member.Roles)
		if mongodb.Ready {
			if allowed {
				if onWhitelist {
					message = fmt.Sprintf("Removing everyone from whitelist")
				} else {
					message = fmt.Sprintf("No one is not on the whitelist")
				}
			} else {
				message = "Operation not permitted!"
			}
		} else {
			message = "Database is not ready, please try again later"
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: message,
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistremove: %v", err)
		}

	},
}
