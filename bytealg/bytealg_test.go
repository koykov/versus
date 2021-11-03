package bytealg

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koykov/bytealg"
	"github.com/koykov/fastconv"
)

var (
	trimOrigin = []byte("..foo bar!!???")
	trimExpect = []byte("foo bar")
	trimCutStr = "?!."
	trimCut    = []byte(trimCutStr)

	trimOriginS = "..foo bar!!???"
	trimExpectS = "foo bar"
	trimCutS    = "?!."
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
			r := bytealg.TrimStr(trimOriginS, trimCutS)
			if r != trimExpectS {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
			}
		}
	})
}

func BenchmarkNative(b *testing.B) {
	b.Run("trim", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := bytes.Trim(trimOrigin, trimCutStr)
			if !bytes.Equal(r, trimExpect) {
				b.Errorf(`Trim: mismatch result %s and expectation %s`, fastconv.B2S(r), fastconv.B2S(trimExpect))
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
