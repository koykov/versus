package nlp_language

import (
	"testing"

	"github.com/koykov/nlp"
)

func BenchmarkAccess(b *testing.B) {
	b.Run("nlp", func(b *testing.B) {
		b.ReportAllocs()
		var x string
		for i := 0; i < b.N; i++ {
			x = nlp.LanguageRussian.String()
		}
		_ = x
	})
	b.Run("stringer", func(b *testing.B) {
		b.ReportAllocs()
		var x string
		for i := 0; i < b.N; i++ {
			x = Russian.String()
		}
		_ = x
	})
}
