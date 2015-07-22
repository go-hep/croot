package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

// ROOT is the entry point to the ROOT system.
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

func (r *ROOT) GetVersion() string {
	version := C.CRoot_ROOT_GetVersion(r.c)
	// we don't own 'version'. no need for C.free
	return C.GoString(version)
}

func init() {
	GRoot = &ROOT{C.CRoot_gROOT}
}

// EOF
