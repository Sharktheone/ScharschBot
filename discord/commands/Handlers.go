package commands

import (
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/embed"
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
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			alreadyListed, existingAcc, freeAccount, allowed, mcBan, dcBan := whitelist.Add(name, i.Member.User.ID, i.Member.Roles)

			if !mcBan && !dcBan {
				if allowed {
					if freeAccount {
						if existingAcc {
							if alreadyListed {
								messageEmbed = embed.WhitelistAlreadyListed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
							} else {
								messageEmbed = embed.WhitelistAdding(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
							}
						} else {
							messageEmbed = embed.WhitelistNotExisting(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
						}
					} else {
						messageEmbed = embed.WhitelistNoFreeAccounts(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
					}
				} else {
					messageEmbed = embed.WhitelistAddNotAllowed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))

				}
			} else {
				messageEmbed = embed.WhitelistBanned(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, dcBan)
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
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
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed, onWhitelist := whitelist.Remove(name, i.Member.User.ID, i.Member.Roles)

			if allowed {
				if onWhitelist {
					messageEmbed = embed.WhitelistRemoving(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
				} else {
					messageEmbed = embed.WhitelistNotListed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
				}
			} else {
				messageEmbed = embed.WhitelistRemoveNotAllowed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
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
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			userID, allowed, found := whitelist.Whois(name, i.Member.User.ID, i.Member.Roles)
			if allowed {
				if found {
					messageEmbed = embed.WhitelistIsListedBy(name, userID, whitelist.ListedAccountsOf(i.Member.User.ID))
				} else {
					messageEmbed = embed.WhitelistNotListed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
				}
			} else {
				messageEmbed = embed.WhitelistWhoisNotAllowed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
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
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			accounts, allowed, found, bannedPlayers := whitelist.HasListed(userID, i.Member.User.ID, i.Member.Roles)
			if allowed {
				if found {
					messageEmbed = embed.WhitelistHasListed(accounts, userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, bannedPlayers)
				} else {
					messageEmbed = embed.WhitelistNoAccounts(userID, i.Member.User.AvatarURL("40"), i.Member.User.Username)
				}
			} else {
				messageEmbed = embed.WhitelistUserNotAllowed(userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, accounts, bannedPlayers)
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistuser: %v", err)
		}
	},
	"whitelistmyaccounts": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			accounts, allowed, found, bannedPlayers := whitelist.HasListed(i.Member.User.ID, i.Member.User.ID, i.Member.Roles)
			if allowed {
				if found {
					messageEmbed = embed.WhitelistHasListed(accounts, i.Member.User.ID, i.Member.User.AvatarURL("40"), i.Member.User.Username, bannedPlayers)
				} else {
					messageEmbed = embed.WhitelistNoAccounts(i.Member.User.ID, i.Member.User.AvatarURL("40"), i.Member.User.Username)
				}
			} else {
				messageEmbed = embed.WhitelistUserNotAllowed(i.Member.User.ID, i.Member.User.AvatarURL("40"), i.Member.User.Username, accounts, bannedPlayers)
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistmyaccounts: %v", err)
		}
	},
	"remove_yes": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed, onWhitelist := whitelist.RemoveAll(i.Member.User.ID, i.Member.Roles)
			if allowed {
				if onWhitelist {
					messageEmbed = embed.WhitelistRemoveAll(i.Member.User.AvatarURL("40"), i.Member.User.Username)
				} else {
					messageEmbed = embed.WhitelistRemoveAllNoWhitelistEntries(i.Member.User.AvatarURL("40"), i.Member.User.Username)
				}
			} else {
				messageEmbed = embed.WhitelistRemoveAllNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute component remove_yes: %v", err)
		}

	},

	"whitelistremoveall": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var (
			messageEmbed discordgo.MessageEmbed
			err          error
		)
		if mongodb.Ready {
			allowed := whitelist.RemoveAllAllowed(i.Member.Roles)
			if allowed {
				var button discordgo.Button
				messageEmbed, button = embed.WhitelistRemoveAllSure(i.Member.User.AvatarURL("40"), i.Member.User.Username)
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							&messageEmbed,
						},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									&button,
								},
							},
						},
					},
				})

			} else {
				messageEmbed = embed.WhitelistRemoveAllNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							&messageEmbed,
						},
					},
				})
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						&messageEmbed,
					},
				},
			})
		}

		if err != nil {
			log.Printf("Failed execute command whitelistremoveall: %v", err)
		}

	},
	"whitelistbanuserid": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		userID := optionMap["userid"].StringValue()
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed, listedAccounts := whitelist.BanUserID(i.Member.User.ID, i.Member.Roles, userID)
			if allowed {
				messageEmbed = embed.WhitelistBanUserID(listedAccounts, userID, i.Member.User.AvatarURL("40"), i.Member.User.Username)
			} else {
				messageEmbed = embed.WhitelistBanUserIDNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, userID, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistbanuserid: %v", err)
		}

	},
	"whitelistbanaccount": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		name := strings.ToLower(optionMap["name"].StringValue())
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed, listedAccounts, userID := whitelist.BanAccount(i.Member.User.ID, i.Member.Roles, name)
			if allowed {
				messageEmbed = embed.WhitelistBanAccount(name, listedAccounts, userID)
			} else {
				messageEmbed = embed.WhitelistBanAccountNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistbanaccount: %v", err)
		}

	},
	"whitelistunbanuserid": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		userID := optionMap["userid"].StringValue()
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed := whitelist.UnBanUserID(i.Member.User.ID, i.Member.Roles, userID)
			if allowed {
				messageEmbed = embed.WhitelistUnBanUserID(userID, i.Member.User.AvatarURL("40"), i.Member.User.Username)
			} else {
				messageEmbed = embed.WhitelistBanUserIDNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, userID, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistunbanuserid: %v", err)
		}

	},
	"whitelistunbanaccount": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		name := strings.ToLower(optionMap["name"].StringValue())
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed, listedAccounts := whitelist.UnBanAccount(i.Member.User.ID, i.Member.Roles, name)
			if allowed {
				messageEmbed = embed.WhitelistUnBanAccount(name, listedAccounts)
			} else {
				messageEmbed = embed.WhitelistBanAccountNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID))
			}
		} else {
			messageEmbed = embed.DatabaseNotReady
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
			},
		})
		if err != nil {
			log.Printf("Failed execute command whitelistunbanaccount: %v", err)
		}

	},
}
