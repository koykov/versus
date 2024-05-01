package dyntpl

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	qtt "github.com/valyala/quicktemplate/testdata/templates"
)

var qtplPool = sync.Pool{New: func() any { return &bytes.Buffer{} }}

func BenchmarkQuickTemplate(b *testing.B) {
	benchJSON := func(b *testing.B, n int) {
		b.ReportAllocs()
		d := newTemplatesDataQT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := qtplPool.Get().(*bytes.Buffer)
				d.WriteJSON(buf)
				buf.Reset()
				qtplPool.Put(buf)
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
				buf := qtplPool.Get().(*bytes.Buffer)
				d.WriteXML(buf)
				buf.Reset()
				qtplPool.Put(buf)
			}
		})
	}
	b.Run("xml/1", func(b *testing.B) { benchXML(b, 1) })
	b.Run("xml/10", func(b *testing.B) { benchXML(b, 10) })
	b.Run("xml/100", func(b *testing.B) { benchXML(b, 100) })
	b.Run("xml/1000", func(b *testing.B) { benchXML(b, 1000) })

	benchText := func(b *testing.B, n int) {
		b.ReportAllocs()
		rows := getBenchRowsQT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := qtplPool.Get().(*bytes.Buffer)
				qtt.WriteBenchPage(buf, rows)
				buf.Reset()
				qtplPool.Put(buf)
			}
		})
	}
	b.Run("text/1", func(b *testing.B) { benchText(b, 1) })
	b.Run("text/10", func(b *testing.B) { benchText(b, 10) })
	b.Run("text/100", func(b *testing.B) { benchText(b, 100) })
	b.Run("text/1000", func(b *testing.B) { benchText(b, 1000) })
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
