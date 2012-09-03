package croot

// #include "croot/croot.h"
//
import "C"

import (
	"unsafe"
)

// Branch
type Branch struct {
	c C.CRoot_Branch
}

func (b *Branch) GetAddress() uintptr {
	return uintptr(unsafe.Pointer(C.CRoot_Branch_GetAddress(b.c)))
}

// EOF
