package pkg

// import (
// 	"fmt"

// 	log "github.com/sirupsen/logrus"

// 	"github.com/dghubble/go-twitter/twitter"
// 	"github.com/dghubble/oauth1"
// )

// type TwtCredentials struct {
// 	ConsumerKey       string
// 	ConsumerSecret    string
// 	AccessToken       string
// 	AccessTokenSecret string
// }

// func getClient(creds *TwtCredentials) (*twitter.Client, error) {
// 	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
// 	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
// 	// Pass in your Access Token and your Access Token Secret
// 	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

// 	httpClient := config.Client(oauth1.NoContext, token)
// 	client := twitter.NewClient(httpClient)

// 	// Verify Credentials
// 	verifyParams := &twitter.AccountVerifyParams{
// 		SkipStatus:   twitter.Bool(true),
// 		IncludeEmail: twitter.Bool(true),
// 	}

// 	// we can retrieve the user and verify if the credentials
// 	// we have used successfully allow us to log in!
// 	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// log.Printf("User's ACCOUNT:\n%+v\n", user)
// 	return client, nil
// }

// func SearchTweets(creds TwtCredentials) {
// 	client, err := getClient(&creds)
// 	if err != nil {
// 		log.Error("Error getting client: ", err.Error())
// 	}
// 	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
// 		Query: `verified Delhi (bed OR beds OR icu OR oxygen OR ventilator OR ventilators OR fabiflu) -"not verified" -"unverified" -"needed" -"required"`,
// 		TweetMode: "Extended",
// 	})
// 	if err != nil {
// 		log.Error("Error searching: ", err.Error())
// 	}

// 	for _, tweet := range search.Statuses {
// 		// log.Info(tweet)
// 		fmt.Println(tweet.ID, tweet.User.ID, tweet.Text, tweet.FullText)
// 		// imgCount := len(tweet.Entities.Media)
// 		// if imgCount > 0 {
// 		// 	media := tweet.Entities.Media[0]
// 		// 	log.Info(media.MediaURL)
// 		// }
// 	}

// 	// params := &twitter.StreamFilterParams{
// 	// 	Track:         []string{"apple"},
// 	// 	StallWarnings: twitter.Bool(true),
// 	// }

// 	// stream, _ := client.Streams.Filter(params)
// 	// demux := twitter.NewSwitchDemux()

// 	// demux.Tweet = func(tweet *twitter.Tweet) {
// 	// 	fmt.Println(tweet.Text)
// 	// }

// 	// for message := range stream.Messages {
// 	// 	demux.Handle(message)
// 	// }
// }
