package embed

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/discordMember"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
)

var (
	config       = conf.GetConf()
	ErrorIcon    = config.Discord.EmbedErrorIcon
	ErrorURL     = config.Discord.EmbedErrorAuthorURL
	BotAvatarURL string
	bansToMax    = config.Whitelist.BannedUsersToMaxAccounts
	footerIcon   = config.Discord.FooterIcon
)

func WhitelistAdding(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("%v is now on the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistAlreadyListed(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("%v is already on the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}
	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF9900,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed

}
func WhitelistNotExisting(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {

	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("%v is not existing", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: ErrorIcon,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistNoFreeAccounts(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = "You have no free Accounts"
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistAddNotAllowed(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("You have no permission to add %v to the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

var DatabaseNotReady = discordgo.MessageEmbed{
	Title: "Database is not ready, please try again later",
	Color: 0xFF0000,
	Author: &discordgo.MessageEmbedAuthor{
		Name:    "Scharsch-bot",
		IconURL: BotAvatarURL,
		URL:     ErrorURL,
	},
	Footer: &discordgo.MessageEmbedFooter{
		Text: "Bot is starting database connection is not ready",
	},
}

func WhitelistRemoving(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("%v is no longer on the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}
func WhitelistRemoveNotAllowed(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("You have no permission to remove %v from the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistNotListed(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("%v is not on the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF9900,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistIsListedBy(PlayerName string, playerID string, i *discordgo.InteractionCreate, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("%v was whitelisted by", PlayerName)
		Description   = fmt.Sprintf("<@%v>", playerID)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(playerID)
		bannedPlayers = whitelist.CheckBans(playerID)
	)

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistWhoisNotAllowed(PlayerName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("You have no permission to lookup the owner of %v", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistHasListed(PlayerNames []string, playerID string, bannedPlayers []string, i *discordgo.InteractionCreate, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = "Whitelisted accounts of"
		Description   = fmt.Sprintf("<@%v>", playerID)
		embedAccounts []*discordgo.MessageEmbedField
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)

	for _, PlayerName := range PlayerNames {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
			Value:  userURL,
			Inline: false,
		})
	}
	for _, PlayerName := range bannedPlayers {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%v (banned)", PlayerName),
			Value:  userURL,
			Inline: false,
		})
	}

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts (max %v)", username, len(PlayerNames), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts and %v are banned (max %v)", username, len(PlayerNames), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
		},
		Fields: embedAccounts,
		Footer: Footer,
	}

	return Embed
}

func WhitelistNoAccounts(i *discordgo.InteractionCreate, playerID string) discordgo.MessageEmbed {
	var (
		username    = i.Member.User.String()
		avatarURL   = i.Member.User.AvatarURL("40")
		Title       = "The following user has no whitelisted accounts:"
		Description = fmt.Sprintf("<@%v>", playerID)
		Embed       = discordgo.MessageEmbed{
			Title:       Title,
			Description: Description,
			Color:       0xFF0000,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    username,
				IconURL: avatarURL,
			},
		}
	)
	return Embed
}

func WhitelistUserNotAllowed(Players []string, playerID string, bannedPlayers []string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username    = i.Member.User.String()
		avatarURL   = i.Member.User.AvatarURL("40")
		maxAccounts = whitelist.GetMaxAccounts(i.Member.Roles)
		Title       = "You have no permission to lookup the whitelisted accounts of"
		Description = fmt.Sprintf("<@%v>", playerID)
		FooterText  string
		Footer      *discordgo.MessageEmbedFooter
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistRemoveAllNotAllowed(i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = "You have no permission to remove all accounts from the whitelist"
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistRemoveAllNoWhitelistEntries(i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username  = i.Member.User.String()
		avatarURL = i.Member.User.AvatarURL("40")
		Title     = "There is no whitelist entries to remove"

		Embed = discordgo.MessageEmbed{
			Title: Title,
			Color: 0xFF0000,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    username,
				IconURL: avatarURL,
				URL:     ErrorURL,
			},
		}
	)
	return Embed
}

func WhitelistRemoveAllSure(i *discordgo.InteractionCreate) (embed discordgo.MessageEmbed, button discordgo.Button) {
	var (
		username  = i.Member.User.String()
		avatarURL = i.Member.User.AvatarURL("40")
		Title     = "Do you really want to remove all accounts from the whitelist?"

		Embed = discordgo.MessageEmbed{
			Title: Title,
			Color: 0xFF9900,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    username,
				IconURL: avatarURL,
				URL:     ErrorURL,
			},
		}
		Button = discordgo.Button{
			Emoji: discordgo.ComponentEmoji{
				Name: "✅",
			},
			Label:    "Yes, I want to remove all accounts",
			CustomID: "remove_yes",
			Style:    discordgo.DangerButton,
		}
	)
	return Embed, Button
}
func WhitelistRemoveAll(i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username  = i.Member.User.String()
		avatarURL = i.Member.User.AvatarURL("40")
		Title     = "You have successful removed all accounts from the whitelist"

		Embed = discordgo.MessageEmbed{
			Title: Title,
			Color: 0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    username,
				IconURL: avatarURL,
				URL:     ErrorURL,
			},
		}
	)
	return Embed
}

func WhitelistBanUserID(playerID string, reason string, i *discordgo.InteractionCreate, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("Banning following user for the reason %v that has following whitelisted accounts", username)
		Description   = fmt.Sprintf("<@%v>", playerID)
		embedAccounts []*discordgo.MessageEmbedField
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(playerID)
		bannedPlayers = whitelist.CheckBans(playerID)
	)
	var FooterText string

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts and %v banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	for _, Player := range Players {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", Player)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   Player,
			Value:  userURL,
			Inline: false,
		})
	}
	var Fields []*discordgo.MessageEmbedField
	Fields = append(Fields, &discordgo.MessageEmbedField{
		Name:  "Reason",
		Value: reason,
	})
	Fields = append(Fields, embedAccounts...)

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
		},
		Fields: Fields,
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanAccount(PlayerName string, playerID string, reason string, i *discordgo.InteractionCreate, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Players       = whitelist.ListedAccountsOf(playerID)
		bannedPlayers = whitelist.CheckBans(playerID)
		Title         = fmt.Sprintf("%v is now banned from the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		Footer        *discordgo.MessageEmbedFooter
		FooterText    string
		field         discordgo.MessageEmbedField
		reasonField   = discordgo.MessageEmbedField{
			Name:  "Reason",
			Value: reason,
		}
		userID = i.Member.User.ID
	)
	if len(userID) > 0 {
		FieldName := fmt.Sprintf("%v was whitelisted by", PlayerName)
		FieldValue := fmt.Sprintf("<@%v>", playerID)
		field = discordgo.MessageEmbedField{
			Name:  FieldName,
			Value: FieldValue,
		}
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts and %v banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: avatarURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	} else {
		FieldName := fmt.Sprintf("%v is not whitelisted", PlayerName)
		field = discordgo.MessageEmbedField{
			Name:  FieldName,
			Value: "The ban will be executed",
		}
	}
	var Embed discordgo.MessageEmbed
	if len(FooterText) > 0 {
		Embed = discordgo.MessageEmbed{
			Title: Title,
			Color: 0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    PlayerName,
				IconURL: AuthorIconUrl,
				URL:     AuthorUrl,
			},
			Fields: []*discordgo.MessageEmbedField{
				&reasonField,
				&field,
			},
			Footer: Footer,
		}
	} else {
		Embed = discordgo.MessageEmbed{
			Title: Title,
			Color: 0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    PlayerName,
				IconURL: AuthorIconUrl,
				URL:     AuthorUrl,
			},
			Fields: []*discordgo.MessageEmbedField{
				&reasonField,
				&field,
			},
		}
	}

	return Embed
}

