package main

import (
	"fmt"
	"helloworld/image"
	"helloworld/twitter"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest some comment to keep the compiler happy
func HandleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {

	tweetData := twitter.ParseTweetData([]byte(request.Body))

	imageData := image.DownloadImage(tweetData.ImageURL)

	animal, _ := image.RecogniseAnimalInImage(imageData)

	var tweetContent string
	if len(animal) > 0 {
		tweetContent = fmt.Sprintf("@%s I can see a : %s", tweetData.ScreenName, animal)
	} else {
		tweetContent = "Sorry that picture is not of an animal"
	}

	tweetSent := twitter.ReplyToTweet(tweetContent, tweetData.TweetID)
	if tweetSent {
		fmt.Printf("Replied to tweet ID %d successfully", tweetData.TweetID)
	} else {
		fmt.Printf("Failed to reply to tweet ID %d", tweetData.TweetID)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       strconv.FormatBool(tweetSent),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
