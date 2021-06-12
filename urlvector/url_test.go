package urlvector

import (
	"net/url"
	"testing"

	"github.com/koykov/fastconv"
)

func BenchmarkParseUrl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		u, err := url.Parse(fastconv.B2S(url0))
		if err != nil {
			b.Error(err)
		}

		if u.Scheme != "http" {
			reportErr(b, "scheme", "http", u.Scheme)
		}
		if u.User.Username() != "user" {
			reportErr(b, "username", "user", u.User.Username())
		}
		if p, _ := u.User.Password(); p != "pass" {
			reportErr(b, "password", "pass", p)
		}
		if u.Host != "[1080:0:0:0:8:800:200C:417A]:61616" {
			reportErr(b, "host", "[1080:0:0:0:8:800:200C:417A]:61616", u.Host)
		}
		if u.Path != "/foo" {
			reportErr(b, "path", "/foo", u.Path)
		}
		if u.Fragment != "foobar" {
			reportErr(b, "hash", "foobar", u.Fragment)
		}

		q, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			b.Error(err)
		}
		if q["page"][0] != "https://ultra-software-base.ru/system/google-chrome.html?yclid=212247430717539672" {
			b.Error("query[page] mismatch")
		}
		if q["ua"][0] != "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 YaBrowser/21.3.3.230 Yowser/2.5 Safari/537.36" {
			b.Error("query[ua] mismatch")
		}
	}
}
