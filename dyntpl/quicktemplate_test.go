package dyntpl

import (
	"fmt"
	"testing"

	qt "github.com/valyala/quicktemplate"
	qtt "github.com/valyala/quicktemplate/testdata/templates"
)

func BenchmarkMarshalJSONQuickTemplate(b *testing.B) {
	bench := func(b *testing.B, n int) {
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
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
	b.Run("1000", func(b *testing.B) { bench(b, 1000) })
}

func BenchmarkMarshalXMLQuickTemplate(b *testing.B) {
	bench := func(b *testing.B, n int) {
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
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
	b.Run("1000", func(b *testing.B) { bench(b, 1000) })
}

func BenchmarkQuickTemplate(b *testing.B) {
	bench := func(b *testing.B, n int) {
		b.ReportAllocs()
		rows := getBenchRowsQT(n)
		b.RunParallel(func(pb *testing.PB) {
			bb := qt.AcquireByteBuffer()
			for pb.Next() {
				qtt.WriteBenchPage(bb, rows)
				bb.Reset()
			}
			qt.ReleaseByteBuffer(bb)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
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
