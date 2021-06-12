package api

import (
	"fmt"
	"io/ioutil"

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

func DetectDominantLanguageIn(input *string) (result string, err error) {
	var session = GetAWSComprehendSessionInstance()
	response, err := session.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{
		Text: input,
	})
	result = response.GoString()
	return result, err

}

func DetectDominantLanguageInFile(filePath string) (result string, err error) {

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s", filePath)
	}
	var dataString = string(dat)
	result, err = DetectDominantLanguageIn(&dataString)
	return result, err

}
