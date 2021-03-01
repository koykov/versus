package xlint

import (
	"testing"

	uuid2 "github.com/google/uuid"
)

func BenchmarkValidateNativeUUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for uuid, valid := range stage {
			_, err := uuid2.Parse(uuid)
			ok := err == nil
			if ok != valid {
				b.Errorf("validation failed on %s, got %v expected %v", uuid, ok, valid)
			}
		}
	}
}
