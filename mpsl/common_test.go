package mpsl

type stage struct {
	hostname,
	tld, etld, etld1 string
	icann bool
}

var stages = []stage{
	{hostname: "google.org.ac", tld: "ac", etld: "org.ac", etld1: "google.org.ac", icann: true},
	{hostname: "github.ae", tld: "ae", etld: "", etld1: "github.ae", icann: true},
	{hostname: "unknown.no-tld", tld: "", etld: "", etld1: "unknown.no-tld", icann: false},
	{hostname: "go.dev", tld: "dev", etld: "", etld1: "go.dev", icann: true},
	{hostname: "verylongverylongverylongverylongverylongverylonghostname.ipa.xyz", tld: "xyz", etld: "", etld1: "ipa.xyz", icann: true},
	{hostname: "example.com", tld: "com", etld: "", etld1: "example.com", icann: true},
	{hostname: "example.id.au", tld: "au", etld: "id.au", etld1: "example.id.au", icann: true},
	{hostname: "www.ck", tld: "ck", etld: "", etld1: "www.ck", icann: true},
	{hostname: "foo.bar.xn--55qx5d.cn", tld: "cn", etld: "xn--55qx5d.cn", etld1: "bar.xn--55qx5d.cn", icann: true},
	{hostname: "a.b.c.minami.fukuoka.jp", tld: "jp", etld: "minami.fukuoka.jp", etld1: "c.minami.fukuoka.jp", icann: true},
	{hostname: "posts-and-telecommunications.museum", tld: "museum", etld: "", etld1: "posts-and-telecommunications.museum", icann: true},
	{hostname: "www.example.pvt.k12.ma.us", tld: "us", etld: "pvt.k12.ma.us", etld1: "example.pvt.k12.ma.us", icann: true},
	{hostname: "many.lol", tld: "lol", etld: "", etld1: "many.lol", icann: true},
	{hostname: "the.russian.for.moscow.is.xn--80adxhks", tld: "xn--80adxhks", etld: "", etld1: "is.xn--80adxhks", icann: true},
	{hostname: "blah.blah.s3-us-west-1.amazonaws.com", tld: "com", etld: "s3-us-west-1.amazonaws.com", etld1: "blah.s3-us-west-1.amazonaws.com", icann: false},
	{hostname: "thing.dyndns.org", tld: "org", etld: "dyndns.org", etld1: "thing.dyndns.org", icann: false},
	{hostname: "nosuchtld", tld: "", etld: "", etld1: ""},
}
