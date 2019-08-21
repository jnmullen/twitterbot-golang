package image

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// RecogniseAnimalInImage test
func RecogniseAnimalInImage(imageData []byte) (string, error) {

	sess := session.New(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	svc := rekognition.New(sess)

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: imageData,
		},
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		fmt.Print(err)
		//do something
		return "", err
	}

	var foundAnimal bool
	var maximumParents int
	var animalResult string

	for _, label := range result.Labels {
		fmt.Printf("This label is %s is", *label.Name)
		if *label.Name == "Animal" {
			foundAnimal = true
		}
	}

	if foundAnimal {
		for _, label := range result.Labels {
			if len(label.Parents) > maximumParents {
				maximumParents = len(label.Parents)
				animalResult = *label.Name
			}
		}
		fmt.Printf("Found animal %s : Number of parents %d", animalResult, maximumParents)
		return animalResult, nil
	} else {
		//no animal detected
		return "", nil
	}

}
