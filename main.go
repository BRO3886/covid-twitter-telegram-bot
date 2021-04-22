package main

import (
	"os"
	"time"

	"github.com/BRO3886/covid-twt-telegram/pkg"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
	bot := pkg.Telegram{
		BotToken: os.Getenv("TG_BOT_TOKEN"),
		ChatId:   os.Getenv("TG_CHAT_ID"),
	}

	// _ = pkg.TwtCredentials{
	// 	AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
	// 	AccessTokenSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
	// 	ConsumerKey:       os.Getenv("TWITTER_API_KEY"),
	// 	ConsumerSecret:    os.Getenv("TWITTER_API_SECRET"),
	// }

	// pkg.SearchTweets(creds)

	for {
		data := pkg.SearchV2()

		for _, msg := range data {
			pkg.PostTelegram(bot, msg)
			time.Sleep(time.Second * 5)
		}

		time.Sleep(time.Minute * 5)
	}

	// pkg.PostTelegram(tel)
}
