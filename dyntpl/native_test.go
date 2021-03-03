package dyntpl

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/valyala/quicktemplate"
)

func BenchmarkMarshalJSONNative1(b *testing.B) {
	benchmarkMarshalJSONNative(b, 1)
}

func BenchmarkMarshalJSONNative10(b *testing.B) {
	benchmarkMarshalJSONNative(b, 10)
}

func BenchmarkMarshalJSONNative100(b *testing.B) {
	benchmarkMarshalJSONNative(b, 100)
}

func BenchmarkMarshalJSONNative1000(b *testing.B) {
	benchmarkMarshalJSONNative(b, 1000)
}

func benchmarkMarshalJSONNative(b *testing.B, n int) {
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

func BenchmarkMarshalXMLNative1(b *testing.B) {
	benchmarkMarshalXMLNative(b, 1)
}

func BenchmarkMarshalXMLNative10(b *testing.B) {
	benchmarkMarshalXMLNative(b, 10)
}

func BenchmarkMarshalXMLNative100(b *testing.B) {
	benchmarkMarshalXMLNative(b, 100)
}

func BenchmarkMarshalXMLNative1000(b *testing.B) {
	benchmarkMarshalXMLNative(b, 1000)
}

func benchmarkMarshalXMLNative(b *testing.B, n int) {
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

func BenchmarkHTMLTemplate1(b *testing.B) {
	benchmarkHTMLTemplate(b, 1)
}

func BenchmarkHTMLTemplate10(b *testing.B) {
	benchmarkHTMLTemplate(b, 10)
}

func BenchmarkHTMLTemplate100(b *testing.B) {
	benchmarkHTMLTemplate(b, 100)
}

func benchmarkHTMLTemplate(b *testing.B, rowsCount int) {
	b.ReportAllocs()
	rows := getBenchRowsQT(rowsCount)
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

func BenchmarkTextTemplate1(b *testing.B) {
	benchmarkTextTemplate(b, 1)
}

func BenchmarkTextTemplate10(b *testing.B) {
	benchmarkTextTemplate(b, 10)
}

func BenchmarkTextTemplate100(b *testing.B) {
	benchmarkTextTemplate(b, 100)
}

func benchmarkTextTemplate(b *testing.B, rowsCount int) {
	b.ReportAllocs()
	rows := getBenchRowsQT(rowsCount)
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
