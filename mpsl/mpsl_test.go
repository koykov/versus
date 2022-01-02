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
	b.ResetTimer()
	for _, s := range stages {
		b.Run(s.hostname, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				etld1 := psdb.GetEffectiveTLDPlusOneStr(s.hostname)
				if etld1 != s.etld1 {
					b.Errorf("etld+1 mismatch: need '%s', got '%s'", s.etld1, etld1)
				}
			}
		})
	}
}
