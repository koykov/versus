package inspector

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func BenchmarkReflect2Get(b *testing.B) {
	b.ReportAllocs()

	b.Run("id", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ptr := reflect2.TypeOfPtr(testO)
			typ := ptr.Elem().(*reflect2.UnsafeStructType)
			f := typ.FieldByName("Id")
			fptr := f.UnsafeGet(unsafe.Pointer(testO))
			if *((*string)(fptr)) != "foo" {
				b.Error("object.Id: mismatch result and expectation")
			}
		}
	})
	b.Run("name", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ptr := reflect2.TypeOfPtr(testO)
			typ := ptr.Elem().(*reflect2.UnsafeStructType)
			f := typ.FieldByName("Name")
			fptr := f.UnsafeGet(unsafe.Pointer(testO))
			if !bytes.Equal(*((*[]byte)(fptr)), expectFoo) {
				b.Error("object.Name: mismatch result and expectation")
			}
		}
	})
}

func BenchmarkReflect2Cmp(b *testing.B) {
	// In general this func is full analogue of BenchmarkReflect2Get(), due to the same comparisons.
	b.ReportAllocs()

	b.Run("id", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ptr := reflect2.TypeOfPtr(testO)
			typ := ptr.Elem().(*reflect2.UnsafeStructType)
			f := typ.FieldByName("Id")
			fptr := f.UnsafeGet(unsafe.Pointer(testO))
			if *((*string)(fptr)) != "foo" { // the same as https://github.com/koykov/versus/blob/master/inspector/inspector_test.go#L93
				b.Error("object.Id: mismatch result and expectation")
			}
		}
	})
	b.Run("name", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ptr := reflect2.TypeOfPtr(testO)
			typ := ptr.Elem().(*reflect2.UnsafeStructType)
			f := typ.FieldByName("Name")
			fptr := f.UnsafeGet(unsafe.Pointer(testO))
			if !bytes.Equal(*((*[]byte)(fptr)), expectFoo) { // the same as https://github.com/koykov/versus/blob/master/inspector/inspector_test.go#L103
				b.Error("object.Name: mismatch result and expectation")
			}
		}
	})
}

// todo implement me
// func BenchmarkReflect2Set(b *testing.B) {
// 	// ...
// }
