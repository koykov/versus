package urlvector

import (
	"net/url"
	"testing"
)

func BenchmarkParseNetUrl(b *testing.B) {
	b.Run("small", func(b *testing.B) { benchNetUrl(b, stages[0]) })
	b.Run("medium", func(b *testing.B) { benchNetUrl(b, stages[1]) })
	b.Run("large", func(b *testing.B) { benchNetUrl(b, stages[2]) })
	b.Run("big", func(b *testing.B) { benchNetUrl(b, stages[3]) })
	b.Run("giant", func(b *testing.B) { benchNetUrl(b, stages[4]) })
}

func benchNetUrl(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u, _ := url.Parse(s)
			_ = u.Query()
		}
	})
}

// func BenchmarkNetUrl(b *testing.B) {
// 	b.Run("parse url", func(b *testing.B) {
// 		b.ReportAllocs()
// 		b.SetBytes(int64(len(url0)))
// 		for i := 0; i < b.N; i++ {
// 			u, err := url.Parse(byteconv.B2S(url0))
// 			if err != nil {
// 				b.Error(err)
// 			}
//
// 			if u.Scheme != url0scheme {
// 				reportErr(b, "scheme", url0scheme, u.Scheme)
// 			}
// 			if u.User.Username() != url0user {
// 				reportErr(b, "username", url0user, u.User.Username())
// 			}
// 			if p, _ := u.User.Password(); p != url0pass {
// 				reportErr(b, "password", url0pass, p)
// 			}
// 			if u.Host != url0host {
// 				reportErr(b, "host", url0host, u.Host)
// 			}
// 			if u.Path != url0path {
// 				reportErr(b, "path", url0path, u.Path)
// 			}
// 			if u.Fragment != url0fragment {
// 				reportErr(b, "hash", url0fragment, u.Fragment)
// 			}
// 		}
// 	})
// 	b.Run("parse query", func(b *testing.B) {
// 		b.ReportAllocs()
// 		b.SetBytes(int64(len(url0)))
// 		for i := 0; i < b.N; i++ {
// 			u, err := url.Parse(byteconv.B2S(url0))
// 			if err != nil {
// 				b.Error(err)
// 			}
//
// 			q, err := url.ParseQuery(u.RawQuery)
// 			if err != nil {
// 				b.Error(err)
// 			}
// 			if q["page"][0] != url0qPage {
// 				b.Error("query[page] mismatch")
// 			}
// 			if q["ua"][0] != url0qUa {
// 				b.Error("query[ua] mismatch")
// 			}
// 			if q["uid"][0] != url0qUid {
// 				b.Error("query[uid] mismatch")
// 			}
// 			if q["ip"][0] != url0qIp {
// 				b.Error("query[ip] mismatch")
// 			}
// 		}
// 	})
// 	b.Run("mod host", func(b *testing.B) {
// 		b.ReportAllocs()
// 		b.SetBytes(int64(len(url0)))
// 		for i := 0; i < b.N; i++ {
// 			u, err := url.Parse(byteconv.B2S(url0))
// 			if err != nil {
// 				b.Error(err)
// 			}
//
// 			u.Host = "x.com:3333"
// 			if u.String() != url0modHost {
// 				b.Error("mod host mismatch")
// 			}
// 		}
// 	})
// 	b.Run("mod query", func(b *testing.B) {
// 		b.ReportAllocs()
// 		b.SetBytes(int64(len(url0)))
// 		for i := 0; i < b.N; i++ {
// 			u, err := url.Parse(byteconv.B2S(url0))
// 			if err != nil {
// 				b.Error(err)
// 			}
//
// 			q, err := url.ParseQuery(u.RawQuery)
// 			if err != nil {
// 				b.Error(err)
// 			}
// 			if q["ip"][0] != url0qIp {
// 				b.Error("mod query old ip mismatch")
// 			}
//
// 			q, err = url.ParseQuery(url0newQuery)
// 			if q["ip"][0] != url0qIp1 {
// 				b.Error("mod query new ip mismatch")
// 			}
// 		}
// 	})
// }
