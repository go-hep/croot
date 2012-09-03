package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
//
// double _go_croot_double_at(double *array, int idx)
// { return array[idx]; }
//
import "C"

import (
	"reflect"
	"unsafe"

	"github.com/sbinet/go-ffi/pkg/ffi"
)

// Tree
type Tree struct {
	t        C.CRoot_Tree
	branches map[string]gobranch
}

type gobranch struct {
	g    reflect.Value  // pointer to go-value
	c    ffi.Value      // the equivalent C-value
	buf  uintptr        // pointer to C-value buffer
	addr unsafe.Pointer // address of that C-value buffer
	br   Branch
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

	ptr := reflect.ValueOf(obj)
	val := reflect.Indirect(ptr)
	br := gobranch{g: ptr, c: ffi.ValueOf(val.Interface())}
	ct := br.c.Type()
	if ct.GoType() == nil {
		panic("no Go-type for ffi.Type [" + ct.Name() + "] !!")
	}
	// register the type with Reflex
	genreflex(br.c.Type())

	// this scaffolding is needed b/c we need to keep the UnsafeAddr alive.
	br.buf = br.c.UnsafeAddr()
	br.addr = unsafe.Pointer(&br.buf)
	//

	classname := to_cxx_name(ct.GoType())
	c_classname := C.CString(classname)
	defer C.free(unsafe.Pointer(c_classname))

	b := C.CRoot_Tree_Branch(t.t, c_name, c_classname, br.addr, C.int32_t(bufsiz), C.int32_t(splitlevel))
	br.br = Branch{c: b}
	t.branches[name] = br
	return br.br
}

func (t *Tree) Branch2(name string, objaddr interface{}, leaflist string, bufsiz int) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	ptr := reflect.ValueOf(objaddr)
	val := reflect.Indirect(ptr)
	br := gobranch{g: ptr, c: ffi.ValueOf(val.Interface())}
	ct := br.c.Type()
	if ct.GoType() == nil {
		panic("no Go-type for ffi.Type [" + ct.Name() + "] !!")
	}
	// register the type with Reflex
	genreflex(br.c.Type())

	// this scaffolding is needed b/c we need to keep the UnsafeAddr alive.
	br.buf = br.c.UnsafeAddr()
	br.addr = unsafe.Pointer(br.buf)
	//

	c_leaflist := C.CString(leaflist)
	defer C.free(unsafe.Pointer(c_leaflist))

	b := C.CRoot_Tree_Branch2(t.t, c_name, br.addr, c_leaflist, C.int32_t(bufsiz))
	br.br = Branch{c: b}
	t.branches[name] = br
	return br.br
}

func (t *Tree) Fill() int {
	// synchronize branches: update ffi.Value
	for _, br := range t.branches {
		br.c.SetValue(br.g.Elem())
	}
	o := int(C.CRoot_Tree_Fill(t.t))
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
		for _, br := range t.branches {
			br.g.Elem().Set(br.c.GoValue())
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
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
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

	ptr := reflect.ValueOf(obj)
	val := reflect.Indirect(ptr)
	br := gobranch{g: ptr, c: ffi.ValueOf(val.Interface())}
	ct := br.c.Type()
	if ct.GoType() == nil {
		panic("no Go-type for ffi.Type [" + ct.Name() + "] !!")
	}
	// register the type with Reflex
	genreflex(br.c.Type())

	// this scaffolding is needed b/c we need to keep the UnsafeAddr alive.
	br.buf = br.c.UnsafeAddr()
	if ct.Kind() == ffi.Struct {
		br.addr = unsafe.Pointer(&br.buf)
	} else {
		br.addr = unsafe.Pointer(br.buf)
	}
	//

	rc := C.CRoot_Tree_SetBranchAddress(t.t, c_name, br.addr, nil)

	t.branches[name] = br
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

// EOF
