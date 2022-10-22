package whitelist

import (
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"github.com/Sharktheone/Scharsch-bot-discord/pterodactyl"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	config              = conf.GetConf()
	whitelistCollection = config.Whitelist.Mongodb.MongodbWhitelistCollectionName
	banCollection       = config.Whitelist.Mongodb.MongodbBanCollectionName
	addCommand          = config.Pterodactyl.WhitelistAddCommand
	removeCommand       = config.Pterodactyl.WhitelistRemoveCommand
	pterodactylEnabled  = config.Pterodactyl.Enabled
)

func Add(username string, userID string, roles []string) (alreadyListed bool, existing bool, accountFree bool, allowed bool, mcBanned bool, dcBanned bool, banReason string) {
	var addAllowed = false
	mcBan, dcBan, reason := checkBanned(username, userID)
	if !mcBan && !dcBan {
		for _, role := range roles {
			if role == config.Whitelist.Roles.ServerRoleID {
				addAllowed = true
			}
		}
	}
	var hasFreeAccount = false
	result, _ := mongodb.Read(whitelistCollection, bson.M{"dcUserID": userID})
	if config.Whitelist.MaxAccounts <= (len(result) + len(CheckBans(userID))) {
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
			if pterodactylEnabled {
				command := fmt.Sprintf(addCommand, username)
				for _, listedServer := range config.Whitelist.Servers {
					for _, server := range config.Pterodactyl.Servers {
						if server.ServerName == listedServer {
							pterodactyl.SendCommand(command, server.ServerID)
						}
					}
				}
			}
			log.Println(userID + " is adding " + username + " to whitelist")
		}

	}
	return found, existingAcc, hasFreeAccount, addAllowed, mcBan, dcBan, reason
}

func Remove(username string, userID string, roles []string) (allowed bool, onWhitelist bool) {
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
	if removeAllowed && found {
		mongodb.Remove(whitelistCollection, bson.M{
			"mcAccount": username,
		})
		if pterodactylEnabled {
			command := fmt.Sprintf(removeCommand, username)
			for _, listedServer := range config.Whitelist.Servers {
				for _, server := range config.Pterodactyl.Servers {
					if server.ServerName == listedServer {
						pterodactyl.SendCommand(command, server.ServerID)
					}
				}
			}
		}
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
			if pterodactylEnabled {
				command := fmt.Sprintf(removeCommand, entry["mcAccount"])
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
		result, dataFound = mongodb.Read(whitelistCollection, bson.M{
			"mcAccount": username,
		})
		if dataFound {
			dcUser = fmt.Sprintf("%v", result[0]["dcUserID"])
		}
	}
	return dcUser, whoisAllowed, dataFound
}
func HasListed(lookupID string, userID string, roles []string) (accounts []string, allowed bool, found bool, bannedPlayers []string) {
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
	return listedAcc, listedAllowed, dataFound, CheckBans(userID)
}

func existingAccount(username string) (existing bool) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%v", username)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to make check account existebility: %v\n", err)
	}
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Printf("Failed reading Body white account check: %v\n", err)
	}
	return len(string(body)) > 0

}
func ListedAccountsOf(userID string) (Accounts []string) {
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

func BanUserID(userID string, roles []string, banID string, banAccounts bool, reason string) (allowed bool, accountsListed []string) {
	banAllowed := false
	listedAccounts := ListedAccountsOf(banID)
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			banAllowed = true
		}
	}
	if banAllowed {
		log.Printf("%v is banning %v", userID, banID)
		mongodb.Write(banCollection, bson.D{
			{"dcUserID", banID},
			{"reason", reason},
		})
		if banAccounts {
			for _, account := range listedAccounts {
				mongodb.Remove(whitelistCollection, bson.M{
					"mcAccount": account,
				})
				if pterodactylEnabled {
					command := fmt.Sprintf(removeCommand, account)
					for _, listedServer := range config.Whitelist.Servers {
						for _, server := range config.Pterodactyl.Servers {
							if server.ServerName == listedServer {
								pterodactyl.SendCommand(command, server.ServerID)
							}
						}
					}
				}
				mongodb.Write(banCollection, bson.D{
					{"mcAccount", account},
					{"dcUserID", banID},
				})

			}
		}
	}
	return banAllowed, listedAccounts
}

