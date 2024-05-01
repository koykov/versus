package dyntpl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"sync"
	"testing"
)

var nativePool = sync.Pool{New: func() any { return &bytes.Buffer{} }}

func BenchmarkNative(b *testing.B) {
	benchJSON := func(b *testing.B, n int) {
		b.ReportAllocs()
		d := newTemplatesDataQT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := nativePool.Get().(*bytes.Buffer)
				e := json.NewEncoder(buf)
				if err := e.Encode(d); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				buf.Reset()
				nativePool.Put(buf)
			}
		})
	}
	b.Run("json/1", func(b *testing.B) { benchJSON(b, 1) })
	b.Run("json/10", func(b *testing.B) { benchJSON(b, 10) })
	b.Run("json/100", func(b *testing.B) { benchJSON(b, 100) })
	b.Run("json/1000", func(b *testing.B) { benchJSON(b, 1000) })

	benchXML := func(b *testing.B, n int) {
		b.ReportAllocs()
		d := newTemplatesDataQT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := nativePool.Get().(*bytes.Buffer)
				e := xml.NewEncoder(buf)
				if err := e.Encode(d); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				buf.Reset()
				nativePool.Put(buf)
			}
		})
	}
	b.Run("xml/1", func(b *testing.B) { benchXML(b, 1) })
	b.Run("xml/10", func(b *testing.B) { benchXML(b, 10) })
	b.Run("xml/100", func(b *testing.B) { benchXML(b, 100) })
	b.Run("xml/1000", func(b *testing.B) { benchXML(b, 1000) })

	benchHTML := func(b *testing.B, n int) {
		b.ReportAllocs()
		rows := getBenchRowsQT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := nativePool.Get().(*bytes.Buffer)
				if err := tpl.Execute(buf, rows); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				buf.Reset()
				nativePool.Put(buf)
			}
		})
	}
	b.Run("html/1", func(b *testing.B) { benchHTML(b, 1) })
	b.Run("html/10", func(b *testing.B) { benchHTML(b, 10) })
	b.Run("html/100", func(b *testing.B) { benchHTML(b, 100) })
	b.Run("html/1000", func(b *testing.B) { benchHTML(b, 1000) })

	benchText := func(b *testing.B, n int) {
		b.ReportAllocs()
		rows := getBenchRowsQT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := nativePool.Get().(*bytes.Buffer)
				if err := textTpl.Execute(buf, rows); err != nil {
					b.Fatalf("unexpected error: %s", err)
				}
				buf.Reset()
				nativePool.Put(buf)
			}
		})
	}
	b.Run("text/1", func(b *testing.B) { benchText(b, 1) })
	b.Run("text/10", func(b *testing.B) { benchText(b, 10) })
	b.Run("text/100", func(b *testing.B) { benchText(b, 100) })
	b.Run("text/1000", func(b *testing.B) { benchText(b, 1000) })
}
