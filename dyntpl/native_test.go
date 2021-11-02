package dyntpl

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/valyala/quicktemplate"
)

func BenchmarkMarshalJSONNative(b *testing.B) {
	bench := func(b *testing.B, n int) {
		b.ReportAllocs()
		d := newTemplatesDataQT(n)
		b.RunParallel(func(pb *testing.PB) {
			bb := quicktemplate.AcquireByteBuffer()
			e := json.NewEncoder(bb)
			for pb.Next() {
				if err := e.Encode(d); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				bb.Reset()
			}
			quicktemplate.ReleaseByteBuffer(bb)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
	b.Run("1000", func(b *testing.B) { bench(b, 1000) })
}

func BenchmarkMarshalXMLNative(b *testing.B) {
	bench := func(b *testing.B, n int) {
		b.ReportAllocs()
		d := newTemplatesDataQT(n)
		b.RunParallel(func(pb *testing.PB) {
			bb := quicktemplate.AcquireByteBuffer()
			e := xml.NewEncoder(bb)
			for pb.Next() {
				if err := e.Encode(d); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				bb.Reset()
			}
			quicktemplate.ReleaseByteBuffer(bb)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
	b.Run("1000", func(b *testing.B) { bench(b, 1000) })
}

func BenchmarkMarshalHTMLNative(b *testing.B) {
	bench := func(b *testing.B, n int) {
		b.ReportAllocs()
		rows := getBenchRowsQT(n)
		b.RunParallel(func(pb *testing.PB) {
			bb := quicktemplate.AcquireByteBuffer()
			for pb.Next() {
				if err := tpl.Execute(bb, rows); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				bb.Reset()
			}
			quicktemplate.ReleaseByteBuffer(bb)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
}

func BenchmarkMarshalTextNative(b *testing.B) {
	bench := func(b *testing.B, n int) {
		b.ReportAllocs()
		rows := getBenchRowsQT(n)
		b.RunParallel(func(pb *testing.PB) {
			bb := quicktemplate.AcquireByteBuffer()
			for pb.Next() {
				if err := textTpl.Execute(bb, rows); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				bb.Reset()
			}
			quicktemplate.ReleaseByteBuffer(bb)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
}
