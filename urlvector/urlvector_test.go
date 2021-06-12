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

		if vec.SchemeString() != "http" {
			reportErr(b, "scheme", "http", vec.SchemeString())
		}
		if vec.UsernameString() != "user" {
			reportErr(b, "username", "user", vec.UsernameString())
		}
		if vec.PasswordString() != "pass" {
			reportErr(b, "password", "pass", vec.PasswordString())
		}
		if vec.HostString() != "[1080:0:0:0:8:800:200C:417A]:61616" {
			reportErr(b, "host", "[1080:0:0:0:8:800:200C:417A]:61616", vec.HostString())
		}
		if vec.PathString() != "/foo" {
			reportErr(b, "path", "/foo", vec.PathString())
		}
		if vec.HashString() != "#foobar" {
			reportErr(b, "hash", "#foobar", vec.HashString())
		}
		if vec.Query().Get("page").String() != "https://ultra-software-base.ru/system/google-chrome.html?yclid=212247430717539672" {
			b.Error("query[page] mismatch")
		}
		if vec.Query().Get("ua").String() != "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 YaBrowser/21.3.3.230 Yowser/2.5 Safari/537.36" {
			b.Error("query[ua] mismatch")
		}

		urlvector.Release(vec)
	}
}
