package main

import (
	"fmt"
	"jamesm/animalbot/image"
	"jamesm/animalbot/twitter"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest is the entry point to the lambda function
func HandleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {

	fmt.Println("Start to recognise an animal")

	fmt.Printf("%s", request.Body)

	tweetData, err := twitter.ParseTweetData([]byte(request.Body))
	if err != nil {
		return response, err
	}

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
