package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
//"unsafe"
)

// ObjArray
type ObjArray interface {
	Object
	At(idx int64) Object
	GetSize() int64
}

type objarray_impl struct {
	c C.CRoot_ObjArray
}

func (t *objarray_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(t.c)
}

func (t *objarray_impl) as_tobject() *object_impl {
	return &object_impl{t.cptr()}
}

func (t *objarray_impl) ClassName() string {
	return t.as_tobject().ClassName()
}

func (t *objarray_impl) Clone(opt Option) Object {
	return t.as_tobject().Clone(opt)
}

func (t *objarray_impl) FindObject(name string) Object {
	return t.as_tobject().FindObject(name)
}

func (t *objarray_impl) GetName() string {
	return t.as_tobject().GetName()
}

func (t *objarray_impl) GetTitle() string {
	return t.as_tobject().GetTitle()
}

func (t *objarray_impl) InheritsFrom(clsname string) bool {
	return t.as_tobject().InheritsFrom(clsname)
}

func (o *objarray_impl) Print(option Option) {
	o.as_tobject().Print(option)
}

func (o *objarray_impl) At(i int64) Object {
	cptr := C.CRoot_ObjArray_At(o.c, C.int64_t(i))
	obj := object_impl{cptr}
	return to_gocroot(&obj)
}

func (o *objarray_impl) GetSize() int64 {
	return int64(C.CRoot_ObjArray_GetSize(o.c))
}

// EOF
