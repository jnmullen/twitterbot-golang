package twitter

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var client *twitter.Client

func init() {
	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
}

// a function to reply to a tweet
func ReplyToTweet(tweetContent string, replyToStatusID int64) bool {

	tweetParams := &twitter.StatusUpdateParams{
		Status:            tweetContent,
		InReplyToStatusID: replyToStatusID,
	}

	// for now ignore tweet and status
	_, _, err := client.Statuses.Update(tweetContent, tweetParams)
	if err != nil {
		return false
	}
	return true
}
