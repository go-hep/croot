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
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c := C.CRoot_ROOT_GetFile(r.c, c_name)
	if c == nil {
		return nil
	}
	return &fileImpl{c}
}

func init() {
	GRoot = &ROOT{C.CRoot_gROOT}
}

// EOF
