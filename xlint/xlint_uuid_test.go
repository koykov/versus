package xlint

import (
	"testing"

	"github.com/koykov/xlint"
)

func BenchmarkValidateXlintUUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for uuid, valid := range stage {
			ok, _ := xlint.ValidateUUIDStr(uuid)
			if ok != valid {
				b.Errorf("validation failed on %s, got %v expected %v", uuid, ok, valid)
			}
		}
	}
}
