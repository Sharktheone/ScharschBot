package commands

import (
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/embed"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var (
	config     = conf.GetConf()
	footerIcon = config.Discord.FooterIcon
)
var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"whitelist": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options[0].Options {
			optionMap[opt.Name] = opt
		}
		var username = i.Member.User.String()
		switch options[0].Name {
		case "add":
			name := strings.ToLower(optionMap["name"].StringValue())
			var (
				messageEmbed discordgo.MessageEmbed
				removeSelect discordgo.SelectMenu
				noFree       = false
			)

			if mongodb.Ready {
				alreadyListed, existingAcc, freeAccount, allowed, mcBan, dcBan := whitelist.Add(name, i.Member.User.ID, i.Member.Roles)
				listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
				mcBans := whitelist.CheckBans(i.Member.User.ID)
				var (
					removeOptions []discordgo.SelectMenuOption
				)
				for _, acc := range listedAccounts {
					removeOptions = append(removeOptions, discordgo.SelectMenuOption{
						Label: acc,
						Value: acc,
					})
				}
				removeSelect = discordgo.SelectMenu{
					Placeholder: "Remove accounts",
					CustomID:    "remove_select",
					Options:     removeOptions,
				}

				if !mcBan && !dcBan {
					if allowed {
						if freeAccount {
							if existingAcc {
								if alreadyListed {
									messageEmbed = embed.WhitelistAlreadyListed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
								} else {
									messageEmbed = embed.WhitelistAdding(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
								}
							} else {
								messageEmbed = embed.WhitelistNotExisting(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
							}
						} else {
							messageEmbed = embed.WhitelistNoFreeAccounts(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
							noFree = true
						}
					} else {
						messageEmbed = embed.WhitelistAddNotAllowed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)

					}
				} else {
					messageEmbed = embed.WhitelistBanned(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, dcBan)
				}
			} else {
				messageEmbed = embed.DatabaseNotReady
			}
			var err error
			if noFree {
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							&messageEmbed,
						},
						Components: []discordgo.MessageComponent{
							discordgo.ActionsRow{
								Components: []discordgo.MessageComponent{
									&removeSelect,
								},
							},
						},
					},
				})
			} else {
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
				log.Printf("Failed execute command whitelistadd: %v", err)
			}
		case "remove":
			name := strings.ToLower(optionMap["name"].StringValue())
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				allowed, onWhitelist := whitelist.Remove(name, i.Member.User.ID, i.Member.Roles)
				listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
				mcBans := whitelist.CheckBans(i.Member.User.ID)

				if allowed {
					if onWhitelist {
						messageEmbed = embed.WhitelistRemoving(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
					} else {
						messageEmbed = embed.WhitelistNotListed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
					}
				} else {
					messageEmbed = embed.WhitelistRemoveNotAllowed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
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
		case "myaccounts":
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				accounts, allowed, found, bannedPlayers := whitelist.HasListed(i.Member.User.ID, i.Member.User.ID, i.Member.Roles)
				if allowed {
					if found || len(bannedPlayers) > 0 {
						messageEmbed = embed.WhitelistHasListed(accounts, i.Member.User.ID, i.Member.User.AvatarURL("40"), i.Member.User.Username, bannedPlayers, footerIcon, username)
					} else {
						messageEmbed = embed.WhitelistNoAccounts(i.Member.User.ID, i.Member.User.AvatarURL("40"), i.Member.User.Username)
					}
				} else {
					messageEmbed = embed.WhitelistUserNotAllowed(i.Member.User.ID, i.Member.User.AvatarURL("40"), i.Member.User.Username, accounts, bannedPlayers, footerIcon, username)
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
		case "removemyaccounts":
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				hasListedAccounts, listedAccounts := whitelist.RemoveMyAccounts(i.Member.User.ID)
				mcBans := whitelist.CheckBans(i.Member.User.ID)

				if hasListedAccounts {
					messageEmbed = embed.WhitelistRemoveMyAccounts(listedAccounts, i.Member.User.ID, i.Member.AvatarURL("40"), i.Member.User.Username, mcBans)
				} else {
					messageEmbed = embed.WhitelistNoAccounts(i.Member.User.ID, i.Member.AvatarURL("40"), i.Member.User.Username)
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
		}

	},
	"whitelistadmin": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options[0].Options {
			optionMap[opt.Name] = opt
		}
		var username = i.Member.User.String()
		switch options[0].Name {
		case "whois":
			name := strings.ToLower(optionMap["name"].StringValue())
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {

				userID, allowed, found := whitelist.Whois(name, i.Member.User.ID, i.Member.Roles)
				listedAccounts := whitelist.ListedAccountsOf(userID)
				mcBans := whitelist.CheckBans(userID)
				if allowed {
					if found {
						messageEmbed = embed.WhitelistIsListedBy(name, userID, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
					} else {
						messageEmbed = embed.WhitelistNotListed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
					}
				} else {
					messageEmbed = embed.WhitelistWhoisNotAllowed(name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID), footerIcon, i.Member.User.AvatarURL("40"), username)
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
		case "user":
			user := optionMap["user"].UserValue(s)
			userID := user.ID
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				accounts, allowed, found, bannedPlayers := whitelist.HasListed(userID, i.Member.User.ID, i.Member.Roles)
				if allowed {
					if found || len(bannedPlayers) > 0 {
						messageEmbed = embed.WhitelistHasListed(accounts, userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, bannedPlayers, footerIcon, username)
					} else {
						messageEmbed = embed.WhitelistNoAccounts(userID, i.Member.User.AvatarURL("40"), i.Member.User.Username)
					}
				} else {
					messageEmbed = embed.WhitelistUserNotAllowed(userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, accounts, bannedPlayers, footerIcon, username)
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
		case "banuser":
			user := optionMap["user"].UserValue(s)
			userID := user.ID
			banAccounts := true
			banAccounts = optionMap["removeaccounts"].BoolValue()
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
				mcBans := whitelist.CheckBans(i.Member.User.ID)
				allowed, listedAccounts := whitelist.BanUserID(i.Member.User.ID, i.Member.Roles, userID, banAccounts)
				if allowed {
					messageEmbed = embed.WhitelistBanUserID(listedAccounts, userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, mcBans, footerIcon, username)
				} else {
					messageEmbed = embed.WhitelistBanUserIDNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, userID, listedAccounts, mcBans, footerIcon, username)
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
		case "banaccount":
			name := strings.ToLower(optionMap["name"].StringValue())
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {

				allowed, listedAccounts, userID := whitelist.BanAccount(i.Member.User.ID, i.Member.Roles, name)
				mcBans := whitelist.CheckBans(userID)
				if allowed {
					messageEmbed = embed.WhitelistBanAccount(name, listedAccounts, userID, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
				} else {
					messageEmbed = embed.WhitelistBanAccountNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID), footerIcon, username)
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
		case "unbanuser":
			user := optionMap["user"].UserValue(s)
			userID := user.ID
			unbanAccounts := false
			unbanAccounts = optionMap["unbanaccounts"].BoolValue()
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
				mcBans := whitelist.CheckBans(userID)
				allowed := whitelist.UnBanUserID(i.Member.User.ID, i.Member.Roles, userID, unbanAccounts)
				if allowed {
					messageEmbed = embed.WhitelistUnBanUserID(userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, mcBans, listedAccounts, footerIcon, username)
				} else {
					messageEmbed = embed.WhitelistBanUserIDNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, userID, listedAccounts, mcBans, footerIcon, username)
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
		case "unbanaccount":
			name := strings.ToLower(optionMap["name"].StringValue())
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				allowed, listedAccounts := whitelist.UnBanAccount(i.Member.User.ID, i.Member.Roles, name)
				if allowed {
					messageEmbed = embed.WhitelistUnBanAccount(name, listedAccounts, footerIcon, i.Member.User.AvatarURL("40"), username)
				} else {
					messageEmbed = embed.WhitelistBanAccountNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, whitelist.ListedAccountsOf(i.Member.User.ID), whitelist.CheckBans(i.Member.User.ID), footerIcon, username)
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
		case "removeall":
			var (
				messageEmbed discordgo.MessageEmbed
				err          error
			)
			if mongodb.Ready {
				listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
				mcBans := whitelist.CheckBans(i.Member.User.ID)
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
					messageEmbed = embed.WhitelistRemoveAllNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, listedAccounts, mcBans, footerIcon, username)
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

		}
	},
	"remove_yes": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var username = i.Member.User.String()
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
			mcBans := whitelist.CheckBans(i.Member.User.ID)
			allowed, onWhitelist := whitelist.RemoveAll(i.Member.User.ID, i.Member.Roles)
			if allowed {
				if onWhitelist {
					messageEmbed = embed.WhitelistRemoveAll(i.Member.User.AvatarURL("40"), i.Member.User.Username)
				} else {
					messageEmbed = embed.WhitelistRemoveAllNoWhitelistEntries(i.Member.User.AvatarURL("40"), i.Member.User.Username)
				}
			} else {
				messageEmbed = embed.WhitelistRemoveAllNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.Username, listedAccounts, mcBans, footerIcon, username)
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
	"remove_select": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var username = i.Member.User.String()
		data := i.MessageComponentData()
		name := data.Values[0]
		var messageEmbed discordgo.MessageEmbed
		if mongodb.Ready {
			allowed, onWhitelist := whitelist.Remove(name, i.Member.User.ID, i.Member.Roles)
			listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
			mcBans := whitelist.CheckBans(i.Member.User.ID)

			if allowed {
				if onWhitelist {
					messageEmbed = embed.WhitelistRemoving(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
				} else {
					messageEmbed = embed.WhitelistNotListed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
				}
			} else {
				messageEmbed = embed.WhitelistRemoveNotAllowed(name, listedAccounts, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username)
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
}
