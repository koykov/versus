package vmchain

import (
	"testing"

	"github.com/koykov/vmchain"
)

var (
	stage  = "auth"
	op     = "login"
	roleID = 123
)

func BenchmarkVMChain(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			vmchain.Counter("myservice_feature_counter_vmchain").
				WithLabel("stage", stage).
				WithLabel("op", op).
				WithAnyLabel("roleID", roleID).Inc()
		}
	})
	b.Run("parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.ReportAllocs()
			for pb.Next() {
				vmchain.Counter("myservice_feature_counter_vmchain").
					WithLabel("stage", stage).
					WithLabel("op", op).
					WithAnyLabel("roleID", roleID).Inc()
			}
		})
	})
}
