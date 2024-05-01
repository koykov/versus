package dyntpl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"log"
	tt "text/template"

	qt "github.com/valyala/quicktemplate"
	qtt "github.com/valyala/quicktemplate/testdata/templates"
)

var (
	tpl     = template.Must(template.ParseFiles("tpl/bench.tpl"))
	textTpl = tt.Must(tt.ParseFiles("tpl/bench.tpl"))
)

func init() {
	// Make sure both encoding/json and templates generate identical output.
	d := newTemplatesDataQT(3)
	bb := qt.AcquireByteBuffer()
	defer qt.ReleaseByteBuffer(bb)

	expectedData, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	e := json.NewEncoder(bb)
	if err := e.Encode(d); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}
	bb.B = bytes.TrimSpace(bb.B)
	if !bytes.Equal(bb.B, expectedData) {
		log.Fatalf("unexpected data generated with encoding/json:\n%q\n. Expecting\n%q\n", bb.B, expectedData)
	}

	bb.Reset()
	d.WriteJSON(bb)
	if !bytes.Equal(bb.B, expectedData) {
		log.Fatalf("unexpected data generated with quicktemplate:\n%q\n. Expecting\n%q\n", bb.B, expectedData)
	}

	// make sure both encoding/xml and templates generate identical output.
	expectedData, err = xml.Marshal(d)
	if err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	bb.Reset()
	xe := xml.NewEncoder(bb)
	if err := xe.Encode(d); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}
	if !bytes.Equal(bb.B, expectedData) {
		log.Fatalf("unexpected data generated with encoding/xml:\n%q\n. Expecting\n%q\n", bb.B, expectedData)
	}

	bb.Reset()
	d.WriteXML(bb)
	if !bytes.Equal(bb.B, expectedData) {
		log.Fatalf("unexpected data generated with quicktemplate:\n%q\n. Expecting\n%q\n", bb.B, expectedData)
	}

	// make sure that both template engines generate the same result
	rows := getBenchRowsQT(3)

	bb1 := qt.AcquireByteBuffer()
	defer qt.ReleaseByteBuffer(bb1)
	if err := tpl.Execute(bb1, rows); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	bb2 := qt.AcquireByteBuffer()
	defer qt.ReleaseByteBuffer(bb2)
	qtt.WriteBenchPage(bb2, rows)

	if !bytes.Equal(bb1.B, bb2.B) {
		log.Fatalf("results mismatch:\n%q\n%q", bb1, bb2)
	}
}
