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

func WhitelistAdding(PlayerName string, Players []string, banedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is now on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(banedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistAlreadyListed(PlayerName string, Players []string, banedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is already on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(banedPlayers), maxAccounts)
	}
	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF9900,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed

}
func WhitelistNotExisting(PlayerName string, Players []string, banedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is not existing", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(banedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: ErrorIcon,
			URL:     ErrorURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistNoFreeAccounts(PlayerName string, Players []string, banedPlayers []string) discordgo.MessageEmbed {
	Title := "You have no free Accounts"
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(banedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistAddNotAllowed(PlayerName string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to add %v to the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
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

func WhitelistRemoving(PlayerName string, Players []string, banedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is no longer on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(banedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}
func WhitelistRemoveNotAllowed(PlayerName string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to remove %v from the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistNotListed(PlayerName string, Players []string, banedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is not on the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(banedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF9900,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistIsListedBy(PlayerName string, userID string, Players []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v was whitelisted by", PlayerName)
	Description := fmt.Sprintf("<@%v>", userID)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	Footer := fmt.Sprintf("He has whitelisted %v accounts (max %v)", len(Players), maxAccounts)

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistWhoisNotAllowed(PlayerName string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to lookup the owner of %v", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistHasListed(PlayerNames []string, userID string, avatarURL string, name string, bannedPlayers []string) discordgo.MessageEmbed {
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
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("He has whitelisted %v accounts (max %v)", len(PlayerNames), maxAccounts)
	} else {
		Footer = fmt.Sprintf("He has whitelisted %v accounts and %v are banned (max %v)", len(PlayerNames), len(bannedPlayers), maxAccounts)
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
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
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

func WhitelistUserNotAllowed(userID string, avatarURL string, name string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := "You have no permission to lookup the whitelisted accounts of"
	Description := fmt.Sprintf("<@%v>", userID)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatarURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistRemoveAllNotAllowed(avatar string, name string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := "You have no permission to remove all accounts from the whitelist"
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatar,
			URL:     ErrorURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
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

func WhitelistBanUserID(PlayerNames []string, userID string, avatarURL string, name string, bannedAccounts []string) discordgo.MessageEmbed {
	Title := "Banning following user that has following whitelisted accounts"
	Description := fmt.Sprintf("<@%v>", userID)
	var embedAccounts []*discordgo.MessageEmbedField
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("He had whitelisted %v accounts (max %v)", len(PlayerNames), maxAccounts)
	} else {
		Footer = fmt.Sprintf("He had whitelisted %v accounts and %v banned (max %v)", len(PlayerNames), len(bannedAccounts), maxAccounts)
	}

	for _, PlayerName := range PlayerNames {
		userURL := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		embedAccounts = append(embedAccounts, &discordgo.MessageEmbedField{
			Name:   PlayerName,
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
		Fields: embedAccounts,
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistBanAccount(PlayerName string, Players []string, userID string, bannedAccounts []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is now banned from the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	var field discordgo.MessageEmbedField
	if len(userID) > 0 {
		FieldName := fmt.Sprintf("%v was whitelisted by", PlayerName)
		FieldValue := fmt.Sprintf("<@%v>", userID)
		field = discordgo.MessageEmbedField{
			Name:  FieldName,
			Value: FieldValue,
		}
		if !bansToMax {
			Footer = fmt.Sprintf("He had whitelisted %v accounts (max %v)", len(Players), maxAccounts)
		} else {
			Footer = fmt.Sprintf("He had whitelisted %v accounts and %v banned (max %v)", len(Players), len(bannedAccounts), maxAccounts)
		}
	} else {
		FieldName := fmt.Sprintf("%v is not whitelisted", PlayerName)
		field = discordgo.MessageEmbedField{
			Name:  FieldName,
			Value: "The ban will be executed",
		}
	}
	var Embed discordgo.MessageEmbed
	if len(Footer) > 0 {
		Embed = discordgo.MessageEmbed{
			Title: Title,
			Color: 0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    PlayerName,
				IconURL: AuthorIconUrl,
				URL:     AuthorUrl,
			},
			Fields: []*discordgo.MessageEmbedField{
				&field,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: Footer,
			},
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
				&field,
			},
		}
	}

	return Embed
}

func WhitelistUnBanUserID(userID string, avatarURL string, name string, bannedAccounts []string, listedAccounts []string) discordgo.MessageEmbed {
	Title := "Unbanning user"
	Description := fmt.Sprintf("<@%v>", userID)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("Owner has whitelisted %v accounts (max %v)", len(listedAccounts), maxAccounts)
	} else {
		Footer = fmt.Sprintf("Owner has whitelisted %v accounts and %v banned (max %v)", len(listedAccounts), len(bannedAccounts), maxAccounts)
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
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
		Fields: embedAccounts,
	}
	return Embed
}

func WhitelistUnBanAccount(PlayerName string, Players []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("%v is now unbanned from the whitelist", PlayerName)
	AuthorIconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
	AuthorUrl := fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
	var Footer string
	if len(Players) > 0 {
		Footer = fmt.Sprintf("Account owner has whitelisted now %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = "No was not whitelisted"
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0x00FF00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    PlayerName,
			IconURL: AuthorIconUrl,
			URL:     AuthorUrl,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistBanAccountNotAllowed(avatar string, name string, mcName string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := fmt.Sprintf("You have no permission to (un)ban %v", mcName)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatar,
			URL:     ErrorURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistBanUserIDNotAllowed(avatar string, name string, banID string, Players []string, bannedPlayers []string) discordgo.MessageEmbed {
	Title := "You have no permission to (un)ban"
	Description := fmt.Sprintf("<@%v>", banID)
	var Footer string
	if !bansToMax {
		Footer = fmt.Sprintf("You have whitelisted %v accounts (max %v)", len(Players), maxAccounts)
	} else {
		Footer = fmt.Sprintf("You have whitelisted %v accounts and %v are banned (max %v)", len(Players), len(bannedPlayers), maxAccounts)
	}

	Embed := discordgo.MessageEmbed{
		Title:       Title,
		Description: Description,
		Color:       0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatar,
			URL:     ErrorURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: Footer,
		},
	}
	return Embed
}

func WhitelistBanned(avatar string, name string, mcName string, IDBan bool) discordgo.MessageEmbed {
	var (
		Title         string
		AuthorName    string
		AuthorURL     string
		AuthorIconURL string
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
		Title: Title,
		Color: 0xFF0000,
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
	var Footer string
	if len(bannedPlayers) > 0 {
		Footer = fmt.Sprintf("You have %v banned accounts (max %v)", len(bannedPlayers), maxAccounts)
	}
	var Embed discordgo.MessageEmbed
	if len(Footer) > 0 {

		Embed = discordgo.MessageEmbed{
			Title:       Title,
			Description: Description,
			Color:       0x00FF00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    name,
				IconURL: avatarURL,
			},
			Fields: embedAccounts,
			Footer: &discordgo.MessageEmbedFooter{
				Text: Footer,
			},
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
