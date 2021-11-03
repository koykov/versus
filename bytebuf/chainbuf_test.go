package bytebuf

import (
	"bytes"
	"testing"

	"github.com/koykov/bytebuf"
)

func BenchmarkChainBuf(b *testing.B) {
	b.Run("writeX", func(b *testing.B) {
		b.ReportAllocs()
		var cb bytebuf.ChainBuf
		for i := 0; i < b.N; i++ {
			cb.Reset().
				Write(cbBytes).WriteByte('-').
				WriteStr(cbString).WriteByte('-').
				WriteInt(cbInt).WriteByte('-').
				WriteUint(cbUint).WriteByte('-').
				WriteFloat(cbFloat)

			if !bytes.Equal(cb.Bytes(), cbExpect) {
				b.Error("ChainBuf: mismatch result and expectation")
			}
		}
	})
}
