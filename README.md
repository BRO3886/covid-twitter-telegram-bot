# Covid-Twitter Stream Bot
A simple bot that uses [Twitter's Stream Search API](https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/quick-start) to send recent tweets that match the filter, to a [Telegram group](https://t.me/joinchat/Rx2aI5y42YtjOTA1).

## Steps to install
* Register for a Developer Account on twitter
* Make a bot and a group, add the bot to the group, fetch the group ID
* Make a .env file (Check [sample.env](.env.sample)) with the required keys
* Run `go run .`

## Filtered Stream FAQ
* Rule being used by me for Delhi NCR
```
verified (delhi OR Delhi OR Noida OR GURGAON) (bed OR beds OR icu OR oxygen OR ventilator OR ventilators OR plasma OR remdesivir OR remedesevir OR remedesivir OR medicine) available -\"not verified\" -unverified -needed -need -required -is:retweet -is:reply (#covidresources OR #delhicovid OR #covidsos OR #covidhelp OR #covidemergency OR #delhineedsoxygen OR #oxygenbeds OR #delhincr OR #remdesivir OR #Toclizumab OR #delhi) #verified
```
* See [Filtered Stream (Twitter API v2)](https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/introduction) to see how to add filtering rules



> Please ⭐ the repo if this helped you :)

The code is very shitty I wrote it in between exams 😭