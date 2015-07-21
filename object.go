package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

type Object interface {
	ClassName() string
	Clone(option Option) Object
	FindObject(name string) Object
	GetName() string
	GetTitle() string
	InheritsFrom(clsname string) bool
	Print(opt Option)
}

type objectImpl struct {
	c C.CRoot_Object
}

func (o *objectImpl) cptr() C.CRoot_Object {
	return o.c
}

func (o *objectImpl) ClassName() string {
	c_str := C.CRoot_Object_ClassName(o.c)
	return C.GoString(c_str)
}

func (o *objectImpl) Clone(newname Option) Object {
	c_str := C.CString(string(newname))
	defer C.free(unsafe.Pointer(c_str))
	newobj := C.CRoot_Object_Clone(o.c, c_str)
	if newobj == nil {
		return nil
	}
	return &objectImpl{c: newobj}
}

func (o *objectImpl) FindObject(name string) Object {
	c_str := C.CString(name)
	defer C.free(unsafe.Pointer(c_str))
	obj := C.CRoot_Object_FindObject(o.c, c_str)
	if obj == nil {
		return nil
	}
	return &objectImpl{c: obj}
}

func (o *objectImpl) GetName() string {
	c_str := C.CRoot_Object_GetName(o.c)
	// we do not own c_str!!
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

func (o *objectImpl) GetTitle() string {
	c_str := C.CRoot_Object_GetTitle(o.c)
	// we do not own c_str!!
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

func (o *objectImpl) InheritsFrom(clsname string) bool {
	c_str := C.CString(clsname)
	defer C.free(unsafe.Pointer(c_str))
	return c2bool(C.CRoot_Object_InheritsFrom(o.c, c_str))
}

func (o *objectImpl) Print(option Option) {
	c_option := C.CString(string(option))
	defer C.free(unsafe.Pointer(c_option))
	C.CRoot_Object_Print(o.c, (*C.CRoot_Option)(c_option))
}

func init() {
	cnvmap["TObject"] = func(o c_object) Object {
		return &objectImpl{c: (C.CRoot_Object)(o.cptr())}
	}
}

// EOF
