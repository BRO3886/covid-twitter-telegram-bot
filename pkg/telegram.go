package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type TelegramBot struct {
	ChatId   string
	BotToken string
}

func PostTelegramMessage(tel TelegramBot, msg string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tel.BotToken)

	body := map[string]interface{}{
		"chat_id":    tel.ChatId,
		"text":       getTGEscapedMardown(msg),
		"parse_mode": "Markdown",
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Error("Error marshaling: ", err.Error())
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
	checkResponse(resp, "", msg)
}

func PostTelegramImage(tel TelegramBot, imgURL, caption string) {
	r := regexp.MustCompile(`https:\/\/t\.co\/[a-zA-Z0-9]+|&amp;`)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", tel.BotToken)

	body := map[string]interface{}{
		"chat_id": tel.ChatId,
		"photo":   imgURL,
		"caption": getTGEscapedMardown(r.ReplaceAllLiteralString(caption, "")),
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
	checkResponse(resp, imgURL, caption)
}

func checkResponse(resp *http.Response, imgURL, message string) {
	if resp.StatusCode == http.StatusOK {
		log.Info("message sent")
	} else {
		log.Warn("message not sent: ", resp.StatusCode)
		log.Info("url was: ", imgURL)
		log.Info("message was: ", message)
		bbytes, _ := ioutil.ReadAll(resp.Body)
		log.Error(string(bbytes))
	}
}

func getTGEscapedMardown(msg string) string {
	r := strings.NewReplacer("*", "\\*", "_", "\\_", "`", "\\`")
	return r.Replace(msg)
}
