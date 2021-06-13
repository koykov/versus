package urlvector

import (
	"testing"

	"github.com/koykov/urlvector"
)

func BenchmarkUrlvector_ParseUrl(b *testing.B) {
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

		urlvector.Release(vec)
	}
}

func BenchmarkUrlvector_ParseQuery(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		vec := urlvector.Acquire()
		err := vec.ParseCopy(url0)
		if err != nil {
			b.Error(err)
		}

		query := vec.Query()
		if query.Get("page").String() != url0qPage {
			b.Error("query[page] mismatch")
		}
		if query.Get("ua").String() != url0qUa {
			b.Error("query[ua] mismatch")
		}
		if query.Get("uid").String() != url0qUid {
			b.Error("query[uid] mismatch")
		}
		if query.Get("ip").String() != url0qIp {
			b.Error("query[page] mismatch")
		}

		urlvector.Release(vec)
	}
}

func BenchmarkUrlvector_ModHost(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		vec := urlvector.Acquire()
		err := vec.ParseCopy(url0)
		if err != nil {
			b.Error(err)
		}

		vec.SetHostnameString("x.com").SetPort(3333)
		if vec.String() != url0modHost {
			b.Error("mod host mismatch")
		}

		urlvector.Release(vec)
	}
}

func BenchmarkUrlvector_ModQuery(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		vec := urlvector.Acquire()
		err := vec.ParseCopy(url0)
		if err != nil {
			b.Error(err)
		}

		if vec.Query().Get("ip").String() != url0qIp {
			b.Error("mod query old ip mismatch")
		}

		vec.SetQueryString(url0newQuery)
		if vec.Query().Get("ip").String() != url0qIp1 {
			b.Error("mod query new ip mismatch")
		}

		urlvector.Release(vec)
	}
}
