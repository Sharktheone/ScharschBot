package whitelist

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckRoles() {
	entries, found := mongodb.Read(collection, bson.M{
		"dcUserID":  "",
		"mcAccount": "",
	})

	if found {
		for _, entry := range entries {
			userID := fmt.Sprintf("%v", entry["dcUserID"])
			serverPerms := false
			roles := discordgo.Member{GuildID: userID}.Roles
			for _, role := range roles {
				if role == config.Whitelist.Roles.ServerRoleID {
					serverPerms = true
				}
			}
			if serverPerms == false {
				mongodb.Remove(collection, bson.M{
					"dcUserID": userID,
				})
			}

		}
	}
}
