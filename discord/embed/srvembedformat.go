package embed

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/discordMember"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"github.com/Sharktheone/Scharsch-bot-discord/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PlayerJoin(PlayerName string, Players []string, banedPlayers []string, footer bool, oneLine bool, footerIcon bool, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, _   = whitelist.GetOwner(PlayerName)
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("%v joined the game", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)
	if footer {
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
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
	if footer {
		if oneLine {
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
		if oneLine {
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
func PlayerQuit(PlayerName string, Players []string, banedPlayers []string, footer bool, oneLine bool, footerIcon bool, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, _   = whitelist.GetOwner(PlayerName)
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("%v left the game", PlayerName)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)
	if footer {
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
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
	if footer {
		if oneLine {
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
		if oneLine {
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

func PlayerAdvancement(PlayerName string, Players []string, banedPlayers []string, Advancement string, footer bool, oneLine bool, footerIcon bool, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, _   = whitelist.GetOwner(PlayerName)
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("%v made the Advancement %v", PlayerName, Advancement)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)
	if footer {
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
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
	if footer {
		if oneLine {
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
		if oneLine {
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

func PlayerDeath(PlayerName string, Players []string, banedPlayers []string, DeathMessage string, footer bool, oneLine bool, footerIcon bool, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, _   = whitelist.GetOwner(PlayerName)
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		Title         = fmt.Sprintf("%v %v", PlayerName, DeathMessage)
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)
	if footer {
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
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
	if footer {
		if oneLine {
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
		if oneLine {
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
func Chat(PlayerName string, Players []string, banedPlayers []string, Message string, footer bool, oneLine bool, footerIcon bool, footerIconURL string, username string, s *discordgo.Session) discordgo.MessageEmbed {
	var (
		playerID, _   = whitelist.GetOwner(PlayerName)
		maxAccounts   = whitelist.GetMaxAccounts(discordMember.GetRoles(playerID, s))
		AuthorIconUrl = fmt.Sprintf("https://mc-heads.net/avatar/%v.png", PlayerName)
		AuthorUrl     = fmt.Sprintf("https://namemc.com/profile/%v", PlayerName)
		FooterText    string
		Footer        *discordgo.MessageEmbedFooter
	)
	if footer {
		if !bansToMax {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts (max %v)", username, len(Players), maxAccounts)
		} else {
			FooterText = fmt.Sprintf("%v • The owner has whitelisted %v accounts and %v are banned (max %v)", username, len(Players), len(banedPlayers), maxAccounts)
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
	if footer {
		if oneLine {
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
		if oneLine {
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
			URL:     ErrorURL,
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
