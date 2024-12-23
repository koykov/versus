package bytefuzz

import (
	"testing"

	"github.com/jamesturk/go-jellyfish"
	"github.com/koykov/bytefuzz/levenshtein"
	"github.com/koykov/bytefuzz/levenshtein2"
)

func BenchmarkLevenshtein(b *testing.B) {
	var stages = []stage{
		{
			text:     "",
			target:   "a",
			distance: 1,
		},
		{
			text:     "a",
			target:   "aa",
			distance: 1,
		},
		{
			text:     "a",
			target:   "aaa",
			distance: 2,
		},
		{
			text:     "",
			target:   "",
			distance: 0,
		},
		{
			text:     "a",
			target:   "b",
			distance: 2,
		},
		{
			text:     "aaa",
			target:   "aba",
			distance: 2,
		},
		{
			text:     "aaa",
			target:   "ab",
			distance: 3,
		},
		{
			text:     "a",
			target:   "a",
			distance: 0,
		},
		{
			text:     "ab",
			target:   "ab",
			distance: 0,
		},
		{
			text:     "a",
			target:   "",
			distance: 1,
		},
		{
			text:     "aa",
			target:   "a",
			distance: 1,
		},
		{
			text:     "aaa",
			target:   "a",
			distance: 2,
		},
		{
			text:     "kitten",
			target:   "sitting",
			distance: 5,
		},
		{
			text:     "Orange",
			target:   "Apple",
			distance: 9,
		},
		{
			text:     "ab",
			target:   "bc",
			distance: 2,
		},
		{
			text:     "abd",
			target:   "bec",
			distance: 4,
		},
		{
			text:     "me",
			target:   "meme",
			distance: 2,
		},
	}

	b.Run("bytefuzz", func(b *testing.B) {
		for _, st := range stages {
			b.Run(st.text, func(b *testing.B) {
				b.ReportAllocs()
				ctx := levenshtein.Acquire()
				defer levenshtein.Release(ctx)
				for i := 0; i < b.N; i++ {
					ctx.Reset()
					ctx.DistanceString(st.text, st.target)
				}
			})
		}
	})
	b.Run("bytefuzz/2", func(b *testing.B) {
		for _, st := range stages {
			b.Run(st.text, func(b *testing.B) {
				b.ReportAllocs()
				ctx := levenshtein2.Acquire()
				defer levenshtein2.Release(ctx)
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
					jellyfish.Levenshtein(st.text, st.target)
				}
			})
		}
	})
}
