package inspector2

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type readerR2 interface {
	Read(ptr unsafe.Pointer, path ...string) any
}

func getReaderR2(typ reflect2.Type) readerR2 {
	ptrType := typ.(*reflect2.UnsafePtrType)
	decoder := readerOfType(ptrType.Elem())
	return decoder
}

func readerOfType(typ reflect2.Type) readerR2 {
	kind := typ.Kind()
	switch kind {
	case reflect.Interface:
		return &ifaceReaderR2{typ}
	case reflect.Struct:
		return readerStructR2{typ}
	case reflect.Array:
		return readerArrayR2{typ}
	case reflect.Slice:
		return readerSliceR2{typ}
	case reflect.Map:
		return readerMapR2{typ}
	case reflect.Ptr:
		return readerOptionalR2{typ}
	default:
		return nil
	}
}

type ifaceReaderR2 struct {
	valType reflect2.Type
}

func (encoder *ifaceReaderR2) Read(ptr unsafe.Pointer, path ...string) any {
	obj := encoder.valType.UnsafeIndirect(ptr)
	if len(path) == 0 {
		return obj
	}
	return diveReflect2(obj, path[1:]...)
}
