package inspector2

import (
	"reflect"

	"github.com/modern-go/reflect2"
)

func diveReflect2(node any, path ...string) any {
	if len(path) == 0 {
		return node
	}
	typ := reflect2.TypeOf(node)
	return diveR2V2(typ, path...)
}

func diveR2V2(typ reflect2.Type, path ...string) any {
	kind := typ.Kind()
	switch kind {
	case reflect.Interface:
		// todo implement me
	case reflect.Struct:
		utyp := typ.(*reflect2.UnsafeStructType)
		for i := 0; i < utyp.NumField(); i++ {
			field := utyp.Field(i)
			if field.Name() == path[0] {
				return diveR2V2(field.Type(), path[1:]...)
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
		return diveR2V2(elem, path...)
	case reflect.Bool:
	case reflect.Int:

	case reflect.Int8:

	case reflect.Int16:

	case reflect.Int32:

	case reflect.Int64:

	case reflect.Uint:

	case reflect.Uint8:

	case reflect.Uint16:

	case reflect.Uint32:

	case reflect.Uint64:

	case reflect.Uintptr:

	case reflect.Float32:

	case reflect.Float64:

	case reflect.Complex64:

	case reflect.Complex128:

	case reflect.String:
		var s string
		a := typ.Indirect(&s)
		return a
	default:
		// reflect.Invalid, reflect.Chan and reflect.UnsafePointer is unsupported.
		return nil
	}
	return nil
}
