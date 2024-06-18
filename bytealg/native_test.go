package bytealg

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koykov/byteconv"
)

func BenchmarkNative(b *testing.B) {
	b.Run("trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := bytes.Trim(trimOrigin, trimCutStr)
			if !bytes.Equal(r, trimExpect) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, byteconv.B2S(r), byteconv.B2S(trimExpect))
			}
		}
	})
	b.Run("trim str", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := strings.Trim(trimOriginS, trimCutS)
			if r != trimExpectS {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
			}
		}
	})
}
