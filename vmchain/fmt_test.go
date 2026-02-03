package vmchain

import (
	"fmt"
	"testing"

	"github.com/VictoriaMetrics/metrics"
)

func BenchmarkFmt(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			name := fmt.Sprintf(`myservice_feature_counter_fmt{stage="%s",op="%s",roleID="%d"}`, stage, op, roleID)
			metrics.GetOrCreateCounter(name).Inc()
		}
	})
	b.Run("parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.ReportAllocs()
			for pb.Next() {
				name := fmt.Sprintf(`myservice_feature_counter_fmt{stage="%s",op="%s",roleID="%d"}`, stage, op, roleID)
				metrics.GetOrCreateCounter(name).Inc()
			}
		})
	})
}
