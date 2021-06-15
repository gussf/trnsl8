package model

import "fmt"

type TranslationResult struct {
	SourceLanguage string `json:"SourceLanguageCode"`
	TargetLanguage string `json:"TargetLanguageCode"`
	Translation    string `json:"TranslatedText"`
}

func (t TranslationResult) String() string {
	return fmt.Sprintf("Source Language [%s] (automatically detected)\nTarget Language [%s]\nTranslation     [%s]", t.SourceLanguage, t.TargetLanguage, t.Translation)
}
