package xmlvector

import (
	"testing"

	xmldom "github.com/subchen/go-xmldom"
)

func TestParseXmlDOM(t *testing.T) {
	doc := xmldom.Must(xmldom.ParseXML(mediumFixture))
	root := doc.Root
	_ = root
	if root.Name != "breakfast_menu" {
		t.FailNow()
	}
}

func BenchmarkParseXmlDOM(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchXmlDOM(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchXmlDOM(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchXmlDOM(b, largeFixture)
	})
	// b.Run("sigmod", func(b *testing.B) {
	// 	benchXmlDOM(b, sigmodFixture)
	// })
	b.Run("mondial", func(b *testing.B) {
		benchXmlDOM(b, mondialFixture)
	})
}

func benchXmlDOM(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doc, err := xmldom.ParseXML(s)
			if err != nil {
				b.Error()
				break
			}
			root := doc.Root
			_ = root
		}
	})
}
