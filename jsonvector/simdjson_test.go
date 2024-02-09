package jsonvector

import (
	"fmt"
	"testing"

	"github.com/koykov/byteconv"
	"github.com/minio/simdjson-go"
)

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
			_, err := simdjson.Parse(byteconv.S2B(s), nil)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
		}
	})
}
