package commands

import "github.com/bwmarrin/discordgo"

var (
	DefaultPermission = true
	Commands          = []*discordgo.ApplicationCommand{
		{
			Name:              "whitelist",
			Description:       "User commands of the Scharsch-Bot",
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
			},
		},
		{
			Name:              "whitelistadmin",
			Description:       "Admin commands of the Scharsch-Bot",
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
			},
		},
		{
			Name:              "power",
			Description:       "power commands of the Scharsch-Bot",
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
		},
		{
			Name:              "report",
			Description:       "Report a Player or Discord User",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "Discord User to report",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Player to report",
				},
			},
		},
		{
			Name:              "reportadmin",
			Description:       "Report Admin Commands",
			DefaultPermission: &DefaultPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "list",
					Description: "List all reports",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "reject",
					Description: "Reject a report",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "name",
							Description: "Name of the reported Player",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "accept",
					Description: "Accept a report",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "name",
							Description: "Name of the reported Player",
						},
					},
				},
			},
		},
		{
			Name:              "help",
			Description:       "Get help",
			DefaultPermission: &DefaultPermission,
		},
	}
)
