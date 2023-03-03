package reports

import (
	"Scharsch-Bot/conf"
	"Scharsch-Bot/database/mongodb"
	"Scharsch-Bot/discord/discordMember"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var (
	config           = conf.GetConf()
	reportCollection = config.Whitelist.Mongodb.MongodbReportCollectionName
)

type ReportData struct {
	ReporterID     string `bson:"reporterID"`
	ReportedPlayer string `bson:"reportedPlayer"`
	Reason         string `bson:"reason"`
}

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
				var roleMessage string
				for _, role := range config.Whitelist.Report.PingRoleID {
					ping := fmt.Sprintf("<@&%s> ", role)
					roleMessage += ping
				}
				for _, channel := range config.Whitelist.Report.ChannelID {
					_, err := s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
						Content: roleMessage,
						Embed:   &messageEmbed,
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

func Reject(name string, i *discordgo.InteractionCreate, s *discordgo.Session, notifyreporter bool, messageEmbed *discordgo.MessageEmbed, messageEmbedDMFailed *discordgo.MessageEmbed) (rejectAllowed bool, enabled bool) {
	var (
		allowed  = false
		notifyDM = config.Whitelist.Report.PlayerNotifyDM
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
	}

	if allowed {
		report, reportFound := GetReport(name)
		if reportFound {
			if notifyDM {
				if notifyreporter {
					channel, err := s.UserChannelCreate(report.ReporterID)
					if err != nil {
						log.Printf("Failed to create DM with reporter: %v", err)

					}
					_, err = s.ChannelMessageSendEmbed(channel.ID, messageEmbed)
					if err != nil {
						log.Printf("Failed to send DM for reporter: %v, sending Message in normal Channels", err)
						for _, channelID := range config.Whitelist.Report.ChannelID {
							_, err = s.ChannelMessageSendEmbed(channelID, messageEmbedDMFailed)
							if err != nil {
								log.Printf("Failed to send Report message in normal Channel: %v", err)
							}
						}
					}

				}
			} else {
				if notifyreporter {
					for _, channelID := range config.Whitelist.Report.ChannelID {
						_, err := s.ChannelMessageSendEmbed(channelID, messageEmbed)
						if err != nil {
							log.Printf("Failed to send Report message : %v", err)
						}
					}

				}
			}
			DeleteReport(name)
		}
	}

	return allowed, config.Whitelist.Report.Enabled
}
func Accept(name string, i *discordgo.InteractionCreate, s *discordgo.Session, notifyreporter bool, messageEmbed *discordgo.MessageEmbed, messageEmbedDMFailed *discordgo.MessageEmbed) (acceptAllowed bool, enabled bool) {
	var (
		allowed  = false
		notifyDM = config.Whitelist.Report.PlayerNotifyDM
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
	}
	if allowed {
		report, reportFound := GetReport(name)
		if reportFound {
			if notifyDM {
				if notifyreporter {
					discordMember.SendDM(report.ReporterID, s, &discordgo.MessageSend{
						Embed: messageEmbed,
					},
						&discordgo.MessageSend{
							Content: fmt.Sprintf("<@%v>", report.ReporterID),
							Embed:   messageEmbedDMFailed,
						})

				}
			} else {
				if notifyreporter {

					for _, channelID := range config.Whitelist.Report.ChannelID {
						_, err := s.ChannelMessageSendEmbed(channelID, messageEmbed)
						if err != nil {
							log.Printf("Failed to send Report message : %v", err)
						}
					}

				}
			}
			DeleteReport(name)
		}
	}
	return allowed, config.Whitelist.Report.Enabled
}

func DeleteReport(name string) {
	mongodb.Remove(reportCollection, bson.M{
		"reportedPlayer": name,
	})
}
func GetReport(name string) (report ReportData, reportFound bool) {
	data, dataFound := mongodb.Read(reportCollection, bson.M{
		"reportedPlayer": name,
	})
	var reportData ReportData
	// TODO Converte with bson.NewDecoder
	if dataFound {
		reportData.ReportedPlayer = data[0]["reportedPlayer"].(string)
		reportData.ReporterID = data[0]["reporterID"].(string)
		reportData.Reason = data[0]["reason"].(string)
	}

	return reportData, dataFound
}
