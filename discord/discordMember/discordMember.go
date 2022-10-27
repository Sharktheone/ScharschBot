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
		log.Printf("Error getting user: %v", err)
	}

	return user.Roles
}
