package inspector2

import "testing"

func TestReflect2(t *testing.T) {
	t.Run("obj.L1.L2.L3.S", func(t *testing.T) {
		t.Log(diveReflect2(&obj, "L1", "L2", "L3", "S")) // prints "foobar"
	})
}