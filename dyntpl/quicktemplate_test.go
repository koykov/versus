package dyntpl

import (
	"fmt"
	"testing"

	qt "github.com/valyala/quicktemplate"
	qtt "github.com/valyala/quicktemplate/testdata/templates"
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
		bb := qt.AcquireByteBuffer()
		for pb.Next() {
			d.WriteJSON(bb)
			bb.Reset()
		}
		qt.ReleaseByteBuffer(bb)
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
		bb := qt.AcquireByteBuffer()
		for pb.Next() {
			d.WriteXML(bb)
			bb.Reset()
		}
		qt.ReleaseByteBuffer(bb)
	})
}

func newTemplatesDataQT(n int) *qtt.MarshalData {
	var rows []qtt.MarshalRow
	for i := 0; i < n; i++ {
		rows = append(rows, qtt.MarshalRow{
			Msg: fmt.Sprintf("тест %d", i),
			N:   i,
		})
	}
	return &qtt.MarshalData{
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
		bb := qt.AcquireByteBuffer()
		for pb.Next() {
			qtt.WriteBenchPage(bb, rows)
			bb.Reset()
		}
		qt.ReleaseByteBuffer(bb)
	})
}

func getBenchRowsQT(n int) []qtt.BenchRow {
	rows := make([]qtt.BenchRow, n)
	for i := 0; i < n; i++ {
		rows[i] = qtt.BenchRow{
			ID:      i,
			Message: fmt.Sprintf("message %d", i),
			Print:   ((i & 1) == 0),
		}
	}
	return rows
}
