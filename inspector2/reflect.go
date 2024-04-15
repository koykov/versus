package inspector2

import (
	"reflect"
	"strconv"
)

func diveReflect(node any, path ...string) any {
	if len(path) == 0 {
		return node
	}
	v := reflect.ValueOf(node)
	switch v.Kind() {
	case reflect.Ptr:
		if elem := v.Elem(); elem.IsValid() && elem.CanInterface() {
			node = elem.Interface()
			return diveReflect(node, path...)
		}
	case reflect.Map:
		for _, f := range v.MapKeys() {
			fv := f.Interface()
			if fvs, ok := fv.(string); ok {
				if fvs == path[0] {
					mv := v.MapIndex(f)
					if mv.IsValid() && mv.CanInterface() {
						return diveReflect(mv.Interface(), path[1:]...)
					}
					return nil
				}
			}
		}
	case reflect.Struct:
		f := v.FieldByName(path[0])
		if f.IsValid() && f.CanInterface() {
			return diveReflect(f.Interface(), path[1:]...)
		}
	case reflect.Slice:
		idx, err := strconv.Atoi(path[0])
		if err != nil {
			return nil
		}
		sv := v.Index(idx)
		if sv.IsValid() && sv.CanInterface() {
			return diveReflect(sv.Interface(), path[1:]...)
		}
	default:
		return nil
	}
	return nil
}
