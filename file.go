package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type File interface {
	Object

	Cd(path string) bool
	Close(option string)
	GetFd() int
	Get(namecycle string) Object
	IsOpen() bool
	Write(name string, opt, bufsiz int) int
}

type fileImpl struct {
	c C.CRoot_File
}

func OpenFile(name, option, title string, compress, netopt int) (File, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	copt := C.CString(option)
	defer C.free(unsafe.Pointer(copt))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	f := C.CRoot_File_Open(cname, (*C.CRoot_Option)(copt), ctitle, C.int32_t(compress), C.int32_t(netopt))
	if f == nil {
		return nil, fmt.Errorf("croot.OpenFile: could not open file [%s]", name)
	}
	return &fileImpl{c: f}, nil
}

func (f *fileImpl) Cd(path string) bool {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return c2bool(C.CRoot_File_cd(f.c, cpath))
}

func (f *fileImpl) Close(option string) {
	copt := C.CString(option)
	defer C.free(unsafe.Pointer(copt))

	C.CRoot_File_Close(f.c, (*C.CRoot_Option)(copt))
}

func (f *fileImpl) GetFd() int {
	return int(C.CRoot_File_GetFd(f.c))
}

func (f *fileImpl) Get(namecycle string) Object {
	cname := C.CString(namecycle)
	defer C.free(unsafe.Pointer(cname))
	o := C.CRoot_File_Get(f.c, cname)
	if o == nil {
		return nil
	}
	return to_gocroot(&objectImpl{o})
}

func (f *fileImpl) IsOpen() bool {
	return c2bool(C.CRoot_File_IsOpen(f.c))
}

//func (f *fileImpl) ReadBuffer(buf, pos, len)
//func (f *fileImpl) ReadBuffers
//func (f *fileImpl) WriteBuffer

func (f *fileImpl) Write(name string, opt, bufsiz int) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.CRoot_File_Write(f.c, cname, C.int32_t(opt), C.int32_t(bufsiz)))
}
