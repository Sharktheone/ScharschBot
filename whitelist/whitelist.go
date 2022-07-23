package whitelist

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	config     = conf.GetConf()
	collection = config.Whitelist.Mongodb.MongodbCollectionName
)

func Add(username string, userID string) (alreadyListed bool, existing bool) {
	var found bool
	existingAcc := existingAccount(username)
	if existingAcc {
		_, found = mongodb.Read(collection, bson.M{
			"mcAccount": username,
		})
		if !found {

			mongodb.Write(collection, bson.D{
				{"dcUserID", userID},
				{"mcAccount", username},
			})

			log.Println(userID + "is adding " + username + " to whitelist")
		}

	}
	return found, existingAcc
}

func Remove(username string, userID string, roles []string) (allowed bool, onWhitelist bool) {
	var removeAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			removeAllowed = true
		}
	}
	entry, found := mongodb.Read(collection, bson.M{
		"dcUserID":  userID,
		"mcAccount": username,
	})
	if !removeAllowed {
		if entry[0]["dcUserID"] == userID && found {
			removeAllowed = true
		}
	}
	if removeAllowed && found {
		mongodb.Remove(collection, bson.M{
			"mcAccount": username,
		})
		log.Printf("%v is removing %v from whitelist", userID, username)
	}
	return removeAllowed, found
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

func existingAccount(username string) (existing bool) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%v", username)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to make check account existebility: %v\n", err)
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Failed reading Body white account check: %v\n", err)
	}
	if len(string(body)) > 0 {
		return true
	} else {
		return false
	}

}
