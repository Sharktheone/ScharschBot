package checkroles

import (
	"fmt"
	"github.com/Sharktheone/ScharschBot/conf"
	"github.com/Sharktheone/ScharschBot/database/mongodb"
	"github.com/Sharktheone/ScharschBot/discord/bot"
	"github.com/Sharktheone/ScharschBot/pterodactyl"
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
		for _, server := range pterodactyl.Servers {
			for _, player := range server.OnlinePlayers.Players {
				_, found := mongodb.Read(whitelistCollection, bson.M{
					"dcUserID":  bson.M{"$exists": true},
					"mcAccount": player,
				})
				if !found {
					command := fmt.Sprintf(config.Whitelist.KickCommand, player)
					if err := pterodactyl.SendCommand(command, server.Config.ServerID); err != nil {
						log.Printf("Failed to kick %v from server %v: %v", player, server.Config.ServerID, err)
					} else {
						server.OnlinePlayers.Mu.Lock()
						for i, player := range server.OnlinePlayers.Players {
							if player == player {
								players := server.OnlinePlayers.Players
								if i == len(players)-1 {
									server.OnlinePlayers.Players = players[:i]
								} else {
									server.OnlinePlayers.Players = append(players[:i], players[i+1:]...)
								}
							}
						}
						server.OnlinePlayers.Mu.Unlock()
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
					user, _ := session.GuildMember(config.Discord.ServerID, userID)
					if user == nil {

						removedIDs = append(removedIDs, userID)
						mongodb.Remove(whitelistCollection, bson.M{
							"dcUserID": userID,
						})
						mongodb.Write(reWhitelistCollection, bson.D{
							{"dcUserID", userID},
							{"mcAccount", entry["mcAccount"]},
						})
					} else {
						serverPerms := false
						for _, role := range user.Roles {
							for _, neededRole := range config.Whitelist.Roles.ServerRoleID {
								if role == neededRole {
									serverPerms = true
									break
								}
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
					user, _ := session.GuildMember(config.Discord.ServerID, userID)
					if user != nil {
						serverPerms := false
						for _, role := range user.Roles {
							for _, neededRole := range config.Whitelist.Roles.ServerRoleID {
								if role == neededRole {
									serverPerms = true
									break
								}
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
			}
			if len(addedIDs) > 0 {
				log.Printf("Adding accounts of the id(s) %v to whitelist because they have the server role again", addedIDs)
			}

		}
	}
}
