package cbytebuf

// Benchmarks of valyala's bytebufferpool.ByteBuffer.

import (
	"bytes"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func BenchmarkByteBufferPool(b *testing.B) {
	b.Run("write", func(b *testing.B) {
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
			buf := bytebufferpool.Get()
			for _, part := range parts {
				_, _ = buf.Write(part)
				_, _ = buf.Write(space)
			}
			if !bytes.Equal(buf.Bytes(), expected) {
				b.Error("not equal")
			}
			bytebufferpool.Put(buf)
		}
	})
	b.Run("write long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := bytebufferpool.Get()
			for j := 0; j < 1000; j++ {
				_, _ = buf.Write(source)
			}
			if !bytes.Equal(buf.Bytes(), expectedLong) {
				b.Error("not equal")
			}
			bytebufferpool.Put(buf)
		}
	})
}
