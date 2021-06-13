package model

type DetectedLanguagesList struct {
	Languages []DetectedLanguage `json:"Languages"`
}

type DetectedLanguage struct {
	DetectedLanguageCode string  `json:"LanguageCode"`
	Score                float64 `json:"Score"`
}
