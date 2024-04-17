package inspector2

import "github.com/modern-go/reflect2"

func diveReflect2(node any, path ...string) any {
	if len(path) == 0 {
		return node
	}
	typ := reflect2.TypeOf(node)
	dec := getDecoderR2(typ)
	ptr := reflect2.PtrOf(node)
	if ptr == nil {
		return nil
	}
	return dec.Decode(ptr, path...)
}
