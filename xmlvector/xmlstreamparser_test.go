package xmlvector

import (
	"bufio"
	"bytes"
	"os"
	"sync"
	"testing"

	xmlparser "github.com/tamerh/xml-stream-parser"
)

func TestParseXmlStreamParser(t *testing.T) {
	f, _ := os.Open("testdata/medium.xml")
	rdr := bufio.NewReaderSize(f, 65536)
	parser := xmlparser.NewXMLParser(rdr, "food")
	for node := range parser.Stream() {
		_ = node.Childs["name"][0].InnerText
	}
}

func BenchmarkParseXmlStreamParser(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchXmlStreamParser(b, smallFixture, "note")
	})
	b.Run("medium", func(b *testing.B) {
		benchXmlStreamParser(b, mediumFixture, "food")
	})
	b.Run("large", func(b *testing.B) {
		benchXmlStreamParser(b, largeFixture, "PLANT")
	})
	b.Run("sigmod", func(b *testing.B) {
		benchXmlStreamParser(b, sigmodFixture, "issue")
	})
	b.Run("mondial", func(b *testing.B) {
		benchXmlStreamParser(b, mondialFixture, "country")
	})
}

func benchXmlStreamParser(b *testing.B, s, root string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			xsp_ := xspPool.Get().(*xsp)
			xsp_.buf.Reset()
			xsp_.buf.WriteString(s)
			xsp_.rdr.Reset(xsp_.buf)
			parser := xmlparser.NewXMLParser(xsp_.rdr, root)
			for node := range parser.Stream() {
				_ = node.Childs
			}
			xspPool.Put(xsp_)
		}
	})
}

type xsp struct {
	buf *bytes.Buffer
	rdr *bufio.Reader
}

var xspPool = sync.Pool{
	New: func() any {
		xsp_ := xsp{
			buf: &bytes.Buffer{},
			rdr: &bufio.Reader{},
		}
		return &xsp_
	},
}
