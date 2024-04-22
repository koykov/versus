package cbytebuf

// Benchmarks of native bytes.Buffer.

import (
	"bytes"
	"testing"
)

func BenchmarkByteBuffer(b *testing.B) {
	b.Run("write", func(b *testing.B) {
		b.ReportAllocs()
		for j := 0; j < b.N; j++ {
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
	})
	b.Run("write long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			for j := 0; j < 1000; j++ {
				buf.Write(source)
			}
			if !bytes.Equal(buf.Bytes(), expectedLong) {
				b.Error("not equal")
			}
			buf.Reset()
		}
	})
}
