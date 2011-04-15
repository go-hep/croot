package croot

/*
 #cgo LDFLAGS: -lcroot
 #include "croot.h"

 #include <stdlib.h>
 #include <string.h>

 double _go_croot_double_at(double *array, int idx)
 { return array[idx]; }

 */
import "C"

import (
	"unsafe"
	"reflect"
)

// utils
func c2bool(b C.CRoot_Bool) bool {
	if int(b) != 0 {
		return true
	}
	return false
}

//
type Option string

type Object struct {
	o C.CRoot_Object
}

type File struct {
	f C.CRoot_File
}

func OpenFile(name, option, title string, compress, netopt int) File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	f := C.CRoot_File_Open(c_name, c_option, c_title, C.int32_t(compress), C.int32_t(netopt))
	return File{f:f}
}

func (self *File) Cd(path string) bool {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))
	return c2bool(C.CRoot_File_cd(self.f, c_path))
}

func (self *File) Close(option string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_File_Close(self.f, c_option)
}

func (self *File) GetFd() int {
	return int(C.CRoot_File_GetFd(self.f))
}

func (self *File) Get(namecycle string) Object {
	c_name := C.CString(namecycle)
	defer C.free(unsafe.Pointer(c_name))
	o := C.CRoot_File_Get(self.f, c_name)
	return Object{o}
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

// ObjArray
type ObjArray C.CRoot_ObjArray

// Branch
type Branch C.CRoot_Branch

// Tree
type Tree struct {
	t C.CRoot_Tree
}

func NewTree(name, title string, splitlevel int) Tree {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	t := C.CRoot_Tree_new(c_name, c_title, C.int32_t(splitlevel))
	return Tree{t:t}
}

func (self *Tree) Delete() {
	C.CRoot_Tree_delete(self.t)
	self.t = nil
}

func (self *Tree) Branch(name, classname string, obj interface{}, bufsiz, splitlevel int) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_classname := C.CString(classname)
	defer C.free(unsafe.Pointer(c_classname))
	v := reflect.NewValue(obj)
	c_addr := unsafe.Pointer(v.UnsafeAddr()) // FIXME !!!
	b := C.CRoot_Tree_Branch(self.t, c_name, c_classname, c_addr, C.int32_t(bufsiz), C.int32_t(splitlevel))
	return Branch(b)
}

//func (self *Tree) Branch2(name string, address interface{}, leaflist string, int bufsiz) Branch

func (self *Tree) Fill() int {
	return int(C.CRoot_Tree_Fill(self.t))
}

func (self *Tree) GetBranch(name string) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	b := C.CRoot_Tree_GetBranch(self.t, c_name)
	return Branch(b)
}

func (self *Tree) GetEntries() int64 {
	return int64(C.CRoot_Tree_GetEntries(self.t))
}

func (self *Tree) GetEntry(entry int64, getall int) int {
	nbytes := C.CRoot_Tree_GetEntry(self.t, C.int64_t(entry), C.int32_t(getall))
	return int(nbytes)
}

func (self *Tree) GetListOfBranches() ObjArray {
	o := C.CRoot_Tree_GetListOfBranches(self.t)
	return ObjArray(o)
}

func (self *Tree) GetListOfLeaves() ObjArray {
	o := C.CRoot_Tree_GetListOfLeaves(self.t)
	return ObjArray(o)
}

func (self *Tree) GetSelectedRows() int64 {
	return int64(C.CRoot_Tree_GetSelectedRows(self.t))
}

func (self *Tree) GetVal(i int) []float64 {
	c_data := C.CRoot_Tree_GetVal(self.t, C.int32_t(i))
	sz := self.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (self *Tree) GetV1() []float64 {
	c_data := C.CRoot_Tree_GetV1(self.t)
	sz := self.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (self *Tree) GetV2() []float64 {
	c_data := C.CRoot_Tree_GetV2(self.t)
	sz := self.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (self *Tree) GetV3() []float64 {
	c_data := C.CRoot_Tree_GetV3(self.t)
	sz := self.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (self *Tree) GetV4() []float64 {
	c_data := C.CRoot_Tree_GetV4(self.t)
	sz := self.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (self *Tree) GetW() []float64 {
	c_data := C.CRoot_Tree_GetW(self.t)
	sz := self.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (self *Tree) LoadTree(entry int64) int64 {
	return int64(C.CRoot_Tree_LoadTree(self.t, C.int64_t(entry)))
}

// func (self *Tree) MakeClass
// func (self *Tree) Notify

func (self *Tree) Print(option string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_Tree_Print(self.t, c_option)
}

func (self *Tree) Write(name string, option, bufsize int) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.CRoot_Tree_Write(self.t, c_name, C.int32_t(option), C.int32_t(bufsize)))
}

// EOF
