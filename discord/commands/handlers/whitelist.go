package handlers

import (
	"Scharsch-Bot/database/mongodb"
	"Scharsch-Bot/discord/embed/wEmbed"
	"Scharsch-Bot/reports"
	"Scharsch-Bot/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func Whitelist(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options[0].Options {
		optionMap[opt.Name] = opt
	}
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
			listedAccounts := whitelist.ListedAccountsOf(i.Member.User.ID, false)
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
								messageEmbed = wEmbed.WhitelistAlreadyListed(name, i)
							} else {
								messageEmbed = wEmbed.WhitelistAdding(name, i)
							}
						} else {
							messageEmbed = wEmbed.WhitelistNotExisting(name, i)
						}
					} else {
						messageEmbed = wEmbed.WhitelistNoFreeAccounts(name, i)
						noFree = true
					}
				} else {
					messageEmbed = wEmbed.WhitelistAddNotAllowed(name, i)

				}
			} else {
				messageEmbed = wEmbed.WhitelistBanned(name, dcBan, banReason, i)
			}
		} else {
			messageEmbed = wEmbed.DatabaseNotReady
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

			if allowed {
				if onWhitelist {
					messageEmbed = wEmbed.WhitelistRemoving(name, i)
				} else {
					messageEmbed = wEmbed.WhitelistNotListed(name, i)
				}
			} else {
				messageEmbed = wEmbed.WhitelistRemoveNotAllowed(name, i)
			}
		} else {
			messageEmbed = wEmbed.DatabaseNotReady
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
					messageEmbed = wEmbed.WhitelistHasListed(accounts, i.Member.User.ID, bannedPlayers, i, s)
				} else {
					messageEmbed = wEmbed.WhitelistNoAccounts(i, i.Member.User.ID)
				}
			} else {
				messageEmbed = wEmbed.WhitelistUserNotAllowed(accounts, i.Member.User.ID, bannedPlayers, i)
			}
		} else {
			messageEmbed = wEmbed.DatabaseNotReady
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
				messageEmbed = wEmbed.WhitelistRemoveMyAccounts(listedAccounts, mcBans, i)
			} else {
				messageEmbed = wEmbed.WhitelistNoAccounts(i, i.Member.User.ID)
			}
		} else {
			messageEmbed = wEmbed.DatabaseNotReady
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
	case "report":
		var (
			messageEmbed discordgo.MessageEmbed
			name         string
			reason       = "No reason provided"
		)
		if optionMap["name"] != nil {
			name = strings.ToLower(optionMap["name"].StringValue())
		}
		if optionMap["reason"] != nil {
			reason = optionMap["reason"].StringValue()
		}

		if mongodb.Ready {
			reportEmbed := wEmbed.NewReport(name, reason, i)
			allowed, alreadyReportet, enabled := reports.Report(name, reason, i, s, reportEmbed)
			if allowed {
				if enabled {
					if alreadyReportet {
						messageEmbed = wEmbed.AlreadyReported(name)
					} else {
						messageEmbed = wEmbed.ReportPlayer(name, reason, i)
					}
				} else {
					messageEmbed = wEmbed.ReportDisabled(i)
				}
			} else {
				messageEmbed = wEmbed.ReportNotALlowed(i)
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
				log.Printf("Failed execute command report %v", err)
			}
			log.Printf("%v reported %v for %v", i.Member.User.Username, name, reason)

		}
	}

}
