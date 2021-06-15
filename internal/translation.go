package internal

import (
	"encoding/json"
	"fmt"

	"trnsl8/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

var translateSession *translate.Translate

func GetAWSTranslateSessionInstance() *translate.Translate {
	if translateSession == nil {
		translateSession = translate.New(session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		})))
	}
	return translateSession
}

func TranslateToTargetLanguage(TARGET_LANGUAGE string, text *string) model.TranslationResult {
	var translationResult = model.TranslationResult{"", "", ""}
	var session = GetAWSTranslateSessionInstance()

	response, err := session.Text(&translate.TextInput{
		SourceLanguageCode: aws.String("auto"),
		TargetLanguageCode: aws.String(TARGET_LANGUAGE),
		Text:               text,
	})

	if err != nil {
		fmt.Println("Error executing 'trnsl8 to': \n", err)
	} else {
		var jsonbytes, _ = json.Marshal(response)
		err := json.Unmarshal(jsonbytes, &translationResult)

		if err != nil {
			fmt.Println(err)
		}
	}
	return translationResult
}