func BanAccount(userID string, roles []string, account string, reason string) (allowed bool, accountsListed []string, ownerID string) {
	var (
		banAllowed     = false
		listedAccounts []string
	)
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			banAllowed = true
		}
	}

	dcUser, dataFound := GetOwner(account)
	_, alreadyBanned := mongodb.Read(banCollection, bson.M{
		"mcAccount": account,
	})
	if dataFound {
		listedAccounts = ListedAccountsOf(dcUser)
	}

	if banAllowed && !alreadyBanned {
		log.Printf("%v is banning %v", userID, account)
		mongodb.Write(banCollection, bson.D{
			{"mcAccount", account},
			{"dcUserID", dcUser},
			{"reason", reason},
		})
		mongodb.Remove(whitelistCollection, bson.M{
			"mcAccount": account,
		})
		if pterodactylEnabled {
			command := fmt.Sprintf(removeCommand, account)
			for _, listedServer := range config.Whitelist.Servers {
				for _, server := range config.Pterodactyl.Servers {
					if server.ServerName == listedServer {
						pterodactyl.SendCommand(command, server.ServerID)
					}
				}
			}
		}
	}

	return banAllowed, listedAccounts, dcUser
}
func UnBanUserID(userID string, roles []string, banID string, unbanAccounts bool) (allowed bool) {
	unBanAllowed := false
	for _, role := range roles {
		if role == config.Discord.WhitelistBanRoleID {
			unBanAllowed = true
		}
	}
	if unBanAllowed {
		log.Printf("%v is unbanning %v", userID, banID)
		mongodb.Remove(banCollection, bson.M{
			"dcUserID":  banID,
			"mcAccount": bson.M{"$exists": false},
		})
		if unbanAccounts {
			var result []bson.M
			result, dataFound := mongodb.Read(banCollection, bson.M{
				"dcUserID": banID,
			})
			if dataFound {
				for _, entry := range result {
					mongodb.Remove(banCollection, bson.M{
						"mcAccount": entry["mcAccount"],
						"dcUserID":  banID,
					})
				}
			}
		}
	}
	return unBanAllowed
}

func UnBanAccount(userID string, roles []string, account string) (allowed bool, accountsListed []string) {
	unBanAllowed := false
	var (
		dcUser         string
		listedAccounts []string
	)
	dcUser, dataFound := GetOwner(account)
	if dataFound {
		listedAccounts = ListedAccountsOf(dcUser)
	}

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

func checkBanned(mcName string, userID string) (mcBanned bool, dcBanned bool, banReason string) {
	var (
		dataFound bool
		mcData    []bson.M
		dcData    []bson.M
		mc        = false
		dc        = false
		reason    string
	)
	mcData, dataFound = mongodb.Read(banCollection, bson.M{
		"mcAccount": mcName,
	})
	if dataFound {
		mc = true
	}

	dcData, dataFound = mongodb.Read(banCollection, bson.M{
		"dcUserID":  userID,
		"mcAccount": bson.M{"$exists": false},
	})
	if dataFound {
		dc = true
	}
	if mc {
		reason = fmt.Sprintf("%v", mcData[0]["reason"])
	}
	if dc {
		reason += fmt.Sprintf("%v", dcData[0]["reason"])
	}

	return mc, dc, reason
}

func CheckBans(userID string) (bannedPlayers []string) {
	var (
		dataFound bool
		results   []bson.M
	)
	results, dataFound = mongodb.Read(banCollection, bson.M{
		"dcUserID":  userID,
		"mcAccount": bson.M{"$exists": true},
	})

	var bannedAccounts = make([]string, len(results), 10)
	if dataFound {
		for i, result := range results {
			bannedAccounts[i] = fmt.Sprintf("%v", result["mcAccount"])

		}
	}
	return bannedAccounts
}

func RemoveMyAccounts(userID string) (hadListedAccounts bool, listedAccounts []string) {

	var (
		accounts          = ListedAccountsOf(userID)
		hasListedAccounts = false
	)
	if len(accounts) > 0 {
		hasListedAccounts = true
		log.Printf("%v is removing his own accounts from the whitelist", userID)
		for _, account := range accounts {
			_, found := mongodb.Read(whitelistCollection, bson.M{
				"dcUserID":  userID,
				"mcAccount": account,
			})
			if found {
				mongodb.Remove(whitelistCollection, bson.M{
					"mcAccount": account,
				})
				if pterodactylEnabled {
					command := fmt.Sprintf(removeCommand, account)
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
	}

	return hasListedAccounts, accounts
}

func GetOwner(Account string) (ownerID string, onWhitelist bool) {
	var (
		dataFound bool
		result    []bson.M
		dcUser    string
	)
	result, dataFound = mongodb.Read(whitelistCollection, bson.M{
		"mcAccount": strings.ToLower(Account),
	})
	if dataFound {
		dcUser = fmt.Sprintf("%v", result[0]["dcUserID"])
	}
	return dcUser, dataFound
}
