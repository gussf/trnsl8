package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/trnsl8/model"

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

func TranslateToTargetLanguage(TARGET_LANGUAGE string, text *string) (model.TranslationResult, error) {
	var translationResult = model.TranslationResult{"", "", ""}
	var session = GetAWSTranslateSessionInstance()

	response, err := session.Text(&translate.TextInput{
		SourceLanguageCode: aws.String("auto"),
		TargetLanguageCode: aws.String(TARGET_LANGUAGE),
		Text:               text,
	})

	if err == nil {
		var jsonbytes, _ = json.Marshal(response)
		err = json.Unmarshal(jsonbytes, &translationResult)
	}

	return translationResult, err
}

func TranslateToTargetLanguageFromFile(TARGET_LANGUAGE string, filePath string) (model.TranslationResult, error) {

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to open file [%s]\n", filePath)
		return model.TranslationResult{"", "", ""}, err
	}
	var dataString = string(dat)
	return TranslateToTargetLanguage(TARGET_LANGUAGE, &dataString)
}

func WriteTranslationToFile(translationResult model.TranslationResult, filePath string) (string, error) {

	dataToWrite := fmt.Sprintf("Source [%s]\nTarget [%s]\n\n%s\n",
		translationResult.SourceLanguage,
		translationResult.TargetLanguage,
		translationResult.Translation)

	byteData := []byte(dataToWrite)
	err := ioutil.WriteFile(filePath, byteData, 0644)
	if err != nil {
		fmt.Printf("Failed to write to file [%s]\n", filePath)
		return "Failed", err
	}
	return "Success", nil
}
