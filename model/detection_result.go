package model

import (
	"fmt"
)

type DetectedLanguagesList struct {
	Languages []DetectedLanguage `json:"Languages"`
}

type DetectedLanguage struct {
	DetectedLanguageCode string  `json:"LanguageCode"`
	Score                float64 `json:"Score"`
}

func (d DetectedLanguage) String() string {
	return fmt.Sprintf("Dominant Language [%s]\nScore [%.0f%%]", d.DetectedLanguageCode, d.Score)
}
