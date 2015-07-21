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
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	f := C.CRoot_File_Open(c_name, (*C.CRoot_Option)(c_option), c_title, C.int32_t(compress), C.int32_t(netopt))
	if f == nil {
		return nil, fmt.Errorf("croot.OpenFile: could not open file [%s]", name)
	}
	return &fileImpl{c: f}, nil
}

func (f *fileImpl) Cd(path string) bool {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))
	return c2bool(C.CRoot_File_cd(f.c, c_path))
}

func (f *fileImpl) Close(option string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_File_Close(f.c, (*C.CRoot_Option)(c_option))
}

func (f *fileImpl) GetFd() int {
	return int(C.CRoot_File_GetFd(f.c))
}

func (f *fileImpl) Get(namecycle string) Object {
	c_name := C.CString(namecycle)
	defer C.free(unsafe.Pointer(c_name))
	o := C.CRoot_File_Get(f.c, c_name)
	if o == nil {
		return nil
	}
	return to_gocroot(&objectImpl{o})
}

func (f *fileImpl) IsOpen() bool {
	return c2bool(C.CRoot_File_IsOpen(f.c))
}

//func (f *file_impl) ReadBuffer(buf, pos, len)
//func (f *file_impl) ReadBuffers
//func (f *file_impl) WriteBuffer

func (f *fileImpl) Write(name string, opt, bufsiz int) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.CRoot_File_Write(f.c, c_name, C.int32_t(opt), C.int32_t(bufsiz)))
}
