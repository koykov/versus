package urlvector

import (
	"testing"

	"github.com/koykov/urlvector"
)

func BenchmarkParseUrlvector(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		vec := urlvector.Acquire()
		err := vec.ParseCopy(url0)
		if err != nil {
			b.Error(err)
		}

		if vec.SchemeString() != url0scheme {
			reportErr(b, "scheme", url0scheme, vec.SchemeString())
		}
		if vec.UsernameString() != url0user {
			reportErr(b, "username", url0user, vec.UsernameString())
		}
		if vec.PasswordString() != url0pass {
			reportErr(b, "password", url0pass, vec.PasswordString())
		}
		if vec.HostString() != url0host {
			reportErr(b, "host", url0host, vec.HostString())
		}
		if vec.PathString() != url0path {
			reportErr(b, "path", url0path, vec.PathString())
		}
		if vec.HashString() != url0hash {
			reportErr(b, "hash", url0hash, vec.HashString())
		}
		if vec.Query().Get("page").String() != url0qPage {
			b.Error("query[page] mismatch")
		}
		if vec.Query().Get("ua").String() != url0qUa {
			b.Error("query[ua] mismatch")
		}

		urlvector.Release(vec)
	}
}
