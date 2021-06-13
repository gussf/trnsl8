package model

type TranslationResult struct {
	SourceLanguage string `json:"SourceLanguageCode"`
	TargetLanguage string `json:"TargetLanguageCode"`
	Translation    string `json:"TranslatedText"`
}
