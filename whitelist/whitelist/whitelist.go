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
	config              = conf.GetConf()
	whitelistCollection = config.Whitelist.Mongodb.MongodbWhitelistCollectionName
	banCollection       = config.Whitelist.Mongodb.MongodbBanCollectionName
)

func Add(username string, userID string, roles []string) (alreadyListed bool, existing bool, accountFree bool, allowed bool, accountsListed []string, mcBanned bool, dcBanned bool) {
	var addAllowed = false
	mcBan, dcBan := checkBanned(username, userID)
	if !mcBan && !dcBan {
		for _, role := range roles {
			if role == config.Discord.WhitelistServerRoleID {
				addAllowed = true
			}
		}
	}
	var hasFreeAccount = false
	result, _ := mongodb.Read(whitelistCollection, bson.M{"dcUserID": userID})
	if config.Whitelist.MaxAccounts <= len(result) {
		hasFreeAccount = false
	} else {
		hasFreeAccount = true
	}
	var found bool
	existingAcc := existingAccount(username)
	if existingAcc && hasFreeAccount && addAllowed {
		_, found = mongodb.Read(whitelistCollection, bson.M{
			"mcAccount": username,
		})
		if !found {

			mongodb.Write(whitelistCollection, bson.D{
				{"dcUserID", userID},
				{"mcAccount", username},
			})

			log.Println(userID + " is adding " + username + " to whitelist")
		}

	}
	return found, existingAcc, hasFreeAccount, addAllowed, listedAccountsOf(userID), mcBan, dcBan
}

func Remove(username string, userID string, roles []string) (allowed bool, onWhitelist bool, accountsListed []string) {
	var removeAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			removeAllowed = true
		}
	}
	entry, found := mongodb.Read(whitelistCollection, bson.M{
		"dcUserID":  userID,
		"mcAccount": username,
	})
	if !removeAllowed {
		if entry[0]["dcUserID"] == userID && found {
			removeAllowed = true
		}
	}
	if removeAllowed && found && allowed {
		mongodb.Remove(whitelistCollection, bson.M{
			"mcAccount": username,
		})
		log.Printf("%v is removing %v from whitelist", userID, username)
	}
	return removeAllowed, found, listedAccountsOf(userID)
}

func RemoveAll(userID string, roles []string) (allowed bool, onWhitelist bool) {
	var removeAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			removeAllowed = true
		}
	}
	entries, found := mongodb.Read(whitelistCollection, bson.M{
		"dcUserID":  bson.M{"$exists": true},
		"mcAccount": bson.M{"$exists": true},
	})

	if removeAllowed && found {
		log.Printf("%v is removing all accounts from whitelist", userID)
		for _, entry := range entries {
			mongodb.Remove(whitelistCollection, bson.M{
				"mcAccount": entry["mcAccount"],
			})

		}

	}

	return removeAllowed, found
}
func RemoveAllAllowed(roles []string) (allowed bool) {
	var removeAllowed = false
	for _, role := range roles {
		if role == config.Discord.WhitelistRemoveRoleID {
			removeAllowed = true
		}
	}
	return removeAllowed
}

func Whois(username string, userID string, roles []string) (dcUserID string, allowed bool, found bool, accountsListed []string) {
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
		result, dataFound = mongodb.Read(whitelistCollection, bson.M{
			"mcAccount": username,
		})
		if dataFound {
			dcUser = fmt.Sprintf("%v", result[0]["dcUserID"])
		}
	}
	return dcUser, whoisAllowed, dataFound, listedAccountsOf(userID)
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

		results, dataFound = mongodb.Read(whitelistCollection, bson.M{
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
func listedAccountsOf(userID string) (Accounts []string) {
	results, dataFound := mongodb.Read(whitelistCollection, bson.M{
		"dcUserID": userID,
	})
	listedAccounts := make([]string, len(results), 10)
	if dataFound {

		for i, result := range results {
			listedAccounts[i] = fmt.Sprintf("%v", result["mcAccount"])
		}
	}
	return listedAccounts
}

func BanUserID(userID string, roles []string, banID string) (allowed bool, accountsListed []string) {
	banAllowed := false
	listedAccounts := listedAccountsOf(userID)
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			banAllowed = true
		}
	}
	if banAllowed {
		log.Printf("%v is banning %v", userID, banID)
		mongodb.Write(banCollection, bson.D{
			{"dcUserID", banID},
		})
	}
	return banAllowed, listedAccounts
}

func BanAccount(userID string, roles []string, account string) (allowed bool, accountsListed []string, ownerID string) {
	banAllowed := false
	listedAccounts := listedAccountsOf(userID)
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			banAllowed = true
		}
	}
	if banAllowed {
		log.Printf("%v is banning %v", userID, account)
		mongodb.Write(banCollection, bson.D{
			{"mcAccount", account},
		})
	}

	var result []bson.M
	var (
		dcUser    string
		dataFound bool
	)
	result, dataFound = mongodb.Read(whitelistCollection, bson.M{
		"mcAccount": account,
	})
	if dataFound {
		dcUser = fmt.Sprintf("%v", result[0]["dcUserID"])
	}
	return banAllowed, listedAccounts, dcUser
}
func UnBanUserID(userID string, roles []string, banID string) (allowed bool) {
	unBanAllowed := false
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			unBanAllowed = true
		}
	}
	if unBanAllowed {
		log.Printf("%v is unbanning %v", userID, banID)
		mongodb.Remove(banCollection, bson.M{
			"dcUserID": banID,
		})
	}
	return unBanAllowed
}

func UnBanAccount(userID string, roles []string, account string) (allowed bool, accountsListed []string) {
	unBanAllowed := false
	listedAccounts := listedAccountsOf(userID)
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			unBanAllowed = true
		}
	}
	if unBanAllowed {
		log.Printf("%v is unbanning %v", userID, account)
		mongodb.Remove(banCollection, bson.M{
			"mcAccount": account,
		})
	}
	return unBanAllowed, listedAccounts
}

func checkBanned(mcName string, userID string) (mcBanned bool, dcBanned bool) {
	var (
		dataFound bool
		mc        = false
		dc        = false
	)
	_, dataFound = mongodb.Read(banCollection, bson.M{
		"mcAccount": mcName,
	})
	if dataFound {
		mc = true
	}

	_, dataFound = mongodb.Read(banCollection, bson.M{
		"dcUserID": userID,
	})
	if dataFound {
		dc = true
	}
	return mc, dc
}
