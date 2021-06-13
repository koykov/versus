package urlvector

import "testing"

var (
	url0         = []byte("http://user:pass@[1080:0:0:0:8:800:200C:417A]:61616/foo?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=5.5.5.5&subage=102&language=ru#foobar")
	url0scheme   = "http"
	url0user     = "user"
	url0pass     = "pass"
	url0host     = "[1080:0:0:0:8:800:200C:417A]:61616"
	url0path     = "/foo"
	url0hash     = "#foobar"
	url0fragment = "foobar"
	url0qPage    = "https://ultra-software-base.ru/system/google-chrome.html?yclid=212247430717539672"
	url0qUa      = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 YaBrowser/21.3.3.230 Yowser/2.5 Safari/537.36"
	url0qUid     = "4f5d0edc-3a3e-48d0-9872-0b48a7998ac6"
	url0qIp      = "5.5.5.5"
	url0qIp1     = "127.0.0.1"

	url0modHost  = "http://user:pass@x.com:3333/foo?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=5.5.5.5&subage=102&language=ru#foobar"
	url0newQuery = "?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=127.0.0.1&subage=102&language=ru"
	url0modQuery = "http://user:pass@[1080:0:0:0:8:800:200C:417A]:61616/foo?v=default&blockID=319385&page=https%3A%2F%2Fultra-software-base.ru%2Fsystem%2Fgoogle-chrome.html%3Fyclid%3D212247430717539672&domain=ultra-software-base.ru&uid=4f5d0edc-3a3e-48d0-9872-0b48a7998ac6&clientNotice=true&imgX=360&imgY=240&limit=1&subage_dt=2021-01-29&format=json&cur=RUB&ua=Mozilla%2F5.0+%28Windows+NT+6.1%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F89.0.4389.105+YaBrowser%2F21.3.3.230+Yowser%2F2.5+Safari%2F537.36&ip=127.0.0.1&subage=102&language=ru#foobar"
)

func reportErr(x testing.TB, key, need, got string) {
	x.Error(key, "mismatch", "need", need, "got", got)
}
