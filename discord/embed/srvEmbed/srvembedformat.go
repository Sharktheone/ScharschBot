package srvEmbed

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/session"
	"Scharsch-Bot/minecraft/advancements"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config     = conf.GetConf()
	bansToMax  = config.Whitelist.BannedUsersToMaxAccounts
	footerIcon = config.Discord.FooterIcon
)

func PlayerJoin(serverConf conf.Server, PlayerName, footerIconURL, username *string, s *session.Session) discordgo.MessageEmbed {
	var (
		playerID, whitelisted = whitelist.GetOwner(*PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		roles, err            = s.GetRoles(playerID) // TODO: Only when whitelisted is true
		maxAccounts           = whitelist.GetMaxAccounts(roles)
		Title                 = fmt.Sprintf("%v joined the game", PlayerName)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
	if err != nil {
		log.Printf("Failed to get roles: %v", err)
	}
	if serverConf.SRV.Footer {
		if whitelisted {
			if !bansToMax {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
			} else {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
			}
		} else {
			FooterText = fmt.Sprintf("%v is not whitelisted", PlayerName)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: *footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}

	}
	var (
		color = 0x00FF00
		Embed discordgo.MessageEmbed
	)
	if serverConf.SRV.Footer {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		} else {
			Embed = discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    *PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		}
	} else {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		} else {
			Embed = discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    *PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		}
	}
	return Embed
}
func PlayerQuit(serverConf conf.Server, PlayerName, footerIconURL, username *string, s *session.Session) discordgo.MessageEmbed {
	var (
		playerID, whitelisted = whitelist.GetOwner(*PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		roles, err            = s.GetRoles(playerID) // TODO: Only when whitelisted is true
		maxAccounts           = whitelist.GetMaxAccounts(roles)
		Title                 = fmt.Sprintf("%v left the game", PlayerName)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
	if err != nil {
		log.Printf("Failed to get roles: %v", err)
	}
	if serverConf.SRV.Footer {
		if whitelisted {
			if !bansToMax {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
			} else {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
			}
		} else {
			FooterText = fmt.Sprintf("%v is not whitelisted", PlayerName)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: *footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	}
	var (
		Embed discordgo.MessageEmbed
		color = 0xFF0000
	)
	if serverConf.SRV.Footer {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}

		} else {
			Embed = discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    *PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		}
	} else {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		} else {
			Embed = discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    *PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		}
	}
	return Embed
}

func PlayerAdvancement(e *types.WebsocketEvent, serverConf *conf.Server, footerIconURL, username *string, s *session.Session) discordgo.MessageEmbed {
	var (
		PlayerName            = e.Data.Player
		advancement           = advancements.Decode(e.Data.Advancement)
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		roles, err            = s.GetRoles(playerID) // TODO: Only when whitelisted is true
		maxAccounts           = whitelist.GetMaxAccounts(roles)
		Title                 = fmt.Sprintf("%v made the Advancement %v", PlayerName, advancement)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
	if err != nil {
		log.Printf("Failed to get roles: %v", err)
	}
	if serverConf.SRV.Footer {
		if whitelisted {
			if !bansToMax {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
			} else {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
			}
		} else {
			FooterText = fmt.Sprintf("%v is not whitelisted", PlayerName)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: *footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	}
	var (
		Embed discordgo.MessageEmbed
		color = 0xFFFF00
	)
	if serverConf.SRV.Footer {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		} else {
			Embed = discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		}
	} else {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		} else {
			Embed = discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		}
	}
	return Embed
}

func PlayerDeath(e *types.WebsocketEvent, serverConf *conf.Server, footerIconURL, username *string, s *session.Session) *discordgo.MessageEmbed {
	var (
		PlayerName            = e.Data.Player
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		roles, err            = s.GetRoles(playerID) // TODO: Only when whitelisted is true
		maxAccounts           = whitelist.GetMaxAccounts(roles)
		Title                 = fmt.Sprintf("%v %v", PlayerName, e.Data.DeathMessage)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
	if err != nil {
		log.Printf("Failed to get roles: %v", err)
	}
	if serverConf.SRV.Footer {
		if whitelisted {
			if !bansToMax {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
			} else {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
			}
		} else {
			FooterText = fmt.Sprintf("%v is not whitelisted", PlayerName)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: *footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	}
	var (
		Embed *discordgo.MessageEmbed
		color = 0x000000
	)
	if serverConf.SRV.Footer {
		if serverConf.SRV.OneLine {
			Embed = &discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		} else {
			Embed = &discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		}
	} else {
		if serverConf.SRV.OneLine {
			Embed = &discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Title,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		} else {
			Embed = &discordgo.MessageEmbed{
				Title: Title,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		}
	}
	return Embed
}

func Chat(eventJson *types.WebsocketEvent, serverConf *conf.Server, footerIconURL, username *string, s *session.Session) *discordgo.MessageEmbed {
	var (
		PlayerName            = eventJson.Data.Player
		Message               = eventJson.Data.Message
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		roles, err            = s.GetRoles(playerID) // TODO: Only when whitelisted is true
		maxAccounts           = whitelist.GetMaxAccounts(roles)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
	if err != nil {
		log.Printf("Failed to get roles: %v", err)
	}
	if serverConf.SRV.Footer {
		if whitelisted {
			if !bansToMax {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
			} else {
				FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(bannedPlayers), maxAccounts)
			}
		} else {
			FooterText = fmt.Sprintf("%v is not whitelisted", PlayerName)
		}
		if footerIcon {
			Footer = &discordgo.MessageEmbedFooter{
				Text:    FooterText,
				IconURL: *footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	}
	var (
		Embed *discordgo.MessageEmbed
		color = 0x00AAFF
	)
	if serverConf.SRV.Footer {
		if serverConf.SRV.OneLine {
			Embed = &discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Message,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		} else {
			Embed = &discordgo.MessageEmbed{
				Title: Message,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		}
	} else {
		if serverConf.SRV.OneLine {
			Embed = &discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Message,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		} else {
			Embed = &discordgo.MessageEmbed{
				Title: Message,
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    PlayerName,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		}
	}
	return Embed
}
