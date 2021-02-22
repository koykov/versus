package jsonvector

import (
	"fmt"
	"testing"

	"github.com/valyala/fastjson"
)

func BenchmarkParseFastJSON(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchmarkParseFastJSON(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchmarkParseFastJSON(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchmarkParseFastJSON(b, largeFixture)
	})
	b.Run("canada", func(b *testing.B) {
		benchmarkParseFastJSON(b, canadaFixture)
	})
	b.Run("citm", func(b *testing.B) {
		benchmarkParseFastJSON(b, citmFixture)
	})
	b.Run("twitter", func(b *testing.B) {
		benchmarkParseFastJSON(b, twitterFixture)
	})
}

func benchmarkFastJSONParse(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		p := poolFJ.Get()
		for pb.Next() {
			v, err := p.Parse(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			if v.Type() != fastjson.TypeObject {
				panic(fmt.Errorf("unexpected value type; got %s; want %s", v.Type(), fastjson.TypeObject))
			}
		}
		poolFJ.Put(p)
	})
}

func benchmarkParseFastJSON(b *testing.B, s string) {
	b.Run("fastjson", func(b *testing.B) {
		benchmarkFastJSONParse(b, s)
	})
}
