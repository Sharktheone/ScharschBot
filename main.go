package main

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	bot *discordgo.Session
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	var BotToken string
	flag.StringVar(&BotToken, "t", os.Getenv("BOT_TOKEN"), "Bot Token")
	flag.Parse()

	bot, err = discordgo.New(BotToken)
	if err != nil {
		log.Fatal("Invalid Bot Configuration:", err)
	}
}

func main() {
	err := bot.Open()
	if err != nil {
		log.Fatal("Cannot open connection:", err)
	}

	defer bot.Close()
}
