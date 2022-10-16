package embed

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/bwmarrin/discordgo"
)

var (
	config       = conf.GetConf()
	maxAccounts  = config.Whitelist.MaxAccounts
	ErrorIcon    = config.Discord.EmbedErrorIcon
	ErrorURL     = config.Discord.EmbedErrorAuthorURL
	BotAvatarURL string
	bansToMax    = config.Whitelist.BannedUsersToMaxAccounts
)

func WhitelistAdding(PlayerName string, Players []string, banedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is now on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistAlreadyListed(PlayerName string, Players []string, banedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is already on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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
func WhitelistNotExisting(PlayerName string, Players []string, banedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is not existing", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistNoFreeAccounts(PlayerName string, Players []string, banedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := "You have no free Accounts"
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistAddNotAllowed(PlayerName string, Players []string, bannedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to add %v to the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistRemoving(PlayerName string, Players []string, banedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is no longer on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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
func WhitelistRemoveNotAllowed(PlayerName string, Players []string, bannedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to remove %v from the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistNotListed(PlayerName string, Players []string, banedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is not on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistIsListedBy(PlayerName string, userID string, Players []string, bannedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v was whitelisted by", PlayerName)
	Description := fmt.Sprintf("<@%v>", userID)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistWhoisNotAllowed(PlayerName string, Players []string, bannedPlayers []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to lookup the owner of %v", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • You have whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistHasListed(PlayerNames []string, userID string, avatarURL string, name string, bannedPlayers []string, footerIcon bool, username string) discordgo.MessageEmbed {
	Title := "Whitelisted accounts of"
	Description := fmt.Sprintf("<@%v>", userID)
	var embedAccounts []*discordgo.MessageEmbedField

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
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
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
			Name:    name,
			IconURL: avatarURL,
		},
		Fields: embedAccounts,
		Footer: Footer,
	}

	return Embed
}

func WhitelistNoAccounts(userID string, avatarURL string, name string) discordgo.MessageEmbed {
	Title := "The following user has no whitelisted accounts:"
	Description := fmt.Sprintf("<@%v>", userID)
	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatarURL,
		},
	}
	return Embed
}

func WhitelistUserNotAllowed(userID string, avatarURL string, name string, Players []string, bannedPlayers []string, footerIcon bool, username string) discordgo.MessageEmbed {
	Title := "You have no permission to lookup the whitelisted accounts of"
	Description := fmt.Sprintf("<@%v>", userID)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
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
			Name:    name,
			IconURL: avatarURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistRemoveAllNotAllowed(avatarURL string, name string, Players []string, bannedPlayers []string, footerIcon bool, username string) discordgo.MessageEmbed {
	Title := "You have no permission to remove all accounts from the whitelist"
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
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
			Name:    name,
			IconURL: avatarURL,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistRemoveAllNoWhitelistEntries(avatar string, name string) discordgo.MessageEmbed {
	Title := "There is no whitelist entries to remove"

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatar,
			URL:     ErrorURL,
		},
	}
	return Embed
}

func WhitelistRemoveAllSure(avatar string, name string) (embed discordgo.MessageEmbed, button discordgo.Button) {
	Title := "Do you really want to remove all accounts from the whitelist?"

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF9900,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatar,
			URL:     ErrorURL,
		},
	}
	Button := discordgo.Button{
		Emoji: discordgo.ComponentEmoji{
			Name: "✅",
		},
		Label:    "Yes, I want to remove all accounts",
		CustomID: "remove_yes",
		Style:    discordgo.DangerButton,
	}
	return Embed, Button
}
func WhitelistRemoveAll(avatar string, name string) discordgo.MessageEmbed {
	Title := "You have successful removed all accounts from the whitelist"

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatar,
			URL:     ErrorURL,
		},
	}
	return Embed
}

