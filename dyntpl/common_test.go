package dyntpl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"log"

	"github.com/valyala/quicktemplate"
)

func init() {
	// Make sure both encoding/json and templates generate identical output.
	d := newTemplatesDataQT(3)
	bb := quicktemplate.AcquireByteBuffer()

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

	quicktemplate.ReleaseByteBuffer(bb)
}
