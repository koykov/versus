package batch_replace

import (
	"bytes"
	"testing"

	"github.com/koykov/batch_replace"
)

func BenchmarkBatchReplace(b *testing.B) {
	b.Run("b2x", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := batch_replace.AcquireWithSource(breplOrigin)
			n := r.BytesToBytes(brTag0, brTag0Val).
				BytesToBytes(brTag1, brTag1Val).
				BytesToFloat(brTag2, 1234567.0987654321).
				BytesToInt(brTag3, int64(4)).
				Commit()
			if !bytes.Equal(n, breplExpect) {
				b.Error("BatchReplace: mismatch result and expectation")
			}
			batch_replace.Release(r)
		}
	})
	b.Run("s2x", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := batch_replace.AcquireWithSource(breplOriginS)
			n := r.StrToStr(brTag0S, brTag0ValS).
				StrToStr(brTag1S, brTag1ValS).
				StrToFloat(brTag2S, 1234567.0987654321).
				StrToInt(brTag3S, int64(4)).
				CommitStr()
			if n != breplExpectS {
				b.Error("BatchReplace: mismatch string result and expectation")
			}
			batch_replace.Release(r)
		}
	})
}
