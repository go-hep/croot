package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
// #include <stdio.h>
//
// double _go_croot_double_at(double *array, int idx)
// { return array[idx]; }
//
// char* _go_croot_new_string(void* src, int len)
// {
//   char *dst = (char*)malloc((len+1) * sizeof(char));
//   if (dst == NULL) { return NULL; }
//   dst = strncpy(dst, src, len);
//   dst[len] = '\0';
//   return dst;
// }
import "C"

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	"github.com/go-hep/croot/cmem"
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
	GetListOfBranches() []Branch
	GetListOfLeaves() []Leaf
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
	branches map[string]*gobranch
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
	v     reflect.Value  // pointer to go-value
	c     cmem.Value     // pointer to C-value
	cptr  unsafe.Pointer // pointer to C-value buffer
	addr  unsafe.Pointer // address of that C-value buffer
	valid bool           // whether the branch has been correctly connected to the Tree C-buffer
	br    *branch_impl
}

func (br *gobranch) get_c_branch(t *tree_impl, name string) unsafe.Pointer {
	//fmt.Printf("::: get_c_branch...\n")
	var ptr unsafe.Pointer

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	// search for branch first.

	c_br := C.CRoot_Tree_GetBranch(t.c, c_name)
	if c_br == nil {
		//fmt.Printf("==[%s]... sub-branch ?\n", name)
		// sub-branch ? the actual name may have a trailing '.'
		c_name := C.CString(name + ".")
		defer C.free(unsafe.Pointer(c_name))
		c_br = C.CRoot_Tree_GetBranch(t.c, c_name)
	}

	// found a branched object.
	if c_br != nil {
		//fmt.Printf("==[%s]... branched object...\n", name)
		cls_name := t.GetBranch(name).GetClassName()
		cls := GetClass(cls_name)
		//fmt.Printf(">>> [%v] -> class=%q %v\n", name, cls_name, cls)
		if cls != nil {
			//fmt.Printf("==[%s]... TBranch::GetAddress()...\n", name)
			addr := unsafe.Pointer(C.CRoot_Branch_GetAddress(c_br))
			ptr = *(*unsafe.Pointer)(addr)
			return ptr
		}
	}

	// try leaf.
	//fmt.Printf("==[%s]... TTree::GetLeaf()...\n", name)
	c_leaf := C.CRoot_Tree_GetLeaf(t.c, c_name)
	if c_br != nil && c_leaf == nil {
		//fmt.Printf("==[%s]... TBranch::GetLeaf()...\n", name)
		//fmt.Printf("==[%s] c_br=%p c_leaf=%v\n", name, c_br, c_leaf)
		c_leaf = C.CRoot_Branch_GetLeaf(c_br, c_name)
		if c_leaf == nil {
			//fmt.Printf("==[%s]... TBranch::GetListOfLeaves()...\n", name)
			c_leaves := C.CRoot_Branch_GetListOfLeaves(c_br)
			nleaves := int(C.CRoot_ObjArray_GetSize(c_leaves))
			//fmt.Printf("==[%s]... n-leaves=%d...\n", name, nleaves)
			if nleaves == 1 {
				c_leaf = (C.CRoot_Leaf)(unsafe.Pointer(C.CRoot_ObjArray_At(c_leaves, 0)))
			} else if nleaves > 1 {
				//fmt.Fprintf(os.Stderr, "**warn** requested branch [%s] has more than one leaf. picking the first one!\n", name)
				c_leaf = (C.CRoot_Leaf)(unsafe.Pointer(C.CRoot_ObjArray_At(c_leaves, 0)))
			}
		}
	}

	// found a leaf, extract value
	if c_leaf != nil {
		// array types.
		c_len_static := int(C.CRoot_Leaf_GetLenStatic(c_leaf))
		c_leaf_count := C.CRoot_Leaf_GetLeafCount(c_leaf)
		//fmt.Printf("==[%s]... TLeaf::GetLenStatic() = %d...\n", name, c_len_static)
		if 1 < c_len_static || c_leaf_count != nil {
			// c_n_data := int(C.CRoot_Leaf_GetNdata(c_leaf))
			ptr = unsafe.Pointer(C.CRoot_Leaf_GetValuePointer(c_leaf))
			panic("not implemented")
			return ptr
		}

		// value types
		ptr = unsafe.Pointer(C.CRoot_Leaf_GetValuePointer(c_leaf))
		return ptr
	}

	fmt.Printf("==[%s] err... utter confusion!!\n", name)
	panic("boo")
	return ptr
}

func (br *gobranch) update_from_c(t *tree_impl, name string) error {
	if !br.v.IsValid() {
		return fmt.Errorf("croot.update_from_c: invalid branch [%v]", name)
	}

	if !br.valid {
		//fmt.Printf(">>> br.c=%v (%v)\n", br.c.UnsafeAddr(), name)
		ptr := br.get_c_branch(t, name)
		br.c = cmem.NewAt(br.c.Type(), ptr)
		//fmt.Printf(">>> br.c=%v\n", br.c.GoValue().Interface())
		br.valid = true
	}
	if br.c.UnsafeAddr() == 0 {
		return fmt.Errorf(
			"croot.update_from_c: NULL C-pointer for branch [%s]",
			name,
		)
	}

	br.v.Set(br.c.GoValue())
	return nil
}

