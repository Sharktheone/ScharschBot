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
		{
			Name:              "whitelistremove",
			Description:       "Remove your account to the Whitelist",
			DefaultPermission: &DefaultPermission,
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
			Name:              "whitelistwhois",
			Description:       "Look which account was whitelisted by which Discord member",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the account to lookup",
					Required:    true,
				},
			},
		},
		{
			Name:              "whitelistuser",
			Description:       "Look which Discord member has whitelisted which accounts",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "userid",
					Description: "ID of the account to lookup",
					Required:    true,
				},
			},
		},
		{
			Name:              "whitelistmyaccounts",
			Description:       "Look which accounts you have whitelisted",
			DefaultPermission: &DefaultPermission,
		},
		{
			Name:              "whitelistremoveall",
			Description:       "Remove all accounts from whitelist",
			DefaultPermission: &DefaultPermission,
		},
	}
)
