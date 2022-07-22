package whitelist

import (
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"github.com/Sharktheone/Scharsch-bot-discord/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var config = conf.GetConf()

func Add(username string, UserID string) {
	mongodb.Write(config.Whitelist.Mongodb.MongodbCollectionName, bson.D{
		{"dcUserID", UserID},
		{"mcAccount", username},
	})

	log.Println("*Add " + username + " to whitelist")
}
