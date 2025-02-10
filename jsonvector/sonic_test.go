package jsonvector

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic/ast"
)

func BenchmarkSonic(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		benchSonic(b, smallFixture)
	})
	b.Run("medium", func(b *testing.B) {
		benchSonic(b, mediumFixture)
	})
	b.Run("large", func(b *testing.B) {
		benchSonic(b, largeFixture)
	})
	b.Run("canada", func(b *testing.B) {
		benchSonic(b, canadaFixture)
	})
	b.Run("citm", func(b *testing.B) {
		benchSonic(b, citmFixture)
	})
	b.Run("twitter", func(b *testing.B) {
		benchSonic(b, twitterFixture)
	})
}

func benchSonic(b *testing.B, s string) {
	b.ReportAllocs()
	b.SetBytes(int64(len(s)))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, ast_, err := ast.Loads(s)
			if err != nil {
				panic(fmt.Errorf("unexpected error: %s", err))
			}
			_ = ast_
		}
	})
}
