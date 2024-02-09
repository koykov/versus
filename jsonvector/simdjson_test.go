package jsonvector

import (
	"fmt"
	"sync"
	"testing"

	"github.com/koykov/byteconv"
	"github.com/minio/simdjson-go"
)

var poolSimdjson = sync.Pool{New: func() any {
	return &simdjson.ParsedJson{}
}}

func BenchmarkParseSimdjson(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchSimdjson(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchSimdjson(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchSimdjson(b, largeFixture)
	})
	b.Run("canada", func(b *testing.B) {
		benchSimdjson(b, canadaFixture)
	})
	b.Run("citm", func(b *testing.B) {
		benchSimdjson(b, citmFixture)
	})
	b.Run("twitter", func(b *testing.B) {
		benchSimdjson(b, twitterFixture)
	})
}

func benchSimdjson(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var err error
			parsed := poolSimdjson.Get().(*simdjson.ParsedJson)
			if parsed, err = simdjson.Parse(byteconv.S2B(s), parsed); err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			poolSimdjson.Put(parsed)
		}
	})
}
