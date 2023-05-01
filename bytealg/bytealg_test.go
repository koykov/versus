package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/bytealg"
	"github.com/koykov/fastconv"
)

func BenchmarkBytealg(b *testing.B) {
	b.Run("trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := bytealg.Trim(trimOrigin, trimCut)
			if !bytes.Equal(r, trimExpect) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
			}
		}
	})
	b.Run("trim str", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := bytealg.Trim(trimOriginS, trimCutS)
			if r != trimExpectS {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
			}
		}
	})
}
