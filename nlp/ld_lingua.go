package nlp

import (
	"github.com/pemistahl/lingua-go"
)

var (
	ldLingua lingua.LanguageDetector
)

func init() {
	ldLingua = lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()
}

func ldLinguaPredict(s, def string) string {
	if lang, ok := ldLingua.DetectLanguageOf(s); ok {
		return lang.String()
	}
	return def
}
