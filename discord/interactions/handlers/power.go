package handlers

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/embed/srvEmbed"
	"Scharsch-Bot/pterodactyl"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func PowerMain(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, i.ApplicationCommandData().Options[0].Name)
}
func PowerStart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, pterodactyl.PowerSignalStart)
}
func PowerRestart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, pterodactyl.PowerSignalRestart)
}
func PowerStop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	power(s, i, pterodactyl.PowerSignalStop)
}

func power(s *discordgo.Session, i *discordgo.InteractionCreate, option string) {
	var (
		restartDisabled = false
		stopDisabled    = false
		startDisabled   = false
	)

	serverSelect := discordgo.SelectMenu{
		Placeholder: fmt.Sprintf("Select a server to %s", option),
		CustomID:    fmt.Sprintf("power_%s_select", option),
		Options:     getServerOptions(option),
	}

	switch option {
	case pterodactyl.PowerSignalStart:
		startDisabled = true
	case pterodactyl.PowerSignalStop:
		stopDisabled = true
	case pterodactyl.PowerSignalRestart:
		restartDisabled = true
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

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				srvEmbed.Power(option),
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: buttonRow,
				},
				func(option string) discordgo.ActionsRow {
					if option != "status" {
						return discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								&serverSelect,
							},
						}
					}
					return discordgo.ActionsRow{}
				}(option),
			},
		},
	}); err != nil {
		log.Printf("Failed to respond to command power: %v", err)
	}
}

func getServerOptions(option string) []discordgo.SelectMenuOption {
	var fields []discordgo.SelectMenuOption
	for _, server := range pterodactyl.Servers {
		if option == pterodactyl.PowerSignalStart && server.Status.State == pterodactyl.PowerStatusRunning ||
			option == pterodactyl.PowerSignalStop && server.Status.State == pterodactyl.PowerStatusOffline {
			continue
		}
		fields = append(fields, discordgo.SelectMenuOption{
			Label: server.Config.ServerName,
			Value: server.Config.ServerID,
		})
	}
	return fields
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
		messageEmbed = srvEmbed.PowerNotAllowed(i.Member.User.AvatarURL("40"), i.Member.User.String(), action, serverConf.ServerName)
	} else {
		messageEmbed = srvEmbed.PowerAction(action, serverConf.ServerName, i.Member.User.AvatarURL("40"), i.Member.User.Username)
		s, err := pterodactyl.GetServer(serverConf.ServerID)
		if err != nil {
			log.Printf("Failed to get server %s: %v", serverConf.ServerName, err)
			return
		}
		switch action {
		case "start":
			if err := s.Start(); err != nil {
				log.Printf("Failed to start server %s: %v", serverConf.ServerName, err)
			}
		case "stop":
			if err := s.Stop(); err != nil {
				log.Printf("Failed to stop server %s: %v", serverConf.ServerName, err)
			}
		case "restart":
			if err := s.Restart(); err != nil {
				log.Printf("Failed to restart server %s: %v", serverConf.ServerName, err)
			}

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
