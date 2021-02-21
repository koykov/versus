package cbytebuf

// Benchmarks of valyala's bytebufferpool.ByteBuffer.

import (
	"bytes"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func BenchmarkByteBufferPool_Write(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
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
}

func BenchmarkByteBufferPool_WriteLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := bytebufferpool.Get()
		for i := 0; i < 1000; i++ {
			_, _ = buf.Write(source)
		}
		if !bytes.Equal(buf.Bytes(), expectedLong) {
			b.Error("not equal")
		}
		bytebufferpool.Put(buf)
	}
}
