package batch_replace

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkNative(b *testing.B) {
	b.Run("b2x", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			n := bytes.Replace(breplOrigin, brTag0, brTag0Val, -1)
			n = bytes.Replace(n, brTag1, brTag1Val, -1)
			n = bytes.Replace(n, brTag2, []byte(strconv.FormatFloat(1234567.0987654321, 'f', -1, 64)), -1)
			n = bytes.Replace(n, brTag3, []byte(strconv.Itoa(4)), -1)
			if !bytes.Equal(n, breplExpect) {
				b.Error("BatchReplace: mismatch result and expectation")
			}
		}
	})
	b.Run("s2x", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			n := strings.Replace(breplOriginS, brTag0S, brTag0ValS, -1)
			n = strings.Replace(n, brTag1S, brTag1ValS, -1)
			n = strings.Replace(n, brTag2S, strconv.FormatFloat(1234567.0987654321, 'f', -1, 64), -1)
			n = strings.Replace(n, brTag3S, strconv.Itoa(4), -1)
			if n != breplExpectS {
				b.Error("BatchReplaceStr: mismatch result and expectation")
			}
		}
	})
}
