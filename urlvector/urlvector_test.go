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
			reportErr(b, "scheme", "user", vec.UsernameString())
		}
		if vec.PasswordString() != "password" {
			reportErr(b, "scheme", "password", vec.PasswordString())
		}
		if vec.HostString() != "[1080:0:0:0:8:800:200C:417A]:61616" {
			reportErr(b, "scheme", "[1080:0:0:0:8:800:200C:417A]:61616", vec.HostString())
		}
		if vec.PathString() != "/foo" {
			reportErr(b, "scheme", "/foo", vec.PathString())
		}
		if vec.HashString() != "#foobar" {
			reportErr(b, "scheme", "#foobar", vec.HashString())
		}
		if vec.Query().Get("page").String() != "https://ultra-software-base.ru/system/google-chrome.html?yclid=212247430717539672" {
			b.Error("query[page] mismatch")
		}

		urlvector.Release(vec)
	}
}
