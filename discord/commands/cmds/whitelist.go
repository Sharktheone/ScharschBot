package cmds

import "github.com/bwmarrin/discordgo"

var (
	DefaultPermission = true
	WhitelistCommand  = discordgo.ApplicationCommand{
		Name:              "whitelist",
		Description:       "User cmds of the Scharsch-Bot",
		DefaultPermission: &DefaultPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "add",
				Description: "Add your account to the Whitelist",
				Options: []*discordgo.ApplicationCommandOption{

					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the account to add",
						Required:    true,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "remove",
				Description: "Remove your account to the Whitelist",
				Options: []*discordgo.ApplicationCommandOption{

					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the account to remove",
						Required:    true,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "myaccounts",
				Description: "Look which accounts you have whitelisted",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "removemyaccounts",
				Description: "Remove all accounts you have whitelisted",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "report",
				Description: "Report a Player or Discord User",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Player to report",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "Discord User to report",
						Required:    false,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "reason",
						Description: "Reason for the report",
						Required:    false,
					},
				},
			},
		},
	}
)
