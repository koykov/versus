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
