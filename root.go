package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

type ROOT struct {
	c C.CRoot_ROOT
}

var GRoot *ROOT = nil

func (r *ROOT) GetFile(name string) File {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	c := C.CRoot_ROOT_GetFile(r.c, cname)
	if c == nil {
		return nil
	}
	return &fileImpl{c}
}

func init() {
	GRoot = &ROOT{C.CRoot_gROOT}
}

// EOF
