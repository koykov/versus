package inspector2

import (
	"github.com/modern-go/reflect2"
)

func diveReflect2(node any, path ...string) any {
	if len(path) == 0 {
		return node
	}
	ptr := reflect2.TypeOf(node)
	switch path[0] {
	case "L1":
		typ := ptr.(*reflect2.UnsafePtrType)
		utyp := typ.Elem().(*reflect2.UnsafeStructType)
		f := utyp.FieldByName(path[0])
		// var t T
		fnode := f.Get(node)
		_ = fnode
	}
	// if path[0] == "L1" {
	// ptr := reflect2.TypeOf(node)
	// if typ, ok := ptr.(*reflect2.UnsafePtrType); ok {
	// 	utyp := typ.Elem().(*reflect2.UnsafeStructType)
	// 	f := utyp.FieldByName(path[0])
	// 	fnode := f.Get(node)
	// 	return diveReflect2(fnode, path[1:]...)
	// } else if typ, ok := ptr.(*reflect2.UnsafeStructType); ok {
	// 	f := typ.FieldByName(path[0])
	// 	fnode := f.Get(node)
	// 	fmt.Printf("%t\n", fnode)
	// 	return diveReflect2(fnode, path[1:]...)
	// } else {
	// 	return nil
	// }
	// ptyp := ptr.Elem().(*reflect2.UnsafePtrType)
	// typ := ptyp.Elem().(*reflect2.UnsafeStructType)
	// f := typ.FieldByName(path[0])
	// fnode := f.Get(node)
	// return diveReflect2(fnode, path[1:]...)
	// }
	// if path[0] == "L2" {
	// 	ptr := reflect2.TypeOfPtr(node)
	// 	typ := ptr.Elem().(*reflect2.UnsafeStructType)
	// 	f := typ.FieldByName(path[0])
	// 	fnode := f.Get(node)
	// 	return diveReflect2(fnode, path[1:]...)
	// }
	return nil
}