func WhitelistBanUserID(PlayerNames []string, userID string, avatarURL string, name string, bannedAccounts []string, footerIcon bool, username string, reason string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("Banning following user for the reason %v that has following whitelisted accounts", username)
	Description := fmt.Sprintf("<@%v>", userID)
	var embedAccounts []*discordgo.MessageEmbedField
	var Footer *discordgo.MessageEmbedFooter
	var FooterText string

	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts (max %v)", username, len(PlayerNames), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts and %v banned (max %v)", username, len(PlayerNames), len(bannedAccounts), maxAccounts)
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

	for _, PlayerName := range PlayerNames {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
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
			Name:    name,
			IconURL: avatarURL,
		},
		Fields: Fields,
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanAccount(PlayerName string, Players []string, userID string, bannedAccounts []string, footerIcon bool, footerIconURL string, username string, reason string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is now banned from the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer *discordgo.MessageEmbedFooter
	var FooterText string
	var field discordgo.MessageEmbedField
	var reasonField = discordgo.MessageEmbedField{
		Name:  "Reason",
		Value: reason,
	}
	if len(userID) > 0 {
		FieldName := fmt.Sprintf("%v was whitelisted by", PlayerName)
		FieldValue := fmt.Sprintf("<@%v>", userID)
		field = discordgo.MessageEmbedField{
			Name:  FieldName,
			Value: FieldValue,
		}
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • He had whitelisted %v accounts and %v banned (max %v)", username, len(Players), len(bannedAccounts), maxAccounts)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: footerIconURL,
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

func WhitelistUnBanUserID(userID string, avatarURL string, name string, bannedAccounts []string, listedAccounts []string, footerIcon bool, username string) discordgo.MessageEmbed {
	Title := "Unbanning user"
	Description := fmt.Sprintf("<@%v>", userID)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if !bansToMax {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts (max %v)", username, len(listedAccounts), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • He has whitelisted %v accounts and %v are banned (max %v)", username, len(listedAccounts), len(bannedAccounts), maxAccounts)
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
	for _, PlayerName := range listedAccounts {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
			Value:  userURL,
			Inline: false,
		})
	}
	for _, PlayerName := range bannedAccounts {
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
			Name:    name,
			IconURL: avatarURL,
		},
		Footer: Footer,
		Fields: embedAccounts,
	}
	return Embed
}

func WhitelistUnBanAccount(PlayerName string, Players []string, footerIcon bool, footerIconURL string, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is now unbanned from the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
	if len(Players) > 0 {
		FooterText = fmt.Sprintf("%v • Account owner has whitelisted now %v accounts (max %v)", username, len(Players), maxAccounts)
	} else {
		FooterText = fmt.Sprintf("%v • No was not whitelisted", username)
	}
	if footerIcon {
		Footer = &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: footerIconURL,
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

func WhitelistBanAccountNotAllowed(avatarURL string, name string, mcName string, Players []string, bannedPlayers []string, footerIcon bool, username string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to (un)ban %v", mcName)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
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
			Name:    name,
			IconURL: avatarURL,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanUserIDNotAllowed(avatarURL string, name string, banID string, Players []string, bannedPlayers []string, footerIcon bool, username string) discordgo.MessageEmbed {
	Title := "You have no permission to (un)ban"
	Description := fmt.Sprintf("<@%v>", banID)
	var FooterText string
	var Footer *discordgo.MessageEmbedFooter
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
			Name:    name,
			IconURL: avatarURL,
			URL:     ErrorURL,
		},
		Footer: Footer,
	}
	return Embed
}

func WhitelistBanned(avatar string, name string, mcName string, IDBan bool, reason string) discordgo.MessageEmbed {
	var (
		Title         string
		AuthorName    string
		AuthorURL     string
		AuthorIconURL string
		Description   = fmt.Sprintf("Reason: %v", reason)
	)
	if IDBan {
		Title = " You have no permission to whitelist accounts because you are banned from the whitelist"
		AuthorName = name
		AuthorURL = ErrorURL
		AuthorIconURL = avatar
	} else {
		Title = fmt.Sprintf("You have no permission to add %v to the whitelist beacause the account banned from the whitelist", mcName)
		AuthorName = mcName
		AuthorURL = fmt.Sprintf("https://namemc.com/profile/%v", mcName)
		AuthorIconURL = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", mcName)
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
func WhitelistRemoveMyAccounts(PlayerNames []string, userID string, avatarURL string, name string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := "Removing whitelisted accounts of"
	Description := fmt.Sprintf("<@%v>", userID)
	var embedAccounts []*discordgo.MessageEmbedField

	for _, PlayerName := range PlayerNames {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
			Value:  userURL,
			Inline: false,
		})
	}
	var Footer *discordgo.MessageEmbedFooter
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
				Name:    name,
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
				Name:    name,
				IconURL: avatarURL,
			},
			Fields: embedAccounts,
		}
	}

	return Embed
}
