package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/bytealg"
	"github.com/koykov/byteconv"
)

func BenchmarkBytealg(b *testing.B) {
	b.Run("trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := bytealg.TrimBytes(trimOrigin, trimCut)
			if !bytes.Equal(r, trimExpect) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpect))
			}
		}
	})
	b.Run("trim str", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := bytealg.TrimString(trimOriginS, trimCutS)
			if r != trimExpectS {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
			}
		}
	})
}
