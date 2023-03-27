package cmds

import "github.com/bwmarrin/discordgo"

var (
	HelpCommand = &discordgo.ApplicationCommand{
		Name:              "help",
		Description:       "Get help",
		DefaultPermission: &DefaultPermission,
	}
)
