package whitelist

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var (
	config     = conf.GetConf()
	collection = config.Whitelist.Mongodb.MongodbCollectionName
)

func Add(username string, userID string) (alreadyListed bool) {
	_, found := mongodb.Read(collection, bson.M{
		"mcAccount": username,
	})
	if !found {

		mongodb.Write(collection, bson.D{
			{"dcUserID", userID},
			{"mcAccount", username},
		})

		log.Println(userID + "is adding " + username + " to whitelist")
	}
	return found
}

func Remove(username string, userID string, roles []string) (allowed bool) {
	var removeAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			removeAllowed = true
		}
	}

	if !removeAllowed {
		entry, found := mongodb.Read(collection, bson.M{
			"dcUserID":  userID,
			"mcAccount": username,
		})
		if entry[0]["dcUserID"] == userID && found {
			removeAllowed = true
		}
	}
	if removeAllowed {
		mongodb.Remove(collection, bson.M{
			"mcAccount": username,
		})
		log.Printf("%v is removing %v from whitelist", userID, username)
	}
	return removeAllowed
}

func Whois(username string, userID string, roles []string) (dcUserID string, allowed bool, found bool) {
	var whoisAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			whoisAllowed = true
		}
	}
	var (
		dcUser    string
		dataFound bool
	)
	if whoisAllowed {
		log.Printf("%v is looking who whitelisted %v ", userID, username)
		var result []bson.M
		result, dataFound = mongodb.Read(collection, bson.M{
			"mcAccount": username,
		})
		if dataFound {
			dcUser = fmt.Sprintf("%v", result[0]["dcUserID"])
		}
	}
	return dcUser, whoisAllowed, dataFound
}
