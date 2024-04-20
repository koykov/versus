package inspector2

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func diveReflect2(node any, path ...string) any {
	if len(path) == 0 {
		return node
	}
	typ := reflect2.TypeOf(node)
	ptr := reflect2.PtrOf(node)
	return diveR2V2(typ, ptr, path...)
}

func diveR2V2(typ reflect2.Type, ptr unsafe.Pointer, path ...string) any {
	kind := typ.Kind()
	switch kind {
	case reflect.Interface:
		// todo implement me
	case reflect.Struct:
		utyp := typ.(*reflect2.UnsafeStructType)
		for i := 0; i < utyp.NumField(); i++ {
			field := utyp.Field(i)
			if field.Name() == path[0] {
				ptr = unsafe.Pointer(uintptr(ptr) + field.Offset())
				return diveR2V2(field.Type(), ptr, path[1:]...)
			}
		}
	case reflect.Array:
		// todo implement me
	case reflect.Slice:
		// todo implement me
	case reflect.Map:
		// todo implement me
	case reflect.Pointer:
		utyp := typ.(*reflect2.UnsafePtrType)
		elem := utyp.Elem()
		node := elem.UnsafeIndirect(ptr)
		ptr = reflect2.PtrOf(node)
		return diveR2V2(elem, ptr, path...)
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String:
		return typ.UnsafeIndirect(ptr)
	default:
		// reflect.Invalid, reflect.Chan and reflect.UnsafePointer is unsupported.
		return nil
	}
	return nil
}
