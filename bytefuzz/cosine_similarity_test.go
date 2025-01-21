package bytefuzz

import (
	"testing"

	"github.com/koykov/bytefuzz/cosine_similarity"
)

func BenchmarkCosineSimilarity(b *testing.B) {
	var stages = []stage{
		{"apple", "applet", 0.9722222222222223},
		{"Michael", "Michelle", 0.8880952380952382},
		{"McDonald's", "Mcdonells", 0.8777777777777778},
		{"hello", "world", 0.4666666666666666},
		{"telephone", "telephne", 0.9346560846560846},
		{"AI", "Artificial Intelligence", 0},
		{"abacus", "abaxus", 0.8444444444444443},
		{"Thompson", "Thomson", 0.8278571428571428},
	}
	b.Run("bytefuzz", func(b *testing.B) {
		for _, st := range stages {
			b.Run(st.text, func(b *testing.B) {
				b.ReportAllocs()
				ctx := cosine_similarity.Acquire()
				defer cosine_similarity.Release(ctx)
				for i := 0; i < b.N; i++ {
					ctx.Reset()
					ctx.DistanceString(st.text, st.target)
				}
			})
		}
	})
}
