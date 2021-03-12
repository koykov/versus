package bytealg

import (
	"bytes"
	"testing"

	"github.com/koykov/bytealg"
	"github.com/koykov/fastconv"
)

var (
	trimOrigin = []byte("..foo bar!!???")
	trimExpect = []byte("foo bar")
	trimCutStr = "?!."
	trimCut    = []byte(trimCutStr)
)

func BenchmarkTrim_Bytealg(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytealg.Trim(trimOrigin, trimCut)
		if !bytes.Equal(r, trimExpect) {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	}
}

func BenchmarkTrim_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytes.Trim(trimOrigin, trimCutStr)
		if !bytes.Equal(r, trimExpect) {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
		}
	}
}
