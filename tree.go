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
	"fmt"
	"reflect"
	"unsafe"
)

// Tree
type Tree interface {
	Object

	Branch(name string, obj interface{}, bufsiz, splitlevel int) (Branch, error)
	Branch2(name string, objaddr interface{}, leaflist string, bufsiz int) (Branch, error)
	Delete()
	Fill() (int, error)
	GetBranch(name string) Branch
	GetEntries() int64
	GetEntry(entry int64, getall int) int
	GetLeaf(name string) Leaf
	GetListOfBranches() ObjArray
	GetListOfLeaves() ObjArray
	GetSelectedRows() int64
	GetVal(i int) []float64
	GetV1() []float64
	GetV2() []float64
	GetV3() []float64
	GetV4() []float64
	GetW() []float64
	LoadTree(entry int64) int64
	SetBranchAddress(name string, obj interface{}) int32
	SetBranchStatus(name string, status bool) uint32
	Write(name string, option, bufsize int) int
}

type tree_impl struct {
	c        C.CRoot_Tree
	branches map[string]gobranch
}

func (t *tree_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(t.c)
}

func (t *tree_impl) as_tobject() *object_impl {
	return &object_impl{t.cptr()}
}

func (t *tree_impl) ClassName() string {
	return t.as_tobject().ClassName()
}

func (t *tree_impl) Clone(opt Option) Object {
	return t.as_tobject().Clone(opt)
}

func (t *tree_impl) FindObject(name string) Object {
	return t.as_tobject().FindObject(name)
}

func (t *tree_impl) GetName() string {
	return t.as_tobject().GetName()
}

func (t *tree_impl) GetTitle() string {
	return t.as_tobject().GetTitle()
}

func (t *tree_impl) InheritsFrom(clsname string) bool {
	return t.as_tobject().InheritsFrom(clsname)
}

type gobranch struct {
	v    reflect.Value  // pointer to go-value
	buf  uintptr        // pointer to go-value buffer
	addr unsafe.Pointer // address of that go-value buffer
	br   *branch_impl
}

// func (br gobranch) update_from_go() {
// 	br.c.SetValue(br.g.Elem())
// }

func NewTree(name, title string, splitlevel int) Tree {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	t := C.CRoot_Tree_new(c_name, c_title, C.int32_t(splitlevel))
	b := make(map[string]gobranch)
	return &tree_impl{c: t, branches: b}
}

func (t *tree_impl) Delete() {
	C.CRoot_Tree_delete(t.c)
	t.c = nil
	t.branches = nil
}

func (t *tree_impl) Branch(name string, obj interface{}, bufsiz, splitlevel int) (Branch, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	ptr := reflect.ValueOf(obj)
	if ptr.Type().Kind() != reflect.Ptr {
		return nil, fmt.Errorf("croot.Tree.Branch: takes a pointer to a struct (got %v)", ptr.Type())
	}
	val := reflect.Indirect(ptr)
	if val.Type().Kind() != reflect.Struct {
		return nil, fmt.Errorf("croot.Tree.Branch: takes a pointer to a struct (got %v)", ptr.Type())
	}
	br := gobranch{v: val}
	// register the type with Reflex
	genreflex(br.v.Type())

	// this scaffolding is needed b/c we need to keep the UnsafeAddr alive.
	br.buf = br.v.UnsafeAddr()
	br.addr = unsafe.Pointer(&br.buf)
	//

	classname := to_cxx_name(val.Type())
	c_classname := C.CString(classname)
	defer C.free(unsafe.Pointer(c_classname))

	b := C.CRoot_Tree_Branch(t.c, c_name, c_classname, br.addr, C.int32_t(bufsiz), C.int32_t(splitlevel))
	br.br = &branch_impl{c: b}
	t.branches[name] = br
	return br.br, nil
}

func (t *tree_impl) Branch2(name string, objaddr interface{}, leaflist string, bufsiz int) (Branch, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	ptr := reflect.ValueOf(objaddr)
	if ptr.Type().Kind() != reflect.Ptr {
		return nil, fmt.Errorf("croot.Tree.Branch: takes a pointer to a builtin (got %v)", ptr.Type())
	}
	val := reflect.Indirect(ptr)
	switch k := val.Type().Kind(); k {
	default:
		// ok.
	case reflect.Ptr, reflect.Struct, reflect.String, reflect.Array,
		reflect.Slice:
		return nil, fmt.Errorf("croot.Tree.Branch: takes a pointer to a builtin (got %v)", ptr.Type())
	}
	br := gobranch{v: val}
	// register the type with Reflex
	genreflex(br.v.Type())

	// this scaffolding is needed b/c we need to keep the UnsafeAddr alive.
	br.buf = br.v.UnsafeAddr()
	br.addr = unsafe.Pointer(br.buf)
	//

	c_leaflist := C.CString(leaflist)
	defer C.free(unsafe.Pointer(c_leaflist))

	b := C.CRoot_Tree_Branch2(t.c, c_name, br.addr, c_leaflist, C.int32_t(bufsiz))
	br.br = &branch_impl{c: b}
	t.branches[name] = br
	return br.br, nil
}

