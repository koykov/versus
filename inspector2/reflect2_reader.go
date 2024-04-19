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
	ifaceType, isIFace := typ.(*reflect2.UnsafeIFaceType)
	if isIFace {
		return &ifaceReaderR2{valType: ifaceType}
	}
	return &efaceReaderR2{}
}

type ifaceReaderR2 struct {
	valType *reflect2.UnsafeIFaceType
}

func (rdr *ifaceReaderR2) Read(ptr unsafe.Pointer, path ...string) any {
	obj := rdr.valType.UnsafeIndirect(ptr)
	if reflect2.IsNil(obj) {
		return nil
	}
	if len(path) == 0 {
		return obj
	}
	diveReflect2(obj, path[1:]...)
	return obj
}

type efaceReaderR2 struct {
}

func (rdr *efaceReaderR2) Read(ptr unsafe.Pointer, path ...string) any {
	pObj := (*interface{})(ptr)
	obj := *pObj
	if obj == nil {
		*pObj = iter.Read()
		return *pObj
	}
	typ := reflect2.TypeOf(obj)
	if typ.Kind() != reflect.Ptr {
		*pObj = iter.Read()
		return *pObj
	}
	ptrType := typ.(*reflect2.UnsafePtrType)
	ptrElemType := ptrType.Elem()
	if iter.WhatIsNext() == NilValue {
		if ptrElemType.Kind() != reflect.Ptr {
			iter.skipFourBytes('n', 'u', 'l', 'l')
			*pObj = nil
			return nil
		}
	}
	if reflect2.IsNil(obj) {
		obj := ptrElemType.New()
		iter.ReadVal(obj)
		*pObj = obj
		return *pObj
	}
	iter.ReadVal(obj)
	return obj
}
