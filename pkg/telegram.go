package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	log "github.com/sirupsen/logrus"
)

type Telegram struct {
	ChatId   string
	BotToken string
}

func PostTelegramMessage(tel Telegram, msg string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tel.BotToken)

	body := map[string]interface{}{
		"chat_id": tel.ChatId,
		"text":    msg,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshiling", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Error("Error making req: ", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error in response: ", err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		log.Info("message sent")
	} else {
		log.Warn("message not sent")
	}
}

func PostTelegramImage(tel Telegram, imgURL, caption string) {
	r := regexp.MustCompile(`https:\/\/t\.co\/[a-zA-Z0-9]+`)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", tel.BotToken)

	body := map[string]interface{}{
		"chat_id": tel.ChatId,
		"photo":   imgURL,
		"caption": r.ReplaceAllLiteralString(caption, ""),
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshiling", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Error("Error making req: ", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error in response: ", err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		log.Info("message sent")
	} else {
		log.Warn("message not sent: ", resp.StatusCode)
		log.Info("url was: ", imgURL)
		bbytes, _ := ioutil.ReadAll(resp.Body)
		log.Error(string(bbytes))
	}

}
