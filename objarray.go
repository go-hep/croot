package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

// ObjArray
type ObjArray interface {
	Object
	At(idx int64) Object
	GetSize() int64
	GetEntries() int64
}

type objArrayImpl struct {
	c C.CRoot_ObjArray
}

func (o *objArrayImpl) At(i int64) Object {
	cptr := C.CRoot_ObjArray_At(o.c, C.int64_t(i))
	obj := objectImpl{cptr}
	return to_gocroot(&obj)
}

func (o *objArrayImpl) GetSize() int64 {
	return int64(C.CRoot_ObjArray_GetSize(o.c))
}

func (o *objArrayImpl) GetEntries() int64 {
	return int64(C.CRoot_ObjArray_GetEntries(o.c))
}

// EOF
