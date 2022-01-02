package mpsl

import (
	"testing"

	"github.com/koykov/hash/fnv"
	"github.com/koykov/mpsl"
)

func BenchmarkMPSL(b *testing.B) {
	var (
		psdb *mpsl.DB
		err  error
	)
	if psdb, err = mpsl.New(fnv.BHasher{}); err != nil {
		b.Error(err)
	}
	if err = psdb.Load("testdata/full.psdb"); err != nil {
		b.Error(err)
		return
	}

	b.ReportAllocs()
	for _, s := range stages {
		b.Run(s.hostname, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tld, etld, etld1, icann := psdb.ParseStr(s.hostname)
				if tld != s.tld {
					b.Errorf("tld mismatch: need '%s', got '%s'", s.tld, tld)
				}
				if etld != s.etld {
					b.Errorf("etld mismatch: need '%s', got '%s'", s.etld, etld)
				}
				if etld1 != s.etld1 {
					b.Errorf("etld+1 mismatch: need '%s', got '%s'", s.etld1, etld1)
				}
				if icann != s.icann {
					b.Errorf("icann mismatch: need '%t', got '%t'", s.icann, icann)
				}
			}
		})
	}
}
