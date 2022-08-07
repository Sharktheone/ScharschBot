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
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "User to lookup accounts (You can list left users with <@USERID>)",
					Required:    true,
				},
			},
		},
		{
			Name:              "whitelistbanuser",
			Description:       "Ban a user from the whitelist",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "User to ban from the whitelist (You can ban left users with <@USERID>)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "removeaccounts",
					Description: "Remove the accounts of the banned person (default: true)",
					Required:    false,
				},
			},
		},
		{
			Name:              "whitelistbanaccount",
			Description:       "Ban a user from the whitelist by minecraft account name",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the account to ban",
					Required:    true,
				},
			},
		},
		{
			Name:              "whitelistunbanuser",
			Description:       "Unban a user from the whitelist",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "User to unban from the whitelist (You can unban left users with <@USERID>)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "unbanaccounts",
					Description: "Unban the accounts of the unbanned person (default: false)",
					Required:    false,
				},
			},
		},
		{
			Name:              "whitelistunbanaccount",
			Description:       "Unban a user from the whitelist by minecraft account name",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the account to unban",
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
			Name:              "whitelistremovemyaccounts",
			Description:       "Remove all accounts you have whitelisted",
			DefaultPermission: &DefaultPermission,
		},
		{
			Name:              "whitelistremoveall",
			Description:       "Remove all accounts from whitelist",
			DefaultPermission: &DefaultPermission,
		},
	}
)
