package inspector2

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/koykov/versus/inspector2/inspector2_ins"
)

func TestInspector(t *testing.T) {
	t.Run("obj.L1.L2.L3.S", func(t *testing.T) {
		ins, _ := inspector.GetInspector("T")
		x, _ := ins.Get(&obj, "L1", "L2", "L3", "S")
		t.Log(*x.(*string)) // prints "foobar"
	})
}

func BenchmarkInspector(b *testing.B) {
	b.Run("obj.L1.L2.L3.S", func(b *testing.B) {
		b.ReportAllocs()
		var ins types_ins.TInspector
		var buf any
		for i := 0; i < b.N; i++ {
			_ = ins.GetTo(&obj, &buf, "L1", "L2", "L3", "S")
		}
	})
}
