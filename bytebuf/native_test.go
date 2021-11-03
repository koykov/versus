package bytebuf

import (
	"bytes"
	"strconv"
	"testing"
)

func BenchmarkNative(b *testing.B) {
	b.Run("writeX", func(b *testing.B) {
		b.ReportAllocs()
		var buf []byte
		for i := 0; i < b.N; i++ {
			buf = buf[:0]
			buf = append(buf, cbBytes...)
			buf = append(buf, '-')
			buf = append(buf, cbString...)
			buf = append(buf, '-')
			buf = append(buf, strconv.Itoa(int(cbInt))...)
			buf = append(buf, '-')
			buf = append(buf, strconv.Itoa(int(cbUint))...)
			buf = append(buf, '-')
			buf = append(buf, strconv.FormatFloat(cbFloat, 'f', -1, 64)...)

			if !bytes.Equal(buf, cbExpect) {
				b.Error("ByteArray: mismatch result and expectation")
			}
		}
	})
}
