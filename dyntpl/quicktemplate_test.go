package dyntpl

import (
	"fmt"
	"testing"

	"github.com/valyala/quicktemplate"
	"github.com/valyala/quicktemplate/testdata/templates"
)

func BenchmarkMarshalJSONQuickTemplate1(b *testing.B) {
	benchmarkMarshalJSONQuickTemplate(b, 1)
}

func BenchmarkMarshalJSONQuickTemplate10(b *testing.B) {
	benchmarkMarshalJSONQuickTemplate(b, 10)
}

func BenchmarkMarshalJSONQuickTemplate100(b *testing.B) {
	benchmarkMarshalJSONQuickTemplate(b, 100)
}

func BenchmarkMarshalJSONQuickTemplate1000(b *testing.B) {
	benchmarkMarshalJSONQuickTemplate(b, 1000)
}

func benchmarkMarshalJSONQuickTemplate(b *testing.B, n int) {
	b.ReportAllocs()
	d := newTemplatesDataQT(n)
	b.RunParallel(func(pb *testing.PB) {
		bb := quicktemplate.AcquireByteBuffer()
		for pb.Next() {
			d.WriteJSON(bb)
			bb.Reset()
		}
		quicktemplate.ReleaseByteBuffer(bb)
	})
}

func BenchmarkMarshalXMLQuickTemplate1(b *testing.B) {
	benchmarkMarshalXMLQuickTemplate(b, 1)
}

func BenchmarkMarshalXMLQuickTemplate10(b *testing.B) {
	benchmarkMarshalXMLQuickTemplate(b, 10)
}

func BenchmarkMarshalXMLQuickTemplate100(b *testing.B) {
	benchmarkMarshalXMLQuickTemplate(b, 100)
}

func BenchmarkMarshalXMLQuickTemplate1000(b *testing.B) {
	benchmarkMarshalXMLQuickTemplate(b, 1000)
}

func benchmarkMarshalXMLQuickTemplate(b *testing.B, n int) {
	b.ReportAllocs()
	d := newTemplatesDataQT(n)
	b.RunParallel(func(pb *testing.PB) {
		bb := quicktemplate.AcquireByteBuffer()
		for pb.Next() {
			d.WriteXML(bb)
			bb.Reset()
		}
		quicktemplate.ReleaseByteBuffer(bb)
	})
}

func newTemplatesDataQT(n int) *templates.MarshalData {
	var rows []templates.MarshalRow
	for i := 0; i < n; i++ {
		rows = append(rows, templates.MarshalRow{
			Msg: fmt.Sprintf("тест %d", i),
			N:   i,
		})
	}
	return &templates.MarshalData{
		Foo:  1,
		Bar:  "foobar",
		Rows: rows,
	}
}

func BenchmarkQuickTemplate1(b *testing.B) {
	benchmarkQuickTemplate(b, 1)
}

func BenchmarkQuickTemplate10(b *testing.B) {
	benchmarkQuickTemplate(b, 10)
}

func BenchmarkQuickTemplate100(b *testing.B) {
	benchmarkQuickTemplate(b, 100)
}

func benchmarkQuickTemplate(b *testing.B, rowsCount int) {
	b.ReportAllocs()
	rows := getBenchRowsQT(rowsCount)
	b.RunParallel(func(pb *testing.PB) {
		bb := quicktemplate.AcquireByteBuffer()
		for pb.Next() {
			templates.WriteBenchPage(bb, rows)
			bb.Reset()
		}
		quicktemplate.ReleaseByteBuffer(bb)
	})
}

func getBenchRowsQT(n int) []templates.BenchRow {
	rows := make([]templates.BenchRow, n)
	for i := 0; i < n; i++ {
		rows[i] = templates.BenchRow{
			ID:      i,
			Message: fmt.Sprintf("message %d", i),
			Print:   ((i & 1) == 0),
		}
	}
	return rows
}
