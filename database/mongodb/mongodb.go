package mongodb

import (
	"context"
	"fmt"
	"github.com/Sharktheone/Scharsch-bot-discord/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	config             = conf.GetConf()
	url                = fmt.Sprintf("mongodb://%v:%v@%v:%v", config.Whitelist.Mongodb.MongodbUser, config.Whitelist.Mongodb.MongodbPass, config.Whitelist.Mongodb.MongodbHost, config.Whitelist.Mongodb.MongodbPort)
	Client, connectErr = mongo.NewClient(options.Client().ApplyURI(url))
	ctx, Cancel        = context.WithTimeout(context.Background(), 10*time.Second)
	connected          = false
	db                 *mongo.Database
	whitelist          *mongo.Collection
)

func Connect() *mongo.Database {
	if !connected {
		if connectErr != nil {
			log.Fatalf("Failed to apply mongo URI: %v", connectErr)
		}

		err := Client.Connect(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		} else {
			connected = true
		}
	}
	db = Client.Database(config.Whitelist.Mongodb.MongodbDatabaseName)
	return db
}

func Disconnect() {
	err := Client.Disconnect(ctx)
	if err != nil {
		log.Printf("Failed to disconnect: %v \n", err)
	}
}

func Write(collection string, data bson.D) {
	whitelist = db.Collection(collection)
	_, err := whitelist.InsertOne(ctx, data)
	if err != nil {
		log.Printf("Failed o write data: %v", err)
	}
}
