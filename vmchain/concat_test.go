package vmchain

import (
	"strconv"
	"testing"

	"github.com/VictoriaMetrics/metrics"
)

func BenchmarkConcat(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			name := `myservice_feature_counter_concat{stage="` + stage + `",op="` + op + `",roleID="` + strconv.Itoa(roleID) + `"}`
			metrics.GetOrCreateCounter(name).Inc()
		}
	})
	b.Run("parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.ReportAllocs()
			for pb.Next() {
				name := `myservice_feature_counter_concat{stage="` + stage + `",op="` + op + `",roleID="` + strconv.Itoa(roleID) + `"}`
				metrics.GetOrCreateCounter(name).Inc()
			}
		})
	})
}
