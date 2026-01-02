package jsonvector

import (
	"fmt"
	"testing"

	"github.com/koykov/jsonvector"
	"github.com/koykov/vector"
)

func TestParseJsonvector(t *testing.T) {
	p := jsonvector.Acquire()
	err := p.ParseString(mediumFixture)
	if err != nil {
		t.Error(err)
	}
	jsonvector.Release(p)
}

func BenchmarkParseJsonvector(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchJsonvector(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchJsonvector(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchJsonvector(b, largeFixture)
	})
	b.Run("canada", func(b *testing.B) {
		benchJsonvector(b, canadaFixture)
	})
	b.Run("citm", func(b *testing.B) {
		benchJsonvector(b, citmFixture)
	})
	b.Run("twitter", func(b *testing.B) {
		benchJsonvector(b, twitterFixture)
	})
}

func benchJsonvector(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p := jsonvector.Acquire()
			err := p.ParseString(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			v := p.Get()
			if v.Type() != vector.TypeObject {
				panic(fmt.Errorf("unexpected value type; got %d; want %d", v.Type(), vector.TypeObject))
			}
			jsonvector.ReleaseNC(p)
		}
	})
}
