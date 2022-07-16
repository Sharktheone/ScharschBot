package commands

import "github.com/bwmarrin/discordgo"

var (
	DefaultPermission = true
	Commands          = []*discordgo.ApplicationCommand{
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
)
