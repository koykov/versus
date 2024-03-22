package xmlvector

import (
	"fmt"
	"testing"

	"github.com/koykov/vector"
	"github.com/koykov/xmlvector"
)

func TestParseXmlvector(t *testing.T) {
	p := xmlvector.Acquire()
	err := p.ParseStr(mediumFixture)
	if err != nil {
		t.Error(err)
	}
	xmlvector.Release(p)
}

func BenchmarkParseXmlvector(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchXmlvector(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchXmlvector(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchXmlvector(b, largeFixture)
	})
	b.Run("sigmod", func(b *testing.B) {
		benchXmlvector(b, sigmodFixture)
	})
	b.Run("mondial", func(b *testing.B) {
		benchXmlvector(b, mondialFixture)
	})
}

func benchXmlvector(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p := xmlvector.Acquire()
			err := p.ParseStr(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			v := p.Get()
			if v.Type() != vector.TypeObj {
				panic(fmt.Errorf("unexpected value type; got %d; want %d", v.Type(), vector.TypeObj))
			}
			xmlvector.Release(p)
		}
	})
}
