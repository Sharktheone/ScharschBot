package interactions

import (
	"github.com/Sharktheone/ScharschBot/discord/interactions/cmds"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		cmds.WhitelistCommand,
		cmds.AdminCommand,
		cmds.PowerCommand,
		cmds.HelpCommand,
	}
)
