package stringer

import "testing"

func BenchmarkAccess(b *testing.B) {
	b.Run("nlp", func(b *testing.B) {
		b.ReportAllocs()
		var x string
		for i := 0; i < b.N; i++ {
			x = Russian.Name()
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
