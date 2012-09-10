package croot

// #include "croot/croot.h"
//
import "C"

import (
	"unsafe"
)

// Branch
type Branch interface {
	Object
	GetAddress() uintptr
	GetClassName() string
}

type branch_impl struct {
	c C.CRoot_Branch
}

func (b *branch_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(b.c)
}

func (b *branch_impl) as_tobject() *object_impl {
	return &object_impl{b.cptr()}
}

func (b *branch_impl) ClassName() string {
	return b.as_tobject().ClassName()
}

func (b *branch_impl) Clone(opt Option) Object {
	return b.as_tobject().Clone(opt)
}

func (b *branch_impl) FindObject(name string) Object {
	return b.as_tobject().FindObject(name)
}

func (b *branch_impl) GetName() string {
	return b.as_tobject().GetName()
}

func (b *branch_impl) GetTitle() string {
	return b.as_tobject().GetTitle()
}

func (b *branch_impl) InheritsFrom(clsname string) bool {
	return b.as_tobject().InheritsFrom(clsname)
}

func (b *branch_impl) Print(option Option) {
	b.as_tobject().Print(option)
}

func (b *branch_impl) GetAddress() uintptr {
	return uintptr(unsafe.Pointer(C.CRoot_Branch_GetAddress(b.c)))
}

func (b *branch_impl) GetClassName() string {
	c_str := C.CRoot_Branch_GetClassName(b.c)
	return C.GoString(c_str)
}

func init() {
	cnvmap["TBranch"] = func(o c_object) Object {
		return &branch_impl{c: (C.CRoot_Branch)(o.cptr())}
	}
}

// EOF
