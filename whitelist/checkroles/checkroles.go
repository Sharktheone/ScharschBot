package checkroles

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/bot"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"github.com/Sharktheone/Scharsch-bot-discord/srv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var (
	config                = conf.GetConf()
	whitelistCollection   = config.Whitelist.Mongodb.MongodbWhitelistCollectionName
	reWhitelistCollection = config.Whitelist.Mongodb.MongodbReWhitelistCollectionName
	reWhitelist           = config.Whitelist.Roles.ReWhitelistWith
	removeWithout         = config.Whitelist.Roles.RemoveUserWithout
	kickUnWhitelisted     = config.Whitelist.KickUnWhitelisted
)

func CheckRoles() {
	if kickUnWhitelisted {
		for _, player := range srv.OnlinePlayers {
			_, found := mongodb.Read(whitelistCollection, bson.M{
				"dcUserID":  bson.M{"$exists": true},
				"mcAccount": player,
			})
			if !found {
				command := fmt.Sprintf(config.Whitelist.KickCommand, player)
				for _, listedServer := range config.Whitelist.Servers {
					for _, server := range config.Pterodactyl.Servers {
						if server.ServerName == listedServer {
							pterodactyl.SendCommand(command, server.ServerID)
						}
					}
				}
			}
		}
	}
	if removeWithout {
		entries, found := mongodb.Read(whitelistCollection, bson.M{
			"dcUserID":  bson.M{"$exists": true},
			"mcAccount": bson.M{"$exists": true},
		})

		if found {
			session := bot.Session
			var removedIDs []string
			for _, entry := range entries {
				userID := fmt.Sprintf("%v", entry["dcUserID"])

				checkID := true
				for _, removeID := range removedIDs {
					if removeID == userID {
						checkID = false
					}
				}
				if checkID {
					user, err := session.GuildMember(config.Discord.ServerID, userID)
					if err != nil {
						log.Println("Error getting user " + userID)
					}
					serverPerms := false
					for _, role := range user.Roles {
						if role == config.Whitelist.Roles.ServerRoleID {
							serverPerms = true
						}
					}
					if serverPerms == false {
						removedIDs = append(removedIDs, userID)

						mongodb.Remove(whitelistCollection, bson.M{
							"dcUserID": userID,
						})
						mongodb.Write(reWhitelistCollection, bson.D{
							{"dcUserID", userID},
							{"mcAccount", entry["mcAccount"]},
						})

					}
				}

			}
			if len(removedIDs) > 0 {
				log.Printf("Removing accounts of the id(s) %v from whitelist because they have not the server role", removedIDs)
			}
		}
	}
	if reWhitelist {
		entries, found := mongodb.Read(reWhitelistCollection, bson.M{
			"dcUserID":  bson.M{"$exists": true},
			"mcAccount": bson.M{"$exists": true},
		})

		if found {
			session := bot.Session
			var addedIDs []string
			for _, entry := range entries {
				userID := fmt.Sprintf("%v", entry["dcUserID"])

				checkID := true
				for _, addID := range addedIDs {
					if addID == userID {
						checkID = false
					}
				}
				if checkID {
					user, err := session.GuildMember(config.Discord.ServerID, userID)
					if err != nil {
						log.Println("Error getting user " + userID)
					}
					serverPerms := false
					for _, role := range user.Roles {
						if role == config.Whitelist.Roles.ServerRoleID {
							serverPerms = true
						}
					}
					if serverPerms == true {

						addedIDs = append(addedIDs, userID)

						mongodb.Remove(reWhitelistCollection, bson.M{
							"dcUserID":  userID,
							"mcAccount": entry["mcAccount"],
						})
						mongodb.Write(whitelistCollection, bson.D{
							{"dcUserID", userID},
							{"mcAccount", entry["mcAccount"]},
						})

					}
				}
			}
			if len(addedIDs) > 0 {
				log.Printf("Adding accounts of the id(s) %v to whitelist because they have the server role again", addedIDs)
			}

		}
	}
}
