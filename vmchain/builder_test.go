package vmchain

import (
	"strconv"
	"strings"
	"testing"

	"github.com/VictoriaMetrics/metrics"
)

func BenchmarkStringsBuilder(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var sb strings.Builder
			sb.WriteString(`myservice_feature_counter_builder{stage="`)
			sb.WriteString(stage)
			sb.WriteString(`",op="`)
			sb.WriteString(op)
			sb.WriteString(`",roleID="`)
			sb.WriteString(strconv.Itoa(roleID))
			sb.WriteString(`"}`)
			metrics.GetOrCreateCounter(sb.String()).Inc()
		}
	})
	b.Run("parallel", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var sb strings.Builder
				sb.WriteString(`myservice_feature_counter_builder{stage="`)
				sb.WriteString(stage)
				sb.WriteString(`",op="`)
				sb.WriteString(op)
				sb.WriteString(`",roleID="`)
				sb.WriteString(strconv.Itoa(roleID))
				sb.WriteString(`"}`)
				metrics.GetOrCreateCounter(sb.String()).Inc()
			}
		})
	})
}
