package mpsl

import (
	"testing"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

func BenchmarkPublicSuffix(b *testing.B) {
	var (
		list *publicsuffix.List
		err  error
	)
	if list, err = publicsuffix.NewListFromFile("testdata/full.psdb", nil); err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for _, s := range stages {
		b.Run(s.hostname, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				etld1, _ := publicsuffix.DomainFromListWithOptions(list, s.hostname, publicsuffix.DefaultFindOptions)
				if etld1 != s.etld1 {
					b.Errorf("etld+1 mismatch: need '%s', got '%s'", s.etld1, etld1)
				}
			}
		})
	}
}
