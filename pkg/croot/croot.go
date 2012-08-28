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
	"fmt"
	"reflect"
	"unsafe"

	"github.com/sbinet/go-ffi/pkg/ffi"
)

// utils
func c2bool(b C.CRoot_Bool) bool {
	if int(b) != 0 {
		return true
	}
	return false
}
func bool2c(b bool) C.CRoot_Bool {
	if true {
		return C.CRoot_Bool(1)
	}
	return C.CRoot_Bool(0)
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

// ObjArray
type ObjArray struct {
	c C.CRoot_ObjArray
}

// Branch
type Branch struct {
	c C.CRoot_Branch
}

// Tree
type Tree struct {
	t        C.CRoot_Tree
	branches map[string]gobranch
}

type gobranch struct {
	g   reflect.Value // pointer to go-value
	c   ffi.Value     // the equivalent C-value
	enc *ffi.Encoder  // go->c encoder
	dec *ffi.Decoder  // c->go decoder
}

func NewTree(name, title string, splitlevel int) *Tree {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	t := C.CRoot_Tree_new(c_name, c_title, C.int32_t(splitlevel))
	b := make(map[string]gobranch)
	return &Tree{t: t, branches: b}
}

func (t *Tree) Delete() {
	C.CRoot_Tree_delete(t.t)
	t.t = nil
	t.branches = nil
}

func (t *Tree) Branch(name string, obj interface{}, bufsiz, splitlevel int) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	v := reflect.Indirect(reflect.ValueOf(obj))
	// t := v.Type()
	// register the type with Reflex
	//genreflex(t)
	br := gobranch{g: reflect.ValueOf(obj), c: ffi.ValueOf(v.Interface())}
	br.enc = ffi.NewEncoder(br.c)
	ct := br.c.Type()
	println(">>>>>", br.c.Type().Name(), v.Type().Name())
	t.branches[name] = br
	c_addr := br.c.UnsafeAddr()
	//c_addr := unsafe.Pointer(br.c.UnsafeAddress())
	//c_addr := unsafe.Pointer(&br.c.Buffer()[0])

	// err := ffi.Associate(ct, v.Type())
	// if err != nil {
	// 	panic(err.Error())
	// }
	if ct.GoType() == nil {
		panic("no Go-type for ffi.Type ["+ct.Name()+"] !!")
	}
	// register the type with Reflex
	genreflex(br.c.Type())

	classname := to_cxx_name(ct.GoType())
	c_classname := C.CString(classname)
	defer C.free(unsafe.Pointer(c_classname))

	// fmt.Printf("Tree.Branch(%s, [%s], [%v], [%v] [%v])...\n", name, classname, br.c.Type(), br.c.UnsafeAddress(), br.c.Buffer()[0])

	// c_buff := br.c.Buffer()
	// fmt.Printf(":>> %v (%v) (%v) \n:<<(%v) [%d,%d]\n",br.c.Buffer(), br.g, br.c.UnsafeAddress(), c_buff, len(c_buff), len(br.c.Buffer()))

	b := C.CRoot_Tree_Branch(t.t, c_name, c_classname, unsafe.Pointer(&c_addr), C.int32_t(bufsiz), C.int32_t(splitlevel))
	return Branch{c: b}
}

func (t *Tree) Branch2(name string, objaddr interface{}, leaflist string, bufsiz int) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	v := reflect.Indirect(reflect.ValueOf(objaddr))
	br := gobranch{g: reflect.ValueOf(objaddr), c: ffi.ValueOf(v.Interface())}
	br.enc = ffi.NewEncoder(br.c)
	//ct := br.c.Type()
	t.branches[name] = br
	c_addr := br.c.UnsafeAddr()
	genreflex(br.c.Type())

	c_leaflist := C.CString(leaflist)
	defer C.free(unsafe.Pointer(c_leaflist))

	b := C.CRoot_Tree_Branch2(t.t, c_name, unsafe.Pointer(c_addr), c_leaflist, C.int32_t(bufsiz))
	return Branch{c: b}
}

func (t *Tree) Fill() int {
	// synchronize branches: update ffi.Value
	for k, br := range t.branches {
		fmt.Printf("::> [%v] %v (%v) (%v)\n", 
		 	k, br.c.Buffer(), br.g, br.c.UnsafeAddr())
		err := br.enc.Encode(br.g.Elem().Interface())
		if err != nil {
			fmt.Printf("**err** problem encoding branch [%s]: %s\n", k, err)
			panic("")
			continue
		}
		fmt.Printf("<:: [%v] %v (%v) (%v) (%v)\n", 
			k, br.c.Buffer(), br.g, br.c.UnsafeAddr(), br.c.Field(0).UnsafeAddr())
	}
	fmt.Printf("--- Fill[%d] --- (%p)\n", t.GetEntries(), t.t)
	o := int(C.CRoot_Tree_Fill(t.t))
	for k, br := range t.branches {
		fmt.Printf("=== [%v] %v (%v) (%v)\n", 
		k, br.c.Buffer(), br.g, br.c.UnsafeAddr())
	}
	return o
}

