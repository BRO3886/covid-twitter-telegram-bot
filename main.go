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
	// log.Info(bot.ChatId + "92383")

	for {
		data := pkg.SearchV2()
		// log.Println(data)
		for _, msg := range data {
			if msg.HasURL {
				go pkg.PostTelegramImage(bot, msg.URL, msg.Message)
			} else {
				pkg.PostTelegramMessage(bot, msg.Message)
			}
			time.Sleep(time.Second * 3)
		}

		time.Sleep(time.Minute * 2)
	}

}
