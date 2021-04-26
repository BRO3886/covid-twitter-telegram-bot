package main

import (
	"os"

	"github.com/BRO3886/covid-twt-telegram/pkg"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	bot := pkg.TelegramBot{
		BotToken: os.Getenv("TG_BOT_TOKEN"),
		ChatId:   os.Getenv("TG_CHAT_ID"),
		// BotToken: os.Getenv("CT_TOKEN"),
		// ChatId:   os.Getenv("CT_ID"),
	}

	twtClient := pkg.TwitterClient{
		Token: os.Getenv("TWITTER_BEARER_TOKEN"),
	}

	pkg.StreamSearch(twtClient, bot)

	// testLink := "[Link](https://twitter.com/anonyy_mouse/status/1385842298198499334)"
	// testMessage := "@DammnGirll #Ventilator Beds are Available !\n\nLocation :- New Delhi. NCR\n\nContact : 96545 - 35285 \n\n#Verified by @Bae_Hey_Yaa at 6:53 Pm on 26-04-2021. 🔥\n\n(Please call on this &amp; ask if they have icu beds.)"
	// pkg.PostTelegramMessage(bot, testLink, testMessage)

	// log.Info(bot.ChatId + "92383")

	// for {
	// 	data := pkg.SearchV2()
	// 	// log.Println(data)
	// 	for _, msg := range data {
	// 		if msg.HasURL {
	// 			pkg.PostTelegramImage(bot, msg.URL, msg.Message)
	// 		} else {
	// 			pkg.PostTelegramMessage(bot, msg.Message)
	// 		}
	// 		time.Sleep(time.Second * 4)
	// 	}

	// 	time.Sleep(time.Minute * 2)
	// }

}
