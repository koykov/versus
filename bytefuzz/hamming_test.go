package bytefuzz

import (
	"testing"

	"github.com/jamesturk/go-jellyfish"
	"github.com/koykov/bytefuzz/hamming"
)

func BenchmarkHamming(b *testing.B) {
	var stages = []stage{
		{"", "", 0},
		{"A", "A", 0},
		{"G", "T", 1},
		{"GGACTGAAATCTG", "GGACTGAAATCTG", 0},
		{"GGACGGATTCTG", "AGGACGGATTCT", 9},
		{"AATG", "AAA", 2},
		{"", "G", 1},
		{"G", "", 1},
	}
	b.Run("bytefuzz", func(b *testing.B) {
		for _, st := range stages {
			b.Run(st.text, func(b *testing.B) {
				b.ReportAllocs()
				ctx := hamming.Acquire()
				defer hamming.Release(ctx)
				for i := 0; i < b.N; i++ {
					ctx.Reset()
					ctx.DistanceString(st.text, st.target)
				}
			})
		}
	})
	b.Run("jellyfish", func(b *testing.B) {
		for _, st := range stages {
			b.Run(st.text, func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					jellyfish.Hamming(st.text, st.target)
				}
			})
		}
	})
}
