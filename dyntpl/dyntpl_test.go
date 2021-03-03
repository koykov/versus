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

func benchmarkMarshalJSONDyntpl(b *testing.B, n int) {
	treeJSON, _ := dyntpl.Parse(tplMarshalJSON, false)
	treeXML, _ := dyntpl.Parse(tplMarshalXML, false)
	dyntpl.RegisterTpl("tplMarshalJSON", treeJSON)
	dyntpl.RegisterTpl("tplMarshalXML", treeXML)

	b.ResetTimer()
	b.ReportAllocs()

	d := newTemplatesDataDT(n)
	b.RunParallel(func(pb *testing.PB) {
		buf := cbytebuf.LBAcquire()
		ctx := dyntpl.AcquireCtx()
		ctx.Set("d", d, &testobj_ins.MarshalDataInspector{})
		for pb.Next() {
			_ = dyntpl.RenderTo(buf, "tplMarshalJSON", ctx)
		}
		dyntpl.ReleaseCtx(ctx)
		cbytebuf.LBRelease(buf)
	})
}

func BenchmarkMarshalJSONDyntpl1(b *testing.B) {
	benchmarkMarshalJSONDyntpl(b, 1)
}

func BenchmarkMarshalJSONDyntpl10(b *testing.B) {
	benchmarkMarshalJSONDyntpl(b, 10)
}

func BenchmarkMarshalJSONDyntpl100(b *testing.B) {
	benchmarkMarshalJSONDyntpl(b, 100)
}

func BenchmarkMarshalJSONDyntpl1000(b *testing.B) {
	benchmarkMarshalJSONDyntpl(b, 1000)
}

func benchmarkMarshalXMLDyntpl(b *testing.B, n int) {
	b.ResetTimer()
	b.ReportAllocs()

	d := newTemplatesDataDT(n)
	b.RunParallel(func(pb *testing.PB) {
		buf := cbytebuf.LBAcquire()
		ctx := dyntpl.AcquireCtx()
		ctx.Set("d", d, &testobj_ins.MarshalDataInspector{})
		for pb.Next() {
			_ = dyntpl.RenderTo(buf, "tplMarshalXML", ctx)
		}
		dyntpl.ReleaseCtx(ctx)
		cbytebuf.LBRelease(buf)
	})
}

func BenchmarkMarshalXMLDyntpl1(b *testing.B) {
	benchmarkMarshalXMLDyntpl(b, 1)
}

func BenchmarkMarshalXMLDyntpl10(b *testing.B) {
	benchmarkMarshalXMLDyntpl(b, 10)
}

func BenchmarkMarshalXMLDyntpl100(b *testing.B) {
	benchmarkMarshalXMLDyntpl(b, 100)
}

func BenchmarkMarshalXMLDyntpl1000(b *testing.B) {
	benchmarkMarshalXMLDyntpl(b, 1000)
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

func BenchmarkDyntpl1(b *testing.B) {
	benchmarkDyntpl(b, 1)
}

func BenchmarkDyntpl10(b *testing.B) {
	benchmarkDyntpl(b, 10)
}

func BenchmarkDyntpl100(b *testing.B) {
	benchmarkDyntpl(b, 100)
}

func benchmarkDyntpl(b *testing.B, rowsCount int) {
	tree, _ := dyntpl.Parse(tplTemplate, false)
	dyntpl.RegisterTpl("tplTemplate", tree)

	b.ResetTimer()
	b.ReportAllocs()

	bench := getBenchRowsDT(rowsCount)
	b.RunParallel(func(pb *testing.PB) {
		buf := cbytebuf.LBAcquire()
		ctx := dyntpl.AcquireCtx()
		ctx.Set("bench", bench, &testobj_ins.BenchRowsInspector{})
		for pb.Next() {
			_ = dyntpl.RenderTo(buf, "tplTemplate", ctx)
		}
		dyntpl.ReleaseCtx(ctx)
		cbytebuf.LBRelease(buf)
	})
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
