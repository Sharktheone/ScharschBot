package commands

import (
	"Scharsch-Bot/discord/commands/handlers"
	"github.com/bwmarrin/discordgo"
)

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"whitelist":            handlers.Whitelist,
	"admin":                handlers.Admin,
	"remove_yes":           handlers.RemoveYes,
	"power":                handlers.PowerMain,
	"power_start":          handlers.PowerStart,
	"power_restart":        handlers.PowerRestart,
	"power_stop":           handlers.PowerStop,
	"power_start_select":   handlers.PowerStartSelect,
	"power_restart_select": handlers.PowerRestartSelect,
	"power_stop_select":    handlers.PowerStopSelect,
}
