package pkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func SearchV2() []string {
	token := os.Getenv("TWITTER_BEARER_TOKEN")

	client := &twitter.Client{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.TweetRecentSearchOpts{
		Expansions:  []twitter.Expansion{twitter.ExpansionEntitiesMentionsUserName, twitter.ExpansionAuthorID},
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldConversationID, twitter.TweetFieldText},
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

	messages := []string{}
	for _, tweet := range tweets {
		message := fmt.Sprintf("https://twitter.com/%s/status/%s\n", tweet.AuthorID, tweet.ID)
		fmt.Println(message)
		message += strings.Replace(tweet.Text, "RT ", "", 1)
		messages = append(messages, message)
	}

	return messages
}
