package croot

// #include "croot/croot.h"
//
import "C"

import (
	"unsafe"
)

// BranchElement
type BranchElement interface {
	Branch
}

type branchelement_impl struct {
	c C.CRoot_BranchElement
}

func (b *branchelement_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(b.c)
}

func (b *branchelement_impl) as_tobject() *object_impl {
	return &object_impl{b.cptr()}
}

func (b *branchelement_impl) ClassName() string {
	return b.as_tobject().ClassName()
}

func (b *branchelement_impl) Clone(opt Option) Object {
	return b.as_tobject().Clone(opt)
}

func (b *branchelement_impl) FindObject(name string) Object {
	return b.as_tobject().FindObject(name)
}

func (b *branchelement_impl) GetName() string {
	return b.as_tobject().GetName()
}

func (b *branchelement_impl) GetTitle() string {
	return b.as_tobject().GetTitle()
}

func (b *branchelement_impl) InheritsFrom(clsname string) bool {
	return b.as_tobject().InheritsFrom(clsname)
}

func (b *branchelement_impl) Print(option Option) {
	b.as_tobject().Print(option)
}

func (b *branchelement_impl) GetAddress() uintptr {
	return uintptr(unsafe.Pointer(C.CRoot_BranchElement_GetAddress(b.c)))
}

func (b *branchelement_impl) GetClassName() string {
	c_str := C.CRoot_BranchElement_GetClassName(b.c)
	return C.GoString(c_str)
}

func init() {
	cnvmap["TBranchElement"] = func(o c_object) Object {
		return &branchelement_impl{c: (C.CRoot_BranchElement)(o.cptr())}
	}
}
