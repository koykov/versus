package nlp

import (
	"testing"

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

func TestPredictLanguage(t *testing.T) {
	t.Run("hello world", func(t *testing.T) {
		var c, p, n int
		for _, s := range stagesLPHW {
			lang := ldLinguaPredict(s.text, "English")
			if lang == s.lang {
				p++
			} else {
				n++
			}
			c++
		}
		t.Logf("stages %d: positive ld %d; negative ld %d", c, p, n)
	})
	t.Run("dataset", func(t *testing.T) {
		var c, p, n int
		for _, s := range stageLPDS {
			lang := ldLinguaPredict(s.text, "English")
			if lang == s.lang {
				p++
			} else {
				n++
			}
			c++
		}
		t.Logf("stages %d: positive ld %d; negative ld %d", c, p, n)
	})
}
