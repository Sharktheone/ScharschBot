package commands

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/embed"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
				alreadyListed, existingAcc, freeAccount, allowed, mcBan, dcBan, banReason := whitelist.Add(name, i.Member.User.ID, i.Member.Roles)
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
					messageEmbed = embed.WhitelistBanned(i.Member.User.AvatarURL("40"), i.Member.User.Username, name, dcBan, banReason)
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
				log.Printf("Failed execute command whitelist: %v", err)
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
			var reason = "No reason provided"
			if optionMap["reason"] != nil {
				reason = optionMap["reason"].StringValue()
			}
			userID := user.ID
			banAccounts := true
			if optionMap["removeaccounts"] != nil {
				banAccounts = optionMap["removeaccounts"].BoolValue()
			}
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {
				listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID)
				mcBans := whitelist.CheckBans(i.Member.User.ID)
				allowed, listedAccounts := whitelist.BanUserID(i.Member.User.ID, i.Member.Roles, userID, banAccounts, reason)
				if allowed {
					messageEmbed = embed.WhitelistBanUserID(listedAccounts, userID, i.Member.User.AvatarURL("40"), i.Member.User.Username, mcBans, footerIcon, username, reason)
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
			var reason = "No reason provided"
			if optionMap["reason"] != nil {
				reason = optionMap["reason"].StringValue()
			}
			var messageEmbed discordgo.MessageEmbed
			if mongodb.Ready {

				allowed, listedAccounts, userID := whitelist.BanAccount(i.Member.User.ID, i.Member.Roles, name, reason)
				mcBans := whitelist.CheckBans(userID)
				if allowed {
					messageEmbed = embed.WhitelistBanAccount(name, listedAccounts, userID, mcBans, footerIcon, i.Member.User.AvatarURL("40"), username, reason)
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
			if optionMap["unbanaccounts"] != nil {
				unbanAccounts = optionMap["unbanaccounts"].BoolValue()
			}
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
			log.Printf("Failed execute command whitelistadmin: %v", err)
		}

	},
	"power":                powerMain,
	"power_start":          powerStart,
	"power_restart":        powerRestart,
	"power_stop":           powerStop,
	"power_start_select":   powerStartSelect,
	"power_restart_select": powerRestartSelect,
	"power_stop_select":    powerStopSelect,
}

func powerMain(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, i.ApplicationCommandData().Options[0].Name)
}
func powerStart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, "start")
}
func powerRestart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, "restart")
}
func powerStop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, "stop")
}

func power(s *discordgo.Session, i *discordgo.InteractionCreate, option string) {
	var (
		c               = cases.Title(language.English)
		messageEmbed    discordgo.MessageEmbed
		serverOptions   []discordgo.SelectMenuOption
		restartDisabled = true
		stopDisabled    = true
		startDisabled   = true
	)
	for _, server := range pterodactyl.ServerStates {
		if option == "start" && (server.Status == "offline" || server.Status == "stopping") {
			serverOptions = append(serverOptions, discordgo.SelectMenuOption{
				Label: server.Name,
				Value: server.ID,
			})
		} else if option == "stop" && (server.Status == "starting" || server.Status == "running") {
			serverOptions = append(serverOptions, discordgo.SelectMenuOption{
				Label: server.Name,
				Value: server.ID,
			})
		} else if option == "restart" {
			serverOptions = append(serverOptions, discordgo.SelectMenuOption{
				Label: server.Name,
				Value: server.ID,
			})
		}

	}
	serverSelect := discordgo.SelectMenu{
		Placeholder: fmt.Sprintf("Select a server to %s", c.String(option)),
		CustomID:    fmt.Sprintf("power_%s_select", option),
		Options:     serverOptions,
	}
	if option == "start" {
		startDisabled = true
		stopDisabled = false
		restartDisabled = false
	} else if option == "stop" {
		startDisabled = false
		stopDisabled = true
		restartDisabled = false
	} else if option == "restart" {
		startDisabled = false
		stopDisabled = false
		restartDisabled = true
	} else if option == "status" {
		startDisabled = false
		stopDisabled = false
		restartDisabled = false
	}

	buttonRow := []discordgo.MessageComponent{
		&discordgo.Button{
			Label:    "Start",
			Style:    discordgo.SuccessButton,
			CustomID: "power_start",
			Disabled: startDisabled,
		},
		&discordgo.Button{
			Label:    "Restart",
			Style:    discordgo.PrimaryButton,
			CustomID: "power_restart",
			Disabled: restartDisabled,
		},
		&discordgo.Button{
			Label:    "Stop",
			Style:    discordgo.DangerButton,
			CustomID: "power_stop",
			Disabled: stopDisabled,
		},
	}

	messageEmbed = embed.Power(option)
	var err error
	if option != "status" {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&messageEmbed,
				},
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: buttonRow,
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							&serverSelect,
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
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: buttonRow,
					},
				},
			},
		})
	}
	if err != nil {
		log.Printf("Failed execute command power: %v", err)
	}
}

func powerSelect(s *discordgo.Session, i *discordgo.InteractionCreate, action string) {
	var (
		allowed      = false
		options      = i.MessageComponentData()
		requiredRole []string
		serverConf   = conf.GetServerConf(options.Values[0], "")
	)
	switch action {
	case "start":
		requiredRole = serverConf.PowerActionsRoleIDs.Start
	case "stop":
		requiredRole = serverConf.PowerActionsRoleIDs.Stop
	case "restart":
		requiredRole = serverConf.PowerActionsRoleIDs.Restart

	}
	for _, role := range i.Member.Roles {
		for _, required := range requiredRole {
			if required == role {
				allowed = true
				break
			}
		}
	}
	var messageEmbed discordgo.MessageEmbed
	if !allowed {
		messageEmbed = embed.PowerNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.String(), action, serverConf.ServerName)
	} else {
		messageEmbed = embed.PowerAction(action, serverConf.ServerName, i.Member.User.AvatarURL("40"), i.Member.User.Username)
		switch action {
		case "start":
			pterodactyl.Start(serverConf.ServerID)
		case "stop":
			pterodactyl.Stop(serverConf.ServerID)
		case "restart":
			pterodactyl.Restart(serverConf.ServerID)

		}
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
		log.Printf("Failed send power_%v Embed: %v", action, err)
	}
}

func powerStartSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	powerSelect(s, i, "start")
}
func powerRestartSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	powerSelect(s, i, "restart")

}
func powerStopSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	powerSelect(s, i, "stop")

}
