package cbytebuf

// Benchmarks of writing (append) data to byte array ([]byte).

import (
	"bytes"
	"testing"
)

func BenchmarkByteArray_Append(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0)
		for _, part := range parts {
			buf = append(buf, part...)
			buf = append(buf, ' ')
		}
		if !bytes.Equal(buf, expected) {
			b.Error("not equal")
		}
	}
}

func BenchmarkByteArray_AppendLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0)
		for i := 0; i < 1000; i++ {
			buf = append(buf, source...)
		}
		if !bytes.Equal(buf, expectedLong) {
			b.Error("not equal")
		}
	}
}
