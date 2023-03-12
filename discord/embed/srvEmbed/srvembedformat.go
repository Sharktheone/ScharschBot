package srvEmbed

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/discord/discordMember"
	"Scharsch-Bot/minecraft/advancements"
	"Scharsch-Bot/pterodactyl"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	config     = conf.GetConf()
	bansToMax  = config.Whitelist.BannedUsersToMaxAccounts
	footerIcon = config.Discord.FooterIcon
)

func PlayerJoin(serverConf conf.Server, PlayerName string, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		maxAccounts           = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title                 = fmt.Sprintf("%v joined the game", PlayerName)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
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
				IconURL: footerIconURL,
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
func PlayerQuit(serverConf conf.Server, PlayerName string, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		maxAccounts           = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title                 = fmt.Sprintf("%v left the game", PlayerName)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
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
				IconURL: footerIconURL,
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

func PlayerAdvancement(eventJson types.EventJson, serverConf conf.Server, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		PlayerName            = eventJson.Name
		advancement           = advancements.Decode(eventJson.Value)
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		maxAccounts           = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title                 = fmt.Sprintf("%v made the Advancement %v", PlayerName, advancement)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
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
				IconURL: footerIconURL,
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

func PlayerDeath(eventJson types.EventJson, serverConf conf.Server, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		PlayerName            = eventJson.Name
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		maxAccounts           = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title                 = fmt.Sprintf("%v %v", PlayerName, eventJson.Value)
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
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
				IconURL: footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	}
	var (
		Embed discordgo.MessageEmbed
		color = 0x000000
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

func Chat(eventJson types.EventJson, serverConf conf.Server, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		PlayerName            = eventJson.Name
		Message               = eventJson.Value
		playerID, whitelisted = whitelist.GetOwner(PlayerName)
		Players               = whitelist.ListedAccountsOf(playerID, true)
		bannedPlayers         = whitelist.CheckBans(playerID)
		maxAccounts           = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		AuthorIconUrl         = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl             = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText            string
		Footer                *discordgo.MessageEmbedFooter
	)
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
				IconURL: footerIconURL,
			}
		} else {
			Footer = &discordgo.MessageEmbedFooter{
				Text: FooterText,
			}
		}
	}
	var (
		Embed discordgo.MessageEmbed
		color = 0x00AAFF
	)
	if serverConf.SRV.Footer {
		if serverConf.SRV.OneLine {
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Message,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
				Footer: Footer,
			}
		} else {
			Embed = discordgo.MessageEmbed{
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
			Embed = discordgo.MessageEmbed{
				Color: color,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    Message,
					IconURL: AuthorIconUrl,
					URL:     AuthorUrl,
				},
			}
		} else {
			Embed = discordgo.MessageEmbed{
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

func Power(action string) discordgo.MessageEmbed {
	var (
		c      = cases.Title(language.English)
		Title  = fmt.Sprintf("%v server(s)", c.String(action))
		color  int
		Fields []*discordgo.MessageEmbedField
	)
	for _, server := range pterodactyl.ServerStates {
		var StateMsg string
		if server.Status == "starting" {
			StateMsg = config.SRV.States.Starting
		} else if server.Status == "stopping" {
			StateMsg = config.SRV.States.Stopping
		} else if server.Status == "running" {
			StateMsg = config.SRV.States.Online
		} else if server.Status == "offline" {
			StateMsg = config.SRV.States.Offline
		}
		Field := &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("%v:", server.Name),
			Value: StateMsg,
		}
		Fields = append(Fields, Field)
	}
	if action == "start" {
		color = 0x00FF00
	} else if action == "stop" {
		color = 0xFF0000
	} else if action == "restart" {
		color = 0xFFFF00
	} else if action == "status" {
		color = 0x00AAFF
	} else {
		color = 0x000000
	}
	Embed := discordgo.MessageEmbed{
		Title:  Title,
		Color:  color,
		Fields: Fields,
	}
	return Embed
}
func PowerNotAllowed(avatarURL string, name string, action string, serverName string) discordgo.MessageEmbed {

	var (
		Title string
	)
	if serverName != "" {
		Title = fmt.Sprintf("You have no permission to %v %v", action, serverName)
	} else {
		Title = fmt.Sprintf("You have no permission to %v servers", action)
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: 0xFF0000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatarURL,
			URL:     config.Discord.EmbedErrorAuthorURL,
		},
	}
	return Embed
}
func PowerAction(action string, serverName string, avatarURL string, name string) discordgo.MessageEmbed {
	var (
		Title string
		color int
	)
	Title = fmt.Sprintf("Server %v is getting %ved", serverName, action)
	switch action {
	case "start":
		color = 0x00FF00
	case "stop":
		color = 0xFF0000
	case "restart":
		color = 0xFFFF00
	}

	Embed := discordgo.MessageEmbed{
		Title: Title,
		Color: color,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: avatarURL,
		},
	}
	return Embed
}