func (t *Tree) GetBranch(name string) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	b := C.CRoot_Tree_GetBranch(t.t, c_name)
	return Branch{c: b}
}

func (t *Tree) GetEntries() int64 {
	return int64(C.CRoot_Tree_GetEntries(t.t))
}

func (t *Tree) GetEntry(entry int64, getall int) int {
	nbytes := C.CRoot_Tree_GetEntry(t.t, C.int64_t(entry), C.int32_t(getall))
	if nbytes > 0 {
		for k, br := range t.branches {
			//fmt.Printf("::> %v (%v) (%v)\n",br.c.Buffer(), br.g, br.c.UnsafeAddress())
			err := br.dec.Decode(br.g)
			if err != nil {
				fmt.Printf("**err** could not decode branch [%s]: %s\n", k, err)
			}
		}
	}
	return int(nbytes)
}

func (t *Tree) GetListOfBranches() ObjArray {
	o := C.CRoot_Tree_GetListOfBranches(t.t)
	return ObjArray{c: o}
}

func (t *Tree) GetListOfLeaves() ObjArray {
	o := C.CRoot_Tree_GetListOfLeaves(t.t)
	return ObjArray{c: o}
}

func (t *Tree) GetSelectedRows() int64 {
	return int64(C.CRoot_Tree_GetSelectedRows(t.t))
}

func (t *Tree) GetVal(i int) []float64 {
	c_data := C.CRoot_Tree_GetVal(t.t, C.int32_t(i))
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *Tree) GetV1() []float64 {
	c_data := C.CRoot_Tree_GetV1(t.t)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *Tree) GetV2() []float64 {
	c_data := C.CRoot_Tree_GetV2(t.t)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j!=sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data,C.int(j)))
	}
	return d
}

func (t *Tree) GetV3() []float64 {
	c_data := C.CRoot_Tree_GetV3(t.t)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *Tree) GetV4() []float64 {
	c_data := C.CRoot_Tree_GetV4(t.t)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *Tree) GetW() []float64 {
	c_data := C.CRoot_Tree_GetW(t.t)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *Tree) LoadTree(entry int64) int64 {
	return int64(C.CRoot_Tree_LoadTree(t.t, C.int64_t(entry)))
}

// func (t *Tree) MakeClass
// func (t *Tree) Notify

func (t *Tree) Print(option string) {
	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_Tree_Print(t.t, (*C.CRoot_Option)(c_option))
}

func (t *Tree) SetBranchAddress(name string, obj interface{}) int32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	v := reflect.Indirect(reflect.ValueOf(obj))
	br := gobranch{g: reflect.ValueOf(obj), c: ffi.ValueOf(v.Interface())}
	br.dec = ffi.NewDecoder(br.c)
	t.branches[name] = br

	genreflex(br.c.Type())
	c_addr := br.c.UnsafeAddr()
	//fmt.Printf("setBranch(%s, %v)...\n", name, c_addr)
	rc := C.CRoot_Tree_SetBranchAddress(t.t, c_name, unsafe.Pointer(c_addr), nil)
	return int32(rc)
}

func (t *Tree) SetBranchStatus(name string, status bool) uint32 {
	c_found := C.uint32_t(0)
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.CRoot_Tree_SetBranchStatus(t.t, c_name, bool2c(status), &c_found)
	return uint32(c_found)
}

func (t *Tree) Write(name string, option, bufsize int) int {
	if len(name) != 0 {
		c_name := C.CString(name)
		defer C.free(unsafe.Pointer(c_name))
		return int(C.CRoot_Tree_Write(t.t, c_name, C.int32_t(option), C.int32_t(bufsize)))
	}
	c_name := (*C.char)(unsafe.Pointer(nil))
	return int(C.CRoot_Tree_Write(t.t, c_name, C.int32_t(option), C.int32_t(bufsize)))
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

func (r *Random) Rannorf() (a, b float32) {
	c_a := (*C.float)(unsafe.Pointer(&a))
	c_b := (*C.float)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannorf(r.r, c_a, c_b)
	return
}

func (r *Random) Rannord() (a, b float64) {
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
	GRandom = &Random{r: C.CRoot_gRandom}
	fmt.Printf(":: initializing go-croot...[done]\n")
}

// EOF
