package halvector

import (
	"fmt"
	"testing"

	"github.com/koykov/halvector"
	"github.com/koykov/vector"
)

func TestParseHalvector(t *testing.T) {
	p := halvector.Acquire()
	err := p.ParseStr(mediumFixture)
	if err != nil {
		t.Error(err)
	}
	halvector.Release(p)
}

func BenchmarkParseHalvector(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchHalvector(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchHalvector(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchHalvector(b, largeFixture)
	})
}

func benchHalvector(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p := halvector.Acquire()
			err := p.ParseStr(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			v := p.Get()
			if v.Type() != vector.TypeArr {
				panic(fmt.Errorf("unexpected value type; got %d; want %d", v.Type(), vector.TypeObj))
			}
			halvector.Release(p)
		}
	})
}
