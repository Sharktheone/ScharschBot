package handlers

import (
	"Scharsch-Bot/database/mongodb"
	"Scharsch-Bot/discord/embed/wEmbed"
	"Scharsch-Bot/whitelist/whitelist"
	"github.com/bwmarrin/discordgo"
	"log"
)

func RemoveYes(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var messageEmbed discordgo.MessageEmbed
	if mongodb.Ready {
		allowed, onWhitelist := whitelist.RemoveAll(i.Member.User.ID, i.Member.Roles)
		if allowed {
			if onWhitelist {
				messageEmbed = wEmbed.WhitelistRemoveAll(i)
			} else {
				messageEmbed = wEmbed.WhitelistRemoveAllNoWhitelistEntries(i)
			}
		} else {
			messageEmbed = wEmbed.WhitelistRemoveAllNotAllowed(i)
		}
	} else {
		messageEmbed = wEmbed.DatabaseNotReady
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				&messageEmbed,
			},
		},
	})
	if err != nil {
		log.Printf("Failed execute component remove_yes: %v", err)
	}
}
