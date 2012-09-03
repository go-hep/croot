package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

type File struct {
	f C.CRoot_File
}

func OpenFile(name, option, title string, compress, netopt int) *File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	f := C.CRoot_File_Open(c_name, (*C.CRoot_Option)(c_option), c_title, C.int32_t(compress), C.int32_t(netopt))
	return &File{f: f}
}

func (self *File) Cd(path string) bool {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))
	return c2bool(C.CRoot_File_cd(self.f, c_path))
}

func (self *File) Close(option string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_File_Close(self.f, (*C.CRoot_Option)(c_option))
}

func (self *File) GetFd() int {
	return int(C.CRoot_File_GetFd(self.f))
}

func (self *File) Get(namecycle string) *Object {
	c_name := C.CString(namecycle)
	defer C.free(unsafe.Pointer(c_name))
	o := C.CRoot_File_Get(self.f, c_name)
	if o == nil {
		return nil
	}
	return &Object{o}
}

func (self *File) GetTree(namecycle string) *Tree {
	o := self.Get(namecycle)
	if o == nil {
		return nil
	}
	c_t := (C.CRoot_Tree)(unsafe.Pointer(o.o))
	return &Tree{t: c_t, branches: make(map[string]gobranch)}
}

func (self *File) IsOpen() bool {
	return c2bool(C.CRoot_File_IsOpen(self.f))
}

//func (self *File) ReadBuffer(buf, pos, len)
//func (self *File) ReadBuffers
//func (self *File) WriteBuffer

func (self *File) Write(name string, opt, bufsiz int) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.CRoot_File_Write(self.f, c_name, C.int32_t(opt), C.int32_t(bufsiz)))
}

// EOF
