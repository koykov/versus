package main

import (
	"testing"
	"unicode"

	"github.com/koykov/versus/nlp_script/testdata"
)

func BenchmarkEvaluate(b *testing.B) {
	l := len(testdata.Model)
	b.Run("native", func(b *testing.B) {
		var y bool
		for i := 0; i < b.N; i++ {
			x := &testdata.Model[i%l]
			for j := 0; j < len(scripts); j++ {
				y = unicode.Is(scripts[j].b, x.C)
			}
			_ = y
		}
	})
	b.Run("nlp", func(b *testing.B) {
		var y bool
		for i := 0; i < b.N; i++ {
			x := &testdata.Model[i%l]
			for j := 0; j < len(scripts); j++ {
				y = scripts[j].a.EvaluateRune(x.C)
			}
			_ = y
		}
	})
}
