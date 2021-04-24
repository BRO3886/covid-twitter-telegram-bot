package pkg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type authorize struct {
	Token string
}

type TelegramTweet struct {
	Message string
	URL     string
	HasURL  bool
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func SearchV2() []TelegramTweet {
	token := os.Getenv("TWITTER_BEARER_TOKEN")

	client := &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.TweetRecentSearchOpts{
		Expansions:  []twitter.Expansion{twitter.ExpansionEntitiesMentionsUserName, twitter.ExpansionAuthorID, twitter.ExpansionAttachmentsMediaKeys},
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAttachments, "entities"},
		MediaFields: []twitter.MediaField{twitter.MediaFieldPreviewImageURL, twitter.MediaFieldURL},
		StartTime:   time.Now().Add(-time.Hour * 3),
	}

	reqStrings := []string{
		`verified (Delhi OR Noida OR GURGAON) (bed OR beds OR icu OR oxygen OR ventilator OR ventilators OR plasma OR remdesivir OR remedesevir OR remedesivir OR medicine) available -"not verified" -"unverified" -"needed" -"need" -"required"`,
		`verified (Indore) (icu OR ventilator OR ventilators ) available -"not verified" -"unverified" -"needed" -"need" -"required "`,
		`apple iphone`, //test
	}

	tweetResponse, err := client.TweetRecentSearch(context.Background(), reqStrings[0], opts)
	if err != nil {
		log.Panicf("tweet lookup error: %v", err)
	}

	tweets := tweetResponse.Raw.Tweets

	r := regexp.MustCompile("@")

	messages := []TelegramTweet{}
	for _, tweet := range tweets {
		t := TelegramTweet{}

		urls := tweet.Entities.URLs
		// log.Info("urls length is ", len(urls))
		if len(urls) > 0 {
			log.Println("has image")
			t.HasURL = true
			t.URL = urls[0].DisplayURL
		}

		message := fmt.Sprintf("https://twitter.com/%s/status/%s\n\n", tweet.AuthorID, tweet.ID)
		fmt.Println(message)
		message += r.ReplaceAllString(strings.Replace(tweet.Text, "RT ", "", 1), "")
		t.Message = message

		messages = append(messages, t)
	}

	return messages
}
