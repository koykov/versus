package inspector

import (
	"bytes"
	"testing"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj_ins"
)

func BenchmarkInspectorGet(b *testing.B) {
	insp := &testobj_ins.TestObjectInspector{}

	b.Run("id", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p0...)
			if *buf.(*string) != "foo" {
				b.Error("object.Id: mismatch result and expectation")
			}
		}
	})
	b.Run("name", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p1...)
			if !bytes.Equal(*buf.(*[]byte), expectFoo) {
				b.Error("object.Name: mismatch result and expectation")
			}
		}
	})
	b.Run("permission[23]", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p2...)
			if *buf.(*bool) != false {
				b.Error("object.Permission[23]: mismatch result and expectation")
			}
		}
	})
	b.Run("flags[export]", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p3...)
			if *buf.(*int32) != 17 {
				b.Error("object.Flags[export]: mismatch result and expectation")
			}
		}
	})
	b.Run("finance.balance", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p4...)
			if *buf.(*float64) != 9000 {
				b.Error("object.Finance.Balance: mismatch result and expectation")
			}
		}
	})
	b.Run("finance.history[1].dateUnix", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p5...)
			if *buf.(*int64) != 153465345246 {
				b.Error("object.Finance.History[1].DateUnix: mismatch result and expectation")
			}
		}
	})
	b.Run("finance.history[0].comment", func(b *testing.B) {
		var buf interface{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.GetTo(testO, &buf, p6...)
			if !bytes.Equal(*buf.(*[]byte), expectComment) {
				b.Error("object.Finance.History[0].Comment: mismatch result and expectation")
			}
		}
	})
}

func BenchmarkInspectorCmp(b *testing.B) {
	insp := &testobj_ins.TestObjectInspector{}

	b.Run("id", func(b *testing.B) {
		var buf bool
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.Cmp(testO, inspector.OpEq, "foo", &buf, p0...)
			if !buf {
				b.Error("object.Id: mismatch result and expectation")
			}
		}
	})
	b.Run("name", func(b *testing.B) {
		var buf bool
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.Cmp(testO, inspector.OpEq, "bar", &buf, p1...)
			if !buf {
				b.Error("object.Name: mismatch result and expectation")
			}
		}
	})
	b.Run("status", func(b *testing.B) {
		var buf bool
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.Cmp(testO, inspector.OpGtq, "60", &buf, p7...)
			if !buf {
				b.Error("object.Status: mismatch result and expectation")
			}
		}
	})
	b.Run("finance.moneyIn", func(b *testing.B) {
		var buf bool
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = insp.Cmp(testO, inspector.OpLtq, "5000", &buf, p8...)
			if !buf {
				b.Error("object.Finance.MoneyIn: mismatch result and expectation")
			}
		}
	})
}

// todo implement me
// func BenchmarkInspectorSet(b *testing.B) {
// 	// ...
// }
