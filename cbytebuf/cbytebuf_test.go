package cbytebuf

// Benchmarks of cbytebuf.ByteBuf.

import (
	"bytes"
	"testing"

	"github.com/koykov/cbytebuf"
)

func BenchmarkCByteBuf(b *testing.B) {
	b.Run("write", func(b *testing.B) {
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			buf := cbytebuf.LBAcquire()
			for _, part := range parts {
				_, _ = buf.Write(part)
				_, _ = buf.Write(space)
			}
			if !bytes.Equal(buf.Bytes(), expected) {
				b.Error("not equal")
			}
			cbytebuf.LBRelease(buf)
		}
	})
	b.Run("write long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := cbytebuf.LBAcquire()
			for j := 0; j < 1000; j++ {
				_, _ = buf.Write(source)
			}
			if !bytes.Equal(buf.Bytes(), expectedLong) {
				b.Error("not equal")
			}
			cbytebuf.LBRelease(buf)
		}
	})
}
