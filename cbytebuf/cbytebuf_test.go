package cbytebuf

// Benchmarks of cbytebuf.ByteBuf.

import (
	"bytes"
	"testing"

	"github.com/koykov/cbytebuf"
)

func BenchmarkCByteBuf_Write(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := cbytebuf.Acquire()
		for _, part := range parts {
			_, _ = buf.Write(part)
			_, _ = buf.Write(space)
		}
		if !bytes.Equal(buf.Bytes(), expected) {
			b.Error("not equal")
		}
		cbytebuf.Release(buf)
	}
}

func BenchmarkCByteBuf_WriteLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := cbytebuf.Acquire()
		for i := 0; i < 1000; i++ {
			_, _ = buf.Write(source)
		}
		if !bytes.Equal(buf.Bytes(), expectedLong) {
			b.Error("not equal")
		}
		cbytebuf.Release(buf)
	}
}
