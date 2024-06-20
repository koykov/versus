package urlvector

import (
	"testing"

	"github.com/koykov/byteconv"
	"github.com/valyala/fasthttp"
)

func TestParseFastHTTP(t *testing.T) {
	r := fasthttp.Request{}
	r.URI()
}

func BenchmarkParseFastHTTP(b *testing.B) {
	b.Run("small", func(b *testing.B) { benchFastHTTP(b, stages[0]) })
	b.Run("medium", func(b *testing.B) { benchFastHTTP(b, stages[1]) })
	b.Run("large", func(b *testing.B) { benchFastHTTP(b, stages[2]) })
	b.Run("big", func(b *testing.B) { benchFastHTTP(b, stages[3]) })
	b.Run("giant", func(b *testing.B) { benchFastHTTP(b, stages[4]) })
}

func benchFastHTTP(b *testing.B, s string) {
	host := []byte("localhost")
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u := fasthttp.AcquireURI()
			_ = u.Parse(host, byteconv.S2B(s))
			_ = u.QueryArgs()
			fasthttp.ReleaseURI(u)
		}
	})
}