func WhitelistUnBanUserID(playerID string, i *discordgo.InteractionCreate, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = "Unbanning user"
		Description   = fmt.Sprintf("<@%v>", playerID)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(playerID)
		bannedPlayers = whitelist.CheckBans(playerID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(Players), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}
	var embedAccounts []*discordgo.MessageEmbedField
	for _, PlayerName := range Players {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
			Value:  userURL,
			Inline: false,
		})
	}
	for _, PlayerName := range bannedPlayers {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%v (banned)", PlayerName),
			Value:  userURL,
			Inline: false,
		})
	}
	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
		},
		Footer: Footer,
		Fields: embedAccounts,
	}
	return Embed
}

func WhitelistUnBanAccount(PlayerName string, i *discordgo.InteractionCreate, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		playerID, _   = whitelist.GetOwner(PlayerName)
		Players       = whitelist.ListedAccountsOf(playerID)
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("%v is now unbanned from the whitelist", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)
	if len(Players) > 0 {
		FooterText = fmt.Sprintf("%v • He had whitelisted now %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • No was not whitelisted", username)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanAccountNotAllowed(mcName string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = fmt.Sprintf("You have no permission to (un)ban %v", mcName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanUserIDNotAllowed(playerID string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = "You have no permission to (un)ban"
		Description   = fmt.Sprintf("<@%v>", playerID)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
		Players       = whitelist.ListedAccountsOf(i.Member.User.ID)
		bannedPlayers = whitelist.CheckBans(i.Member.User.ID)
	)
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: avatarURL,
		}
	} else {
		Footer = &discordgo.MessageEmbedFooter{
			Text: FooterText,
		}
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    username,
			IconURL: avatarURL,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanned(PlayerName string, IDBan bool, reason string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		Title         string
		AuthorName    string
		AuthorURL     string
		AuthorIconURL string
		Description   = fmt.Sprintf("Reason: %v", reason)
	)
	if IDBan {
		Title = " You have no permission to whitelist accounts because you are banned from the whitelist"
		AuthorName = username
		AuthorURL = ErrorURL
		AuthorIconURL = avatarURL
	} else {
		Title = fmt.Sprintf("You have no permission to add %v to the whitelist beacause the account banned from the whitelist", PlayerName)
		AuthorName = PlayerName
		AuthorURL = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		AuthorIconURL = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    AuthorName,
			IconURL: AuthorIconURL,
			URL:     AuthorURL,
		},
	}
	return Embed

}
func WhitelistRemoveMyAccounts(PlayerNames []string, bannedPlayers []string, i *discordgo.InteractionCreate) discordgo.MessageEmbed {
	var (
		username      = i.Member.User.String()
		avatarURL     = i.Member.User.AvatarURL("40")
		maxAccounts   = whitelist.GetMaxAccounts(i.Member.Roles)
		Title         = "Removing whitelisted accounts of"
		playerID      = i.Member.User.ID
		Description   = fmt.Sprintf("<@%v>", playerID)
		embedAccounts []*discordgo.MessageEmbedField
		Footer        *discordgo.MessageEmbedFooter
	)

	for _, PlayerName := range PlayerNames {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
			Value:  userURL,
			Inline: false,
		})
	}

	var FooterText string
	if len(bannedPlayers) > 0 {
		FooterText = fmt.Sprintf("You have %v banned accounts (max %v)", len(bannedPlayers), maxAccounts)
	}
	Footer = &discordgo.MessageEmbedFooter{
		Text:    FooterText,
		IconURL: avatarURL,
	}
	var Embed discordgo.MessageEmbed
	if len(FooterText) > 0 {

		Embed = discordgo.MessageEmbed{
			Title:       Title,
			Description: Description,
			Color:       0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    username,
				IconURL: avatarURL,
			},
			Fields: embedAccounts,
			Footer: Footer,
		}
	} else {
		Embed = discordgo.MessageEmbed{
			Title:       Title,
			Description: Description,
			Color:       0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    username,
				IconURL: avatarURL,
			},
			Fields: embedAccounts,
		}
	}

	return Embed
}
