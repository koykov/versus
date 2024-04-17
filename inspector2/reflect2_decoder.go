package inspector2

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type decR2 interface {
	Decode(ptr unsafe.Pointer, path ...string) any
}

func getDecoderR2(typ reflect2.Type) decR2 {
	ptrType := typ.(*reflect2.UnsafePtrType)
	decoder := decoderOfType(ptrType.Elem())
	return decoder
}

func decoderOfType(typ reflect2.Type) decR2 {
	ifaceType, isIFace := typ.(*reflect2.UnsafeIFaceType)
	if isIFace {
		return &ifaceDecoder{valType: ifaceType}
	}
	return &efaceDecoder{}
}

type ifaceDecoder struct {
	valType *reflect2.UnsafeIFaceType
}

func (decoder *ifaceDecoder) Decode(ptr unsafe.Pointer, path ...string) {
	if iter.ReadNil() {
		decoder.valType.UnsafeSet(ptr, decoder.valType.UnsafeNew())
		return
	}
	obj := decoder.valType.UnsafeIndirect(ptr)
	if reflect2.IsNil(obj) {
		iter.ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	iter.ReadVal(obj)
}

type efaceDecoder struct {
}

func (decoder *efaceDecoder) Decode(ptr unsafe.Pointer, path ...string) {
	pObj := (*interface{})(ptr)
	obj := *pObj
	if obj == nil {
		*pObj = iter.Read()
		return
	}
	typ := reflect2.TypeOf(obj)
	if typ.Kind() != reflect.Ptr {
		*pObj = iter.Read()
		return
	}
	ptrType := typ.(*reflect2.UnsafePtrType)
	ptrElemType := ptrType.Elem()
	if iter.WhatIsNext() == NilValue {
		if ptrElemType.Kind() != reflect.Ptr {
			iter.skipFourBytes('n', 'u', 'l', 'l')
			*pObj = nil
			return
		}
	}
	if reflect2.IsNil(obj) {
		obj := ptrElemType.New()
		iter.ReadVal(obj)
		*pObj = obj
		return
	}
	iter.ReadVal(obj)
}
