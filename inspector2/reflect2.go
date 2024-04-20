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
	case reflect.Struct:
	case reflect.Array:
		// todo implement me
	case reflect.Slice:
		// todo implement me
	case reflect.Map:
		// todo implement me
	case reflect.Ptr:
		utyp := typ.(*reflect2.UnsafePtrType)
		elem := utyp.Elem()
		return diveR2V2(elem, path...)
	default:
		return nil
	}
	return nil
}
