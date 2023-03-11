package cmds

import "github.com/bwmarrin/discordgo"

var (
	PowerCommand = discordgo.ApplicationCommand{
		Name:              "power",
		Description:       "power cmds of the Scharsch-Bot",
		DefaultPermission: &DefaultPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "start",
				Description: "Start a server(s)",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the server(s) to start (Comma seperated or 'all', none sends a list)",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "stop",
				Description: "Stop a server(s)",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the server(s) to stop (Comma seperated or 'all', none sends a list)",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "restart",
				Description: "Restart a server(s) ",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the server(s) to restart (Comma seperated or 'all', none sends a list)",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "status",
				Description: "Get the status of all servers",
			},
		},
	}
)