func NewTree(name, title string, splitlevel int) Tree {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	t := C.CRoot_Tree_new(c_name, c_title, C.int32_t(splitlevel))
	b := make(map[string]*gobranch)
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
	br := &gobranch{v: val, c: cmem.ValueOf(val.Interface())}
	// register the type with Reflex
	genreflex(br.v.Type())

	br.cptr = unsafe.Pointer(br.c.UnsafeAddr())
	br.addr = unsafe.Pointer(&br.cptr)

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
	br := &gobranch{v: val, c: cmem.ValueOf(val.Interface())}
	// register the type with Reflex
	genreflex(br.v.Type())

	br.cptr = unsafe.Pointer(br.c.UnsafeAddr())
	br.addr = unsafe.Pointer(br.cptr)

	c_leaflist := C.CString(leaflist)
	defer C.free(unsafe.Pointer(c_leaflist))

	b := C.CRoot_Tree_Branch2(t.c, c_name, br.addr, c_leaflist, C.int32_t(bufsiz))
	br.br = &branch_impl{c: b}
	t.branches[name] = br
	return br.br, nil
}

func (t *tree_impl) Fill() (int, error) {
	// fmt.Printf("=== fill ===...\n")
	for _, br := range t.branches {
		// fmt.Printf("branch[%s]: %v (%p) (%p)\n", n, br.v.Interface(),
		// 	unsafe.Pointer(br.v.UnsafeAddr()),
		// 	unsafe.Pointer(br.c.UnsafeAddr()))
		br.c.SetValue(br.v)
		// fmt.Printf("branch[%s]: %v (%p) (%p) --> %v\n", n, br.v.Interface(),
		// 	unsafe.Pointer(br.v.UnsafeAddr()),
		// 	unsafe.Pointer(br.c.UnsafeAddr()),
		// 	br.c.GoValue().Interface())
		// if br.v.Type().Name() == "DataString" {
		// 	type cgo_string struct {
		// 		Len  int32
		// 		Data unsafe.Pointer
		// 	}
		// 	cstr := (*cgo_string)(unsafe.Pointer(br.c.UnsafeAddr() + unsafe.Sizeof(int64(0)) + unsafe.Sizeof(float64(0))))
		// 	fmt.Printf("cstr-len: %v\n", cstr.Len)
		// 	fmt.Printf("cstr-ptr: 0x%x\n", cstr.Data)
		// 	gostr := C.GoStringN((*C.char)(cstr.Data), C.int(cstr.Len))
		// 	fmt.Printf("gostr:    %q\n", gostr)
		// 	ptr := br.get_c_branch(t, n)
		// 	fmt.Printf("br-ptr:  0x%x\n", ptr)
		// }
	}
	nb := int(C.CRoot_Tree_Fill(t.c))
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
	if b == nil {
		return nil
	}
	return &branch_impl{c: b}
}

func (t *tree_impl) GetEntries() int64 {
	return int64(C.CRoot_Tree_GetEntries(t.c))
}

func (t *tree_impl) GetEntry(entry int64, getall int) int {
	//fmt.Fprintf(os.Stderr, ">> GetEntry(%v, %v)...\n", entry, getall)
	nbytes := C.CRoot_Tree_GetEntry(t.c, C.int64_t(entry), C.int32_t(getall))
	if nbytes > 0 {
		for nn, br := range t.branches {
			err := br.update_from_c(t, nn)
			if err != nil {
				fmt.Fprintf(os.Stderr, "**error: %v\n", err)
				return -1
			}
		}
	}
	return int(nbytes)
}

func (t *tree_impl) GetLeaf(name string) Leaf {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c := C.CRoot_Tree_GetLeaf(t.c, c_name)
	if c == nil {
		return nil
	}
	return &leaf_impl{c: c}
}

func (t *tree_impl) GetListOfBranches() []Branch {
	c := C.CRoot_Tree_GetListOfBranches(t.c)
	objs := objarray_impl{c: c}
	branches := make([]Branch, objs.GetEntries())
	for i := 0; i < len(branches); i++ {
		obj := objs.At(int64(i))
		br := t.GetBranch(obj.GetName())
		branches[i] = br
	}
	return branches
}

func (t *tree_impl) GetListOfLeaves() []Leaf {
	c := C.CRoot_Tree_GetListOfLeaves(t.c)
	objs := objarray_impl{c: c}
	leaves := make([]Leaf, objs.GetEntries())
	for i := 0; i < len(leaves); i++ {
		obj := objs.At(int64(i))
		leaf := t.GetLeaf(obj.GetName())
		leaves[i] = leaf
	}
	return leaves
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

	br := &gobranch{v: val}
	typ := br.v.Type()
	// register the type with Reflex
	genreflex(typ)

	br.c = cmem.ValueOf(val.Interface())
	br.cptr = unsafe.Pointer(br.c.UnsafeAddr())
	br.addr = unsafe.Pointer(&br.cptr)
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
