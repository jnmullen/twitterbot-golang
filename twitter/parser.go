package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ProcessedData struct {
	TweetID    int64
	ScreenName string
	ImageURL   string
}

// Data comment
type Data struct {
	Tweets []struct {
		ID       string `json:"id_str"`
		Entities struct {
			Urls []struct {
				URL string `json:"url"`
			} `json:"urls"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []struct {
				MediaURL string `json:"media_url"`
			} `json:"media"`
		} `json:"extended_entities"`
		User struct {
			ScreenName string `json:"screen_name"`
		} `json:"user"`
	} `json:"tweet_create_events"`
}

// ParseTweetData blah
func ParseTweetData(body []byte) ProcessedData {

	var tweetData Data
	var response ProcessedData
	err := json.Unmarshal(body, &tweetData)
	if err != nil {
		fmt.Println(err)
	}

	var urlFromEntities string
	if (len(tweetData.Tweets[0].Entities.Urls)) > 0 {
		urlFromEntities = tweetData.Tweets[0].Entities.Urls[0].URL
	}

	var urlFromExtendedEntities string
	if (len(tweetData.Tweets[0].ExtendedEntities.Media)) > 0 {
		urlFromExtendedEntities = tweetData.Tweets[0].ExtendedEntities.Media[0].MediaURL
	}

	var tweetID = tweetData.Tweets[0].ID
	var twitterScreenName = tweetData.Tweets[0].User.ScreenName

	intTweetID, err := strconv.ParseInt(tweetID, 10, 64)
	if err == nil {
		response.TweetID = intTweetID
	}

	response.ScreenName = twitterScreenName
	if urlFromEntities == "" && urlFromExtendedEntities == "" {
		response.ImageURL = ""
	} else if urlFromEntities != "" {
		response.ImageURL = urlFromEntities
	} else {
		response.ImageURL = urlFromExtendedEntities
	}

	fmt.Printf("ID : %d\n", intTweetID)
	fmt.Printf("ScreenName : %s\n", twitterScreenName)
	fmt.Printf("Url From Entities : %s\n", urlFromEntities)
	fmt.Printf("Url From Extended Entities : %s\n", urlFromExtendedEntities)

	return response
}
