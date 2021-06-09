package api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

var comprehendSession *comprehend.Comprehend

func GetAWSComprehendSessionInstance() *comprehend.Comprehend {
	if comprehendSession == nil {
		comprehendSession = comprehend.New(session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})))
	}
	return comprehendSession
}

func DetectDominantLanguageIn(input *string) {
	var session = GetAWSComprehendSessionInstance()
	response, err := session.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{
		Text: input,
	})

	if err != nil {
		fmt.Println("Error executing trnsl8 detect: \n", err)
	} else {
		fmt.Println(response)
	}

}
