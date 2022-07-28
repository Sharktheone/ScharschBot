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

func Add(username string, userID string, roles []string) (alreadyListed bool, existing bool, accountFree bool, allowed bool) {
	var addAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			addAllowed = true
		}
	}
	var hasFreeAccount = false
	result, _ := mongodb.Read(collection, bson.M{"dcUserID": userID})
	if config.Whitelist.MaxAccounts <= len(result) {
		hasFreeAccount = false
	} else {
		hasFreeAccount = true
	}
	var found bool
	existingAcc := existingAccount(username)
	if existingAcc && hasFreeAccount {
		_, found = mongodb.Read(collection, bson.M{
			"mcAccount": username,
		})
		if !found {

			mongodb.Write(collection, bson.D{
				{"dcUserID", userID},
				{"mcAccount", username},
			})

			log.Println(userID + " is adding " + username + " to whitelist")
		}

	}
	return found, existingAcc, hasFreeAccount, addAllowed
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

func RemoveAll(userID string, roles []string) (allowed bool, onWhitelist bool) {
	var removeAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			removeAllowed = true
		}
	}
	entries, found := mongodb.Read(collection, bson.M{
		"dcUserID":  bson.M{"$exists": true},
		"mcAccount": bson.M{"$exists": true},
	})

	if removeAllowed && found {
		log.Printf("%v is removing all accounts from whitelist", userID)
		for _, entry := range entries {
			mongodb.Remove(collection, bson.M{
				"mcAccount": entry["mcAccount"],
			})

		}

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
func HasListed(lookupID string, userID string, roles []string) (accounts []string, allowed bool, found bool) {
	var listedAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			listedAllowed = true
		}
	}
	var (
		results   []bson.M
		dataFound bool
	)
	var listedAcc []string
	if listedAllowed {
		log.Printf("%v is looking on whitelisted accounts of %v ", userID, lookupID)

		results, dataFound = mongodb.Read(collection, bson.M{
			"dcUserID": lookupID,
		})
		listedAccounts := make([]string, len(results), 10)
		for i, result := range results {
			listedAccounts[i] = fmt.Sprintf("%v", result["mcAccount"])

		}
		listedAcc = listedAccounts
	}
	return listedAcc, listedAllowed, dataFound
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