func (t *tree_impl) Fill() (int, error) {
	// fmt.Printf("=== fill ===...\n")
	// for n, v := range t.branches {
	// 	fmt.Printf("branch[%s]: %v (%p)\n", n, v.v.Interface(),
	// 		unsafe.Pointer(v.v.UnsafeAddr()))
	// 	if n == "evt" {
	// 		vv := v.v.Field(1).Field(2)
	// 		fmt.Printf("   evt.A.Fs: %d %v (%p)\n", vv.Len(), vv.Interface(),
	// 			unsafe.Pointer(vv.UnsafeAddr()))
	// 	}
	// }
	nb := int(C.CRoot_Tree_Fill(t.c))
	// fmt.Printf("--addrs: %d\n", len(addrs))
	// if o > 0 {
	// 	addrs = addrs[:0]
	// }
	// fmt.Printf("=== fill ===... [done]\n")
	if nb < 0 {
		return nb, fmt.Errorf("croot.Tree.Fill: error")
	}
	return nb, nil
}

func (t *tree_impl) GetBranch(name string) Branch {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	b := C.CRoot_Tree_GetBranch(t.c, c_name)
	return &branch_impl{c: b}
}

func (t *tree_impl) GetEntries() int64 {
	return int64(C.CRoot_Tree_GetEntries(t.c))
}

func (t *tree_impl) GetEntry(entry int64, getall int) int {
	nbytes := C.CRoot_Tree_GetEntry(t.c, C.int64_t(entry), C.int32_t(getall))
	// if nbytes > 0 {
	// 	for _, br := range t.branches {
	// 		if br.v.IsValid() {
	// 			br.v.Elem().Set(br.c.GoValue())
	// 		}
	// 	}
	// }
	return int(nbytes)
}

func (t *tree_impl) GetLeaf(name string) Leaf {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	l := C.CRoot_Tree_GetLeaf(t.c, c_name)
	return &leaf_impl{c: l}
}

func (t *tree_impl) GetListOfBranches() ObjArray {
	o := C.CRoot_Tree_GetListOfBranches(t.c)
	return &objarray_impl{c: o}
}

func (t *tree_impl) GetListOfLeaves() ObjArray {
	o := C.CRoot_Tree_GetListOfLeaves(t.c)
	return &objarray_impl{c: o}
}

func (t *tree_impl) GetSelectedRows() int64 {
	return int64(C.CRoot_Tree_GetSelectedRows(t.c))
}

func (t *tree_impl) GetVal(i int) []float64 {
	c_data := C.CRoot_Tree_GetVal(t.c, C.int32_t(i))
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *tree_impl) GetV1() []float64 {
	c_data := C.CRoot_Tree_GetV1(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *tree_impl) GetV2() []float64 {
	c_data := C.CRoot_Tree_GetV2(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *tree_impl) GetV3() []float64 {
	c_data := C.CRoot_Tree_GetV3(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *tree_impl) GetV4() []float64 {
	c_data := C.CRoot_Tree_GetV4(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *tree_impl) GetW() []float64 {
	c_data := C.CRoot_Tree_GetW(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(c_data, C.int(j)))
	}
	return d
}

func (t *tree_impl) LoadTree(entry int64) int64 {
	return int64(C.CRoot_Tree_LoadTree(t.c, C.int64_t(entry)))
}

// func (t *tree_impl) MakeClass
// func (t *tree_impl) Notify

func (t *tree_impl) Print(option Option) {
	c_option := C.CString(string(option))
	defer C.free(unsafe.Pointer(c_option))

	C.CRoot_Tree_Print(t.c, (*C.CRoot_Option)(c_option))
}

func (t *tree_impl) SetBranchAddress(name string, obj interface{}) int32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var ptr reflect.Value
	var val reflect.Value

	switch obj.(type) {
	case reflect.Value:
		val = obj.(reflect.Value)
	default:
		ptr = reflect.ValueOf(obj)
		val = reflect.Indirect(ptr)
	}

	br := gobranch{v: val}
	typ := br.v.Type()
	// register the type with Reflex
	genreflex(typ)

	// this scaffolding is needed b/c we need to keep the UnsafeAddr alive.
	br.buf = br.v.UnsafeAddr()
	if typ.Kind() == reflect.Struct {
		br.addr = unsafe.Pointer(&br.buf)
	} else {
		br.addr = unsafe.Pointer(br.buf)
	}
	//

	rc := C.CRoot_Tree_SetBranchAddress(t.c, c_name, br.addr, nil)

	t.branches[name] = br
	return int32(rc)
}

func (t *tree_impl) SetBranchStatus(name string, status bool) uint32 {
	c_found := C.uint32_t(0)
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	C.CRoot_Tree_SetBranchStatus(t.c, c_name, bool2c(status), &c_found)
	return uint32(c_found)
}

func (t *tree_impl) Write(name string, option, bufsize int) int {
	if len(name) != 0 {
		c_name := C.CString(name)
		defer C.free(unsafe.Pointer(c_name))
		return int(C.CRoot_Tree_Write(t.c, c_name, C.int32_t(option), C.int32_t(bufsize)))
	}
	c_name := (*C.char)(unsafe.Pointer(nil))
	return int(C.CRoot_Tree_Write(t.c, c_name, C.int32_t(option), C.int32_t(bufsize)))
}

func init() {
	cnvmap["TTree"] = func(o c_object) Object {
		return &tree_impl{c: (C.CRoot_Tree)(o.cptr())}
	}
}

// EOF
