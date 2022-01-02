package mpsl

import (
	"testing"

	"golang.org/x/net/publicsuffix"
)

func BenchmarkXPublicSuffix(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for _, s := range stages {
		b.Run(s.hostname, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				etld1, _ := publicsuffix.EffectiveTLDPlusOne(s.hostname)
				if etld1 != s.etld1 {
					b.Errorf("etld+1 mismatch: need '%s', got '%s'", s.etld1, etld1)
				}
			}
		})
	}
}
