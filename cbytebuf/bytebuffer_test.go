package cbytebuf

// Benchmarks of native bytes.Buffer.

import (
	"bytes"
	"testing"
)

func BenchmarkByteBufferNative_Write(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, part := range parts {
			buf.Write(part)
			buf.WriteByte(' ')
		}
		if !bytes.Equal(buf.Bytes(), expected) {
			b.Error("not equal")
		}
		buf.Reset()
	}
}

func BenchmarkByteBufferNative_WriteLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for i := 0; i < 1000; i++ {
			buf.Write(source)
		}
		if !bytes.Equal(buf.Bytes(), expectedLong) {
			b.Error("not equal")
		}
		buf.Reset()
	}
}
