package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
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
	GetListOfLeaves() []Leaf
	GetLeaf(n string) Leaf
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

// func (b *branch_impl) GetObject() uintptr {
// 	return uintptr(unsafe.Pointer(C.CRoot_Branch_GetObject(b.c)))
// }

func (b *branch_impl) GetClassName() string {
	c_str := C.CRoot_Branch_GetClassName(b.c)
	return C.GoString(c_str)
}

func (b *branch_impl) GetListOfLeaves() []Leaf {
	c := C.CRoot_Branch_GetListOfLeaves(b.c)
	objs := objarray_impl{c: c}
	leaves := make([]Leaf, objs.GetEntries())
	for i := 0; i < len(leaves); i++ {
		obj := objs.At(int64(i))
		leaf := b.GetLeaf(obj.GetName())
		leaves[i] = leaf
	}
	return leaves
}

func (b *branch_impl) GetLeaf(name string) Leaf {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c := C.CRoot_Branch_GetLeaf(b.c, c_name)
	if c == nil {
		return nil
	}
	return &leaf_impl{c: c}
}

func init() {
	cnvmap["TBranch"] = func(o c_object) Object {
		return &branch_impl{c: (C.CRoot_Branch)(o.cptr())}
	}
}

// EOF
