package bytealg

import (
	"strings"
	"testing"

	"github.com/koykov/bytealg"
)

var (
	trimOriginS = "..foo bar!!???"
	trimExpectS = "foo bar"
	trimCutS    = "?!."
)

func BenchmarkTrimStr_Bytealg(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := bytealg.TrimStr(trimOriginS, trimCutS)
		if r != trimExpectS {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
		}
	}
}

func BenchmarkTrimStr_Native(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.Trim(trimOriginS, trimCutS)
		if r != trimExpectS {
			b.Errorf(`Trim: mismatch result %s and expectation %s`, r, trimExpectS)
		}
	}
}
