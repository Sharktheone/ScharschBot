package banembed

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/discordMember"
	"github.com/bwmarrin/discordgo"
)

var (
	config     = conf.GetConf()
	serverName = config.Discord.ServerName
)

func DMBan(dmFailed bool, userID string, reason string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		user, _     = discordMember.GetUserProfile(userID, s)
		authorName  = user.User.String()
		avatarURL   = user.AvatarURL("40")
		Title       = fmt.Sprintf("You got banned on the server %v", serverName)
		Description = fmt.Sprintf("You have been banned for the reason %v from the server. If you think this is a mistake, please contact a staff member directly.", reason)
		FooterText  string
		FooterIcon  = user.AvatarURL("40")
	)

	if dmFailed {
		FooterText = fmt.Sprintf("Failed to send DM to %v. Maybe you have DMs disabled? Sending to channel instead.", user.User.String())
	} else {
		FooterText = authorName
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FFC9,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    authorName,
			IconURL: avatarURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: FooterIcon,
		},
	}
	return Embed
}

func DMUnBan(dmFailed bool, userID string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		user, _     = discordMember.GetUserProfile(userID, s)
		authorName  = user.User.String()
		avatarURL   = user.AvatarURL("40")
		Title       = fmt.Sprintf("You got Unbanned on the server %v", serverName)
		Description string
		FooterText  string
		FooterIcon  = user.AvatarURL("40")
	)

	if dmFailed {
		FooterText = fmt.Sprintf("Failed to send DM to %v. Maybe you have DMs disabled? Sending to channel instead.", user.User.String())
	} else {
		FooterText = authorName
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FFC9,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    authorName,
			IconURL: avatarURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: FooterIcon,
		},
	}
	return Embed
}

func DMBanAccount(name string, dmFailed bool, userID string, reason string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		user, _     = discordMember.GetUserProfile(userID, s)
		authorName  = user.User.String()
		avatarURL   = user.AvatarURL("40")
		Title       = fmt.Sprintf("Your account %v got banned on the server %v", name, serverName)
		Description = fmt.Sprintf("The Account has been banned for the reason %v from the server. If you think this is a mistake, please contact a staff member directly.", reason)
		FooterText  string
		FooterIcon  = user.AvatarURL("40")
	)

	if dmFailed {
		FooterText = fmt.Sprintf("Failed to send DM to %v. Maybe you have DMs disabled? Sending to channel instead.", user.User.String())
	} else {
		FooterText = authorName
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FFC9,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    authorName,
			IconURL: avatarURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: FooterIcon,
		},
	}
	return Embed
}

func DMUnBanAccount(name string, dmFailed bool, userID string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		user, _     = discordMember.GetUserProfile(userID, s)
		authorName  = user.User.String()
		avatarURL   = user.AvatarURL("40")
		Title       = fmt.Sprintf("Your account %v got Unbanned on the server %v", name, serverName)
		Description string
		FooterText  string
		FooterIcon  = user.AvatarURL("40")
	)

	if dmFailed {
		FooterText = fmt.Sprintf("Failed to send DM to %v. Maybe you have DMs disabled? Sending to channel instead.", user.User.String())
	} else {
		FooterText = authorName
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FFC9,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    authorName,
			IconURL: avatarURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: FooterIcon,
		},
	}
	return Embed
}
