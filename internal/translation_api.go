package translation_api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

var translateSession *translate.Translate

func InitAWSTranslateSessionInstance() {
	if translateSession == nil {
		translateSession = translate.New(session.Must(session.NewSession(
			&aws.Config{
				Region: aws.String("us-east-1"),
			})))
	}
}

func TranslateToTargetLanguage(TARGET_LANGUAGE string, text *string) {

	InitAWSTranslateSessionInstance()

	response, err := translateSession.Text(&translate.TextInput{
		SourceLanguageCode: aws.String("auto"),
		TargetLanguageCode: aws.String(TARGET_LANGUAGE),
		Text:               text,
	})

	if err != nil {
		fmt.Println("Error executing trnsl8: \n", err)
	} else {
		fmt.Println(response)
	}

}
