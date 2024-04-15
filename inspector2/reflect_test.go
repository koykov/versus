package inspector2

import "testing"

func TestReflect(t *testing.T) {
	t.Run("obj.L1.L2.L3.S", func(t *testing.T) {
		t.Log(diveReflect(&obj, "L1", "L2", "L3", "S")) // prints "foobar"
	})
}

func BenchmarkReflect(b *testing.B) {
	b.Run("obj.L1.L2.L3.S", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			diveReflect(&obj, "L1", "L2", "L3", "S")
		}
	})
}
