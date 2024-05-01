package dyntpl

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/koykov/dyntpl"
	"github.com/koykov/dyntpl/testobj"
	"github.com/koykov/dyntpl/testobj_ins"
)

var (
	tplTemplate = []byte(`<html>
	<head><title>test</title></head>
	<body>
		<ul>
		{% for _, row := range d.Rows %}
			<li>ID={%= row.ID %}, Message={%= row.Message %}</li>
		{% endfor %}
		</ul>
	</body>
</html>`)
	tplMarshalJSON = []byte(`{
	"Foo": {%= d.Foo %},
	"Bar": "{%j= d.Bar %}",
	"Rows":[
		{% for _, r := range d.Rows separator , %}
			{
				"Msg": "{%j= r.Msg %}",
				"N": {%= r.N %}
			}
		{% endfor %}
	]
}`)
	tplMarshalXML = []byte(`<MarshalData>
	<Foo>{%= d.Foo %}</Foo>
	<Bar>{%= d.Bar %}</Bar>
	{% for _, r := range d.Rows %}
		<Rows>
			<Msg>{%= r.Msg %}</Msg>
			<N>{%= r.N %}</N>
		</Rows>
	{% endfor %}
</MarshalData>`)

	dtplPool = sync.Pool{New: func() any { return &bytes.Buffer{} }}
)

func init() {
	treeJSON, _ := dyntpl.Parse(tplMarshalJSON, false)
	dyntpl.RegisterTplKey("tplMarshalJSON", treeJSON)
	treeXML, _ := dyntpl.Parse(tplMarshalXML, false)
	dyntpl.RegisterTplKey("tplMarshalXML", treeXML)
	tree, _ := dyntpl.Parse(tplTemplate, false)
	dyntpl.RegisterTplKey("tplTemplate", tree)
}

func BenchmarkDyntpl(b *testing.B) {
	benchJSON := func(b *testing.B, n int) {
		b.ResetTimer()
		b.ReportAllocs()

		d := newTemplatesDataDT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := dtplPool.Get().(*bytes.Buffer)
				ctx := dyntpl.AcquireCtx()
				ctx.Set("d", d, testobj_ins.MarshalDataInspector{})
				_ = dyntpl.Write(buf, "tplMarshalJSON", ctx)
				dyntpl.ReleaseCtx(ctx)
				buf.Reset()
				dtplPool.Put(buf)
			}
		})
	}
	b.Run("json/1", func(b *testing.B) { benchJSON(b, 1) })
	b.Run("json/10", func(b *testing.B) { benchJSON(b, 10) })
	b.Run("json/100", func(b *testing.B) { benchJSON(b, 100) })
	b.Run("json/1000", func(b *testing.B) { benchJSON(b, 1000) })

	benchXML := func(b *testing.B, n int) {
		b.ResetTimer()
		b.ReportAllocs()

		d := newTemplatesDataDT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := dtplPool.Get().(*bytes.Buffer)
				ctx := dyntpl.AcquireCtx()
				ctx.Set("d", d, testobj_ins.MarshalDataInspector{})
				_ = dyntpl.Write(buf, "tplMarshalXML", ctx)
				dyntpl.ReleaseCtx(ctx)
				buf.Reset()
				dtplPool.Put(buf)
			}
		})
	}
	b.Run("xml/1", func(b *testing.B) { benchXML(b, 1) })
	b.Run("xml/10", func(b *testing.B) { benchXML(b, 10) })
	b.Run("xml/100", func(b *testing.B) { benchXML(b, 100) })
	b.Run("xml/1000", func(b *testing.B) { benchXML(b, 1000) })

	benchText := func(b *testing.B, n int) {
		b.ResetTimer()
		b.ReportAllocs()

		bench := getBenchRowsDT(n)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				buf := dtplPool.Get().(*bytes.Buffer)
				ctx := dyntpl.AcquireCtx()
				ctx.Set("d", bench, testobj_ins.BenchRowsInspector{})
				_ = dyntpl.Write(buf, "tplTemplate", ctx)
				dyntpl.ReleaseCtx(ctx)
				buf.Reset()
				dtplPool.Put(buf)
			}
		})
	}
	b.Run("text/1", func(b *testing.B) { benchText(b, 1) })
	b.Run("text/10", func(b *testing.B) { benchText(b, 10) })
	b.Run("text/100", func(b *testing.B) { benchText(b, 100) })
	b.Run("text/1000", func(b *testing.B) { benchText(b, 1000) })
}

func newTemplatesDataDT(n int) *testobj.MarshalData {
	var rows []testobj.MarshalRow
	for i := 0; i < n; i++ {
		rows = append(rows, testobj.MarshalRow{
			Msg: fmt.Sprintf("тест %d", i),
			N:   i,
		})
	}
	return &testobj.MarshalData{
		Foo:  1,
		Bar:  "foobar",
		Rows: rows,
	}
}

func getBenchRowsDT(n int) *testobj.BenchRows {
	bench := &testobj.BenchRows{
		Rows: make([]testobj.BenchRow, n),
	}
	for i := 0; i < n; i++ {
		bench.Rows[i] = testobj.BenchRow{
			ID:      i,
			Message: fmt.Sprintf("message %d", i),
			Print:   (i & 1) == 0,
		}
	}
	return bench
}
