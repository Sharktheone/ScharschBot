package discordMember

import (
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config = conf.GetConf()
)

func GetUserProfile(userID string, s *discordgo.Session) (User *discordgo.Member, success bool) {
	user, err := s.GuildMember(config.Discord.ServerID, userID)
	if err != nil {
		log.Printf("Failed to get user profile: %v", err)
		return nil, false
	}
	return user, true
}

func GetRoles(userID string, s *discordgo.Session) []string {

	user, err := s.GuildMember(config.Discord.ServerID, userID)
	if err != nil {
		log.Printf("Error getting user %v: %v", userID, err)
	} else {
		return user.Roles
	}

	return nil
}

func SendDM(userID string, s *discordgo.Session, messageComplexDM *discordgo.MessageSend, messageComplexDMFailed *discordgo.MessageSend) (success bool) {
	var (
		successDM = false
	)
	channel, err := s.UserChannelCreate(userID)
	if err != nil {
		log.Printf("Failed to create DM with reporter: %v", err)

	}
	_, err = s.ChannelMessageSendComplex(channel.ID, messageComplexDM)
	if err != nil {
		log.Printf("Failed to send DM: %v, sending Message in normal Channels", err)
		for _, channelID := range config.Whitelist.Report.ChannelID {
			_, err = s.ChannelMessageSendComplex(channelID, messageComplexDMFailed)
			if err != nil {
				log.Printf("Failed to send message in dm alternative channel on server: %v", err)
			}
		}
		successDM = false
	} else {
		successDM = true
	}
	return successDM
}
