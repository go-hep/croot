package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

type File interface {
	Object

	Cd(path string) bool
	Close(option string)
	GetFd() int
	Get(namecycle string) Object
	GetTree(namecycle string) Tree //FIXME: should use Get+type-cast
	IsOpen() bool
	Write(name string, opt, bufsiz int) int
}

type file_impl struct {
	c C.CRoot_File
}

func (f *file_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(f.c)
}

func (f *file_impl) as_tobject() *object_impl {
	return &object_impl{f.cptr()}
}

func (f *file_impl) ClassName() string {
	return f.as_tobject().ClassName()
}

func (f *file_impl) Clone(opt Option) Object {
	return f.as_tobject().Clone(opt)
}

func (f *file_impl) FindObject(name string) Object {
	return f.as_tobject().FindObject(name)
}

func (f *file_impl) GetName() string {
	return f.as_tobject().GetName()
}

func (f *file_impl) GetTitle() string {
	return f.as_tobject().GetTitle()
}

func (f *file_impl) InheritsFrom(clsname string) bool {
	return f.as_tobject().InheritsFrom(clsname)
}

func (f *file_impl) Print(option Option) {
	f.as_tobject().Print(option)
}


func OpenFile(name, option, title string, compress, netopt int) File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	f := C.CRoot_File_Open(c_name, (*C.CRoot_Option)(c_option), c_title, C.int32_t(compress), C.int32_t(netopt))
	return &file_impl{c: f}
}

func (f *file_impl) Cd(path string) bool {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))
	return c2bool(C.CRoot_File_cd(f.c, c_path))
}

func (f *file_impl) Close(option string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_File_Close(f.c, (*C.CRoot_Option)(c_option))
}

func (f *file_impl) GetFd() int {
	return int(C.CRoot_File_GetFd(f.c))
}

func (f *file_impl) Get(namecycle string) Object {
	c_name := C.CString(namecycle)
	defer C.free(unsafe.Pointer(c_name))
	o := C.CRoot_File_Get(f.c, c_name)
	if o == nil {
		return nil
	}
	return &object_impl{o}
}

func (f *file_impl) GetTree(namecycle string) Tree {
	o := f.Get(namecycle)
	if o == nil {
		return nil
	}
	c_t := (C.CRoot_Tree)(unsafe.Pointer(o.(c_object).cptr()))
	return &tree_impl{c: c_t, branches: make(map[string]gobranch)}
}

func (f *file_impl) IsOpen() bool {
	return c2bool(C.CRoot_File_IsOpen(f.c))
}

//func (f *file_impl) ReadBuffer(buf, pos, len)
//func (f *file_impl) ReadBuffers
//func (f *file_impl) WriteBuffer

func (f *file_impl) Write(name string, opt, bufsiz int) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.CRoot_File_Write(f.c, c_name, C.int32_t(opt), C.int32_t(bufsiz)))
}

func init() {
	cnvmap["TFile"] = func(o c_object) Object {
		return &file_impl{c: (C.CRoot_File)(o.cptr())}
	}
}
// EOF
