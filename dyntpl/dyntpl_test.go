package dyntpl

import (
	"fmt"
	"testing"

	"github.com/koykov/cbytebuf"

	"github.com/koykov/dyntpl"
	"github.com/koykov/dyntpl/testobj"
	"github.com/koykov/dyntpl/testobj_ins"
)

var (
	tplMarshalJSON = []byte(`{
	"Foo": {%= d.Foo %},
	"Bar": {%q= d.Bar %},
	"Rows":[
		{% for _, r := range d.Rows separator , %}
			{
				"Msg": {%q= r.Msg %},
				"N": {%= r.N %}
			}
		{% endfor %}
	]
}`)
	tplMarshalXML = []byte(`<MarshalData>
	<Foo>{%= d.Foo %}</Foo>
	<Bar>{%h= d.Bar %}</Bar>
	{% for _, r := range d.Rows %}
		<Rows>
			<Msg>{%h= r.Msg %}</Msg>
			<N>{%= r.N %}</N>
		</Rows>
	{% endfor %}
</MarshalData>`)
)

func BenchmarkMarshalJSONDyntpl(b *testing.B) {
	bench := func(b *testing.B, n int) {
		treeJSON, _ := dyntpl.Parse(tplMarshalJSON, false)
		treeXML, _ := dyntpl.Parse(tplMarshalXML, false)
		dyntpl.RegisterTpl(0, "tplMarshalJSON", treeJSON)
		dyntpl.RegisterTpl(0, "tplMarshalXML", treeXML)

		b.ResetTimer()
		b.ReportAllocs()

		d := newTemplatesDataDT(n)
		b.RunParallel(func(pb *testing.PB) {
			buf := cbytebuf.LBAcquire()
			ctx := dyntpl.AcquireCtx()
			ctx.Set("d", d, &testobj_ins.MarshalDataInspector{})
			for pb.Next() {
				_ = dyntpl.Write(buf, "tplMarshalJSON", ctx)
			}
			dyntpl.ReleaseCtx(ctx)
			cbytebuf.LBRelease(buf)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
	b.Run("1000", func(b *testing.B) { bench(b, 1000) })
}

func BenchmarkMarshalXMLDyntpl(b *testing.B) {
	bench := func(b *testing.B, n int) {
		b.ResetTimer()
		b.ReportAllocs()

		d := newTemplatesDataDT(n)
		b.RunParallel(func(pb *testing.PB) {
			buf := cbytebuf.LBAcquire()
			ctx := dyntpl.AcquireCtx()
			ctx.Set("d", d, &testobj_ins.MarshalDataInspector{})
			for pb.Next() {
				_ = dyntpl.Write(buf, "tplMarshalXML", ctx)
			}
			dyntpl.ReleaseCtx(ctx)
			cbytebuf.LBRelease(buf)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
	b.Run("1000", func(b *testing.B) { bench(b, 1000) })
}

func BenchmarkDyntpl(b *testing.B) {
	bench := func(b *testing.B, n int) {
		tree, _ := dyntpl.Parse(tplTemplate, false)
		dyntpl.RegisterTpl(0, "tplTemplate", tree)

		b.ResetTimer()
		b.ReportAllocs()

		bench := getBenchRowsDT(n)
		b.RunParallel(func(pb *testing.PB) {
			buf := cbytebuf.LBAcquire()
			ctx := dyntpl.AcquireCtx()
			ctx.Set("bench", bench, &testobj_ins.BenchRowsInspector{})
			for pb.Next() {
				_ = dyntpl.Write(buf, "tplTemplate", ctx)
			}
			dyntpl.ReleaseCtx(ctx)
			cbytebuf.LBRelease(buf)
		})
	}
	b.Run("1", func(b *testing.B) { bench(b, 1) })
	b.Run("10", func(b *testing.B) { bench(b, 10) })
	b.Run("100", func(b *testing.B) { bench(b, 100) })
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
