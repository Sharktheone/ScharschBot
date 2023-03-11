package handlers

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/embed"
	"Scharsch-Bot/pterodactyl"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
)

func PowerMain(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, i.ApplicationCommandData().Options[0].Name)
}
func PowerStart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, "start")
}
func PowerRestart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, "restart")
}
func PowerStop(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

func PowerStartSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	powerSelect(s, i, "start")
}
func PowerRestartSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	powerSelect(s, i, "restart")

}
func PowerStopSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	powerSelect(s, i, "stop")

}
