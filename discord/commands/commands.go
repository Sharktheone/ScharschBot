package commands

import (
	"Scharsch-Bot/discord/commands/cmds"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		&cmds.WhitelistCommand,
		&cmds.AdminCommand,
		&cmds.PowerCommand,
		&cmds.HelpCommand,
	}
)
