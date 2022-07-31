package checkroles

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/discord/bot"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var (
	collection = config.Whitelist.Mongodb.MongodbWhitelistCollectionName
	config     = conf.GetConf()
)

func CheckRoles() {
	entries, found := mongodb.Read(collection, bson.M{
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

					mongodb.Remove(collection, bson.M{
						"dcUserID": userID,
					})

				}
			}

		}
		if len(removedIDs) > 0 {
			log.Printf("Removing accounts of the id(s) %v from whitelist because they have not the server role", removedIDs)
		}
	}
}
