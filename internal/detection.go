package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"trnsl8/model"

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

func DetectDominantLanguageIn(input *string) model.DetectedLanguage {
	var detectionResult = model.DetectedLanguage{"", 0}
	var session = GetAWSComprehendSessionInstance()

	response, err := session.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{
		Text: input,
	})

	if err != nil {
		fmt.Println("Error executing 'trnsl8 detect': \n", err)
	} else {
		var resultArray model.DetectedLanguagesList
		var jsonbytes, _ = json.Marshal(response)
		err := json.Unmarshal(jsonbytes, &resultArray)

		if err != nil {
			fmt.Println(err)
			return detectionResult
		}
		detectionResult = resultArray.Languages[0]
	}
	return detectionResult

}

func DetectDominantLanguageInFile(filePath string) model.DetectedLanguage {

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to open file %s", filePath)
		return model.DetectedLanguage{"", 0}
	}
	var dataString = string(dat)
	return DetectDominantLanguageIn(&dataString)
}
