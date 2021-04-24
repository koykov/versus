package jsonvector

import (
	"fmt"
	"testing"

	"github.com/koykov/jsonvector"
	"github.com/koykov/vector"
)

func TestParseJsonvector(t *testing.T) {
	p := jsonvector.Acquire()
	err := p.ParseStr(mediumFixture)
	if err != nil {
		t.Error(err)
	}
	jsonvector.Release(p)
}

func BenchmarkParseJsonvector(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchmarkParseJsonvector(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchmarkParseJsonvector(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchmarkParseJsonvector(b, largeFixture)
	})
	b.Run("canada", func(b *testing.B) {
		benchmarkParseJsonvector(b, canadaFixture)
	})
	b.Run("citm", func(b *testing.B) {
		benchmarkParseJsonvector(b, citmFixture)
	})
	b.Run("twitter", func(b *testing.B) {
		benchmarkParseJsonvector(b, twitterFixture)
	})
}

func benchmarkJsonvectorParse(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		p := jsonvector.Acquire()
		for pb.Next() {
			err := p.ParseStr(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			v := p.Get()
			if v.Type() != vector.TypeObj {
				panic(fmt.Errorf("unexpected value type; got %d; want %d", v.Type(), vector.TypeObj))
			}
			p.Reset()
		}
		jsonvector.Release(p)
	})
}

func benchmarkParseJsonvector(b *testing.B, s string) {
	b.Run("jsonvector", func(b *testing.B) {
		benchmarkJsonvectorParse(b, s)
	})
}
