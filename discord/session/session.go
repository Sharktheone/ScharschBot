package session

import (
	"Scharsch-Bot/conf"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	config = conf.GetConf()
)

type Session struct {
	*discordgo.Session
}

func (s *Session) GetUserProfile(userID string) (*discordgo.Member, error) {
	if user, err := s.GuildMember(config.Discord.ServerID, userID); err != nil {
		return &discordgo.Member{}, fmt.Errorf("failed to get user profile: %v", err)
	} else {
		return user, nil
	}
}

func (s *Session) GetRoles(userID string) ([]string, error) {
	if user, err := s.GuildMember(config.Discord.ServerID, userID); err != nil {
		return nil, err
	} else {
		return user.Roles, nil
	}
}

func (s *Session) SendDM(userID string, messageComplexDM *discordgo.MessageSend, messageComplexDMFailed *discordgo.MessageSend) error {
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
				return fmt.Errorf("failed to send message in dm alternative channel on server: %v", err)
			}
		}
	}
	return nil
}

func HasRole(member *discordgo.Member, roleIDs []string) bool {
	for _, role := range member.Roles {
		for _, neededRole := range roleIDs {
			if role == neededRole {
				return true
			}
		}
	}
	return false
}
