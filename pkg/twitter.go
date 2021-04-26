package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
)

func UnmarshalData(data []byte) (TweetData, error) {
	var r TweetData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TweetData) marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TweetData struct {
	Data          DataClass      `json:"data"`
	Includes      Includes       `json:"includes"`
	MatchingRules []MatchingRule `json:"matching_rules"`
}

type DataClass struct {
	Entities  Entities  `json:"entities"`
	AuthorID  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	ID        string    `json:"id"`
}

type Entities struct {
	Hashtags []Hashtag `json:"hashtags"`
	Mentions []Mention `json:"mentions"`
	Urls     []URL     `json:"urls"`
}

type Hashtag struct {
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Tag   string `json:"tag"`
}

type Mention struct {
	Start    int64  `json:"start"`
	End      int64  `json:"end"`
	Username string `json:"username"`
}

type URL struct {
	Start       int64   `json:"start"`
	End         int64   `json:"end"`
	URL         string  `json:"url"`
	ExpandedURL string  `json:"expanded_url"`
	DisplayURL  string  `json:"display_url"`
	Images      []Image `json:"images"`
	Status      int64   `json:"status"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnwoundURL  string  `json:"unwound_url"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Includes struct {
	Users []User `json:"users"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type MatchingRule struct {
	ID  int64  `json:"id"`
	Tag string `json:"tag"`
}

type TwitterClient struct {
	Token string
}

func StreamSearch(twitter TwitterClient, bot TelegramBot) {
	r := regexp.MustCompile("@")
	loc, _ := time.LoadLocation("Asia/Kolkata")

	query := "tweet.fields=attachments,created_at,entities,author_id&expansions=attachments.media_keys,author_id&media.fields=duration_ms,height,media_key,preview_image_url,public_metrics,type,url,width"
	url := fmt.Sprintf("https://api.twitter.com/2/tweets/search/stream?%s", query)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error("Error making http request: ", err)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", twitter.Token))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Error("Error reading lines: ", err)
			time.Sleep(time.Second * 2)
			continue
		}

		if len(line) == 0 {
			continue
		}

		tweet, err := UnmarshalData(line)
		if err != nil {
			log.Error("error unmarshaling: ", err)
			time.Sleep(time.Second * 2)
			continue
		}
		link := fmt.Sprintf("[Twitter Link](https://twitter.com/%s/status/%s/)", tweet.Includes.Users[0].Username, tweet.Data.ID)
		message := fmt.Sprintf("\n\n🕘 %s %s\n\n", tweet.Data.CreatedAt.In(loc).Format("Mon, Jan 2"), tweet.Data.CreatedAt.In(loc).Format(time.Kitchen))
		message += fmt.Sprintf("%s on Twitter:\n", tweet.Includes.Users[0].Name)
		log.Info(link, message, tweet.Data.Text)
		message += r.ReplaceAllString(tweet.Data.Text, " ")

		// urls := tweet.Data.Entities.Urls
		// if len(urls) > 0 {
		// 	log.Println("probably has image")
		// 	urls[0].
		// }

		PostTelegramMessage(bot, link, message)
		time.Sleep(time.Second * 1)
	}

}
