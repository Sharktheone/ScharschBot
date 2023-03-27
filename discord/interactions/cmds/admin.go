package cmds

import "github.com/bwmarrin/discordgo"

var (
	AdminCommand = &discordgo.ApplicationCommand{
		Name:              "admin",
		Description:       "Admin cmds of the Scharsch-Bot",
		DefaultPermission: &DefaultPermission,
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "whois",
				Description: "Look which account was whitelisted by which Discord member",
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
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "user",
				Description: "Look which Discord member has whitelisted which accounts",
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
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "banuser",
				Description: "Ban a user from the whitelist",
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
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "reason",
						Description: "Reason for the ban",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "banaccount",
				Description: "Ban a user from the whitelist by minecraft account name",
				Options: []*discordgo.ApplicationCommandOption{

					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the account to ban",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "reason",
						Description: "Reason for the ban",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "unbanuser",
				Description: "Unban a user from the whitelist",
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
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "unbanaccount",
				Description: "Unban a user from the whitelist by minecraft account name",
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
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "removeall",
				Description: "Remove all accounts from whitelist",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "listreports",
				Description: "List all reports",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "rejectreport",
				Description: "Reject a report",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the reported Player",
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "notifyreporter",
						Description: "Notify the reporter about the rejection (default: true)",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "acceptreport",
				Description: "Accept a report",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name of the reported Player",
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        "notifyreporter",
						Description: "Notify the reporter about the acceptance (default: true)",
					},
				},
			},
		},
	}
)
