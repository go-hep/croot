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
	"fmt"
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

func OpenFile(name, option, title string, compress, netopt int) *File {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	f := C.CRoot_File_Open(c_name, c_option, c_title, C.int32_t(compress), C.int32_t(netopt))
	return &File{f:f}
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
	return &Tree{t:c_t}
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

func NewTree(name, title string, splitlevel int) *Tree {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	t := C.CRoot_Tree_new(c_name, c_title, C.int32_t(splitlevel))
	return &Tree{t:t}
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
	v := reflect.ValueOf(obj)
	c_addr := unsafe.Pointer(v.UnsafeAddr()) // FIXME !!!
	b := C.CRoot_Tree_Branch(self.t, c_name, c_classname, c_addr, C.int32_t(bufsiz), C.int32_t(splitlevel))
	return Branch(b)
}

func (self *Tree) Branch2(name string, objaddr interface{}, leaflist string, bufsiz int) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	v := reflect.ValueOf(objaddr)
	c_addr := unsafe.Pointer(v.Elem().UnsafeAddr())

	c_leaflist := C.CString(leaflist)
	defer C.free(unsafe.Pointer(c_leaflist))

	b := C.CRoot_Tree_Branch2(self.t, c_name, c_addr, c_leaflist, C.int32_t(bufsiz))
	return Branch(b)
}

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

func (self *Tree) SetBranchAddress(name string, objaddr interface{}) int32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	v := reflect.ValueOf(objaddr)
	c_addr := unsafe.Pointer(v.Elem().UnsafeAddr())
	
	rc := C.CRoot_Tree_SetBranchAddress(self.t, c_name, c_addr, nil)
	return int32(rc)
}

func (self *Tree) Write(name string, option, bufsize int) int {
	if len(name) != 0 {
		c_name := C.CString(name)
		defer C.free(unsafe.Pointer(c_name))
		return int(C.CRoot_Tree_Write(self.t, c_name, C.int32_t(option), C.int32_t(bufsize)))
	}
	c_name := (*C.char)(unsafe.Pointer(nil))
	return int(C.CRoot_Tree_Write(self.t, c_name, C.int32_t(option), C.int32_t(bufsize)))
}

// TRandom
type Random struct {
	r C.CRoot_Random
}
var GRandom *Random = nil

func (r *Random) Gaus(mean, sigma float64) float64 {
	val := C.CRoot_Random_Gaus(r.r, C.double(mean), C.double(sigma))
	return float64(val)
}

func (r *Random) Rannorf() (a,b float32) {
	c_a := (*C.float)(unsafe.Pointer(&a))
	c_b := (*C.float)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannorf(r.r, c_a, c_b)
	return
}

func (r *Random) Rannord() (a,b float64) {
	c_a := (*C.double)(unsafe.Pointer(&a))
	c_b := (*C.double)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannord(r.r, c_a, c_b)
	return
}

func (r *Random) Rndm(i int) float64 {
	val := C.CRoot_Random_Rndm(r.r, C.int32_t(i))
	return float64(val)
}

func init() {
	fmt.Printf(":: initializing go-croot...\n")
	GRandom = &Random{r:C.CRoot_gRandom}
	fmt.Printf(":: initializing go-croot...[done]\n")
}

// EOF
