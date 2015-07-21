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

type branchElementImpl struct {
	c C.CRoot_BranchElement
}

func (b *branchElementImpl) GetAddress() uintptr {
	return uintptr(unsafe.Pointer(C.CRoot_BranchElement_GetAddress(b.c)))
}

func (b *branchElementImpl) GetClassName() string {
	cstr := C.CRoot_BranchElement_GetClassName(b.c)
	return C.GoString(cstr)
}
