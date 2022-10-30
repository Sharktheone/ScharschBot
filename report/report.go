package report

import (
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var (
	config           = conf.GetConf()
	reportCollection = config.Whitelist.Mongodb.MongodbReportCollectionName
)

func GetReports() (reports []bson.M, anyReports bool) {
	report, dataFound := mongodb.Read(reportCollection, bson.M{
		"reportedPlayer": bson.M{"$exists": true},
	})
	return report, dataFound
}

func Report(name string, reason string, i *discordgo.InteractionCreate, s *discordgo.Session, messageEmbed discordgo.MessageEmbed) (reporAllowed bool, alreadyReported bool, enabled bool) {
	var (
		allowed   = false
		dataFound bool
	)
	if config.Whitelist.Report.Enabled {
		for _, role := range i.Member.Roles {
			for _, requiredRole := range config.Whitelist.Report.Roles {
				if role == requiredRole {
					allowed = true
					break
				}
			}
		}
		if allowed {
			_, dataFound = mongodb.Read(reportCollection, bson.M{
				"reportedPlayer": name,
			})
			if !dataFound {
				mongodb.Write(reportCollection, bson.D{
					{"reporterID", i.Member.User.ID},
					{"reportedPlayer", name},
					{"reason", reason},
				})
				for _, channel := range config.Whitelist.Report.ChannelID {
					_, err := s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
						Embed: &messageEmbed,
						AllowedMentions: &discordgo.MessageAllowedMentions{
							Roles: config.Whitelist.Report.Roles,
						},
					})
					if err != nil {
						log.Printf("Failed to send report embed: %v", err)
					}
				}
			}
		}
	}
	return allowed, dataFound, config.Whitelist.Report.Enabled
}

func Reject(name string, i *discordgo.InteractionCreate, s *discordgo.Session) (allowed bool) {

	return
}
func Accept(name string, i *discordgo.InteractionCreate, s *discordgo.Session) (allowed bool) {

	return
}
