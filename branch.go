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

type branchImpl struct {
	c C.CRoot_Branch
}

func (b *branchImpl) GetAddress() uintptr {
	return uintptr(unsafe.Pointer(C.CRoot_Branch_GetAddress(b.c)))
}

// func (b *branch_impl) GetObject() uintptr {
// 	return uintptr(unsafe.Pointer(C.CRoot_Branch_GetObject(b.c)))
// }

func (b *branchImpl) GetClassName() string {
	cstr := C.CRoot_Branch_GetClassName(b.c)
	return C.GoString(cstr)
}

func (b *branchImpl) GetListOfLeaves() []Leaf {
	c := C.CRoot_Branch_GetListOfLeaves(b.c)
	objs := objArrayImpl{c: c}
	leaves := make([]Leaf, objs.GetEntries())
	for i := 0; i < len(leaves); i++ {
		obj := objs.At(int64(i))
		leaf := b.GetLeaf(obj.GetName())
		leaves[i] = leaf
	}
	return leaves
}

func (b *branchImpl) GetLeaf(name string) Leaf {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.CRoot_Branch_GetLeaf(b.c, cname)
	if c == nil {
		return nil
	}
	return &leafImpl{c: c}
}

// EOF
