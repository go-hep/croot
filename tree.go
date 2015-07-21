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
	"log"
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

type treeImpl struct {
	c        C.CRoot_Tree
	branches map[string]*gobranch
}

type gobranch struct {
	v     reflect.Value  // pointer to go-value
	c     cmem.Value     // pointer to C-value
	cptr  unsafe.Pointer // pointer to C-value buffer
	addr  unsafe.Pointer // address of that C-value buffer
	valid bool           // whether the branch has been correctly connected to the Tree C-buffer
	br    *branchImpl
}

func (br *gobranch) get_c_branch(t *treeImpl, name string) unsafe.Pointer {
	//fmt.Printf("::: get_c_branch...\n")
	var ptr unsafe.Pointer

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	// search for branch first.

	cbr := C.CRoot_Tree_GetBranch(t.c, cname)
	if cbr == nil {
		//fmt.Printf("==[%s]... sub-branch ?\n", name)
		// sub-branch ? the actual name may have a trailing '.'
		cname := C.CString(name + ".")
		defer C.free(unsafe.Pointer(cname))
		cbr = C.CRoot_Tree_GetBranch(t.c, cname)
	}

	// found a branched object.
	if cbr != nil {
		//fmt.Printf("==[%s]... branched object...\n", name)
		cls_name := t.GetBranch(name).GetClassName()
		cls := GetClass(cls_name)
		//fmt.Printf(">>> [%v] -> class=%q %v\n", name, cls_name, cls)
		if cls != nil {
			//fmt.Printf("==[%s]... TBranch::GetAddress()...\n", name)
			addr := unsafe.Pointer(C.CRoot_Branch_GetAddress(cbr))
			ptr = *(*unsafe.Pointer)(addr)
			return ptr
		}
	}

	// try leaf.
	//fmt.Printf("==[%s]... TTree::GetLeaf()...\n", name)
	cleaf := C.CRoot_Tree_GetLeaf(t.c, cname)
	if cbr != nil && cleaf == nil {
		//fmt.Printf("==[%s]... TBranch::GetLeaf()...\n", name)
		//fmt.Printf("==[%s] cbr=%p cleaf=%v\n", name, cbr, cleaf)
		cleaf = C.CRoot_Branch_GetLeaf(cbr, cname)
		if cleaf == nil {
			//fmt.Printf("==[%s]... TBranch::GetListOfLeaves()...\n", name)
			cleaves := C.CRoot_Branch_GetListOfLeaves(cbr)
			nleaves := int(C.CRoot_ObjArray_GetSize(cleaves))
			//fmt.Printf("==[%s]... n-leaves=%d...\n", name, nleaves)
			if nleaves == 1 {
				cleaf = (C.CRoot_Leaf)(unsafe.Pointer(C.CRoot_ObjArray_At(cleaves, 0)))
			} else if nleaves > 1 {
				//fmt.Fprintf(os.Stderr, "**warn** requested branch [%s] has more than one leaf. picking the first one!\n", name)
				cleaf = (C.CRoot_Leaf)(unsafe.Pointer(C.CRoot_ObjArray_At(cleaves, 0)))
			}
		}
	}

	// found a leaf, extract value
	if cleaf != nil {
		// array types.
		clenstatic := int(C.CRoot_Leaf_GetLenStatic(cleaf))
		cleafcount := C.CRoot_Leaf_GetLeafCount(cleaf)
		//fmt.Printf("==[%s]... TLeaf::GetLenStatic() = %d...\n", name, clenstatic)
		if 1 < clenstatic || cleafcount != nil {
			ptr = unsafe.Pointer(C.CRoot_Leaf_GetValuePointer(cleaf))
			return ptr
		}

		// value types
		ptr = unsafe.Pointer(C.CRoot_Leaf_GetValuePointer(cleaf))
		return ptr
	}

	log.Panicf("==[%s] err... utter confusion!!\n", name)
	return ptr
}

func (br *gobranch) update_from_c(t *treeImpl, name string) error {
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
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	t := C.CRoot_Tree_new(cname, ctitle, C.int32_t(splitlevel))
	b := make(map[string]*gobranch)
	return &treeImpl{c: t, branches: b}
}

func (t *treeImpl) Delete() {
	C.CRoot_Tree_delete(t.c)
	t.c = nil
	t.branches = nil
}

func (t *treeImpl) Branch(name string, obj interface{}, bufsiz, splitlevel int) (Branch, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

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
	gendict(br.v.Type())

	br.cptr = unsafe.Pointer(br.c.UnsafeAddr())
	br.addr = unsafe.Pointer(&br.cptr)

	classname := toCxxName(val.Type())
	ccls := C.CString(classname)
	defer C.free(unsafe.Pointer(ccls))

	b := C.CRoot_Tree_Branch(t.c, cname, ccls, br.addr, C.int32_t(bufsiz), C.int32_t(splitlevel))
	br.br = &branchImpl{c: b}
	t.branches[name] = br

	return br.br, nil
}

func (t *treeImpl) Branch2(name string, objaddr interface{}, leaflist string, bufsiz int) (Branch, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ptr := reflect.ValueOf(objaddr)
	if ptr.Type().Kind() != reflect.Ptr {
		return nil, fmt.Errorf("croot.Tree.Branch2: takes a pointer to a builtin (got %v)", ptr.Type())
	}
	val := reflect.Indirect(ptr)
	switch k := val.Type().Kind(); k {
	default:
		// ok.
	case reflect.Ptr, reflect.Struct, reflect.String,
		reflect.Slice:
		return nil, fmt.Errorf("croot.Tree.Branch2: takes a pointer to a builtin (got %v)", ptr.Type())
	}
	br := &gobranch{v: val, c: cmem.ValueOf(val.Interface())}
	// register the type with Reflex
	gendict(br.v.Type())

	br.cptr = unsafe.Pointer(br.c.UnsafeAddr())
	br.addr = unsafe.Pointer(br.cptr)

	cleaflist := C.CString(leaflist)
	defer C.free(unsafe.Pointer(cleaflist))

	b := C.CRoot_Tree_Branch2(t.c, cname, br.addr, cleaflist, C.int32_t(bufsiz))
	br.br = &branchImpl{c: b}
	t.branches[name] = br
	return br.br, nil
}

func (t *treeImpl) Fill() (int, error) {
	// fmt.Printf("=== fill ===...\n")
	for _, br := range t.branches {
		br.c.SetValue(br.v)
	}
	nb := int(C.CRoot_Tree_Fill(t.c))
	// fmt.Printf("=== fill ===... [done]\n")
	if nb < 0 {
		return nb, fmt.Errorf("croot.Tree.Fill: error")
	}
	return nb, nil
}

func (t *treeImpl) GetBranch(name string) Branch {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	b := C.CRoot_Tree_GetBranch(t.c, cname)
	if b == nil {
		return nil
	}
	return &branchImpl{c: b}
}

func (t *treeImpl) GetEntries() int64 {
	return int64(C.CRoot_Tree_GetEntries(t.c))
}

func (t *treeImpl) GetEntry(entry int64, getall int) int {
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

func (t *treeImpl) GetLeaf(name string) Leaf {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.CRoot_Tree_GetLeaf(t.c, cname)
	if c == nil {
		return nil
	}
	return &leafImpl{c: c}
}

func (t *treeImpl) GetListOfBranches() []Branch {
	c := C.CRoot_Tree_GetListOfBranches(t.c)
	objs := objArrayImpl{c: c}
	branches := make([]Branch, objs.GetEntries())
	for i := 0; i < len(branches); i++ {
		obj := objs.At(int64(i))
		br := t.GetBranch(obj.GetName())
		branches[i] = br
	}
	return branches
}

func (t *treeImpl) GetListOfLeaves() []Leaf {
	c := C.CRoot_Tree_GetListOfLeaves(t.c)
	objs := objArrayImpl{c: c}
	leaves := make([]Leaf, objs.GetEntries())
	for i := 0; i < len(leaves); i++ {
		obj := objs.At(int64(i))
		leaf := t.GetLeaf(obj.GetName())
		leaves[i] = leaf
	}
	return leaves
}

func (t *treeImpl) GetSelectedRows() int64 {
	return int64(C.CRoot_Tree_GetSelectedRows(t.c))
}

func (t *treeImpl) GetVal(i int) []float64 {
	cdata := C.CRoot_Tree_GetVal(t.c, C.int32_t(i))
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(cdata, C.int(j)))
	}
	return d
}

func (t *treeImpl) GetV1() []float64 {
	cdata := C.CRoot_Tree_GetV1(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(cdata, C.int(j)))
	}
	return d
}

func (t *treeImpl) GetV2() []float64 {
	cdata := C.CRoot_Tree_GetV2(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(cdata, C.int(j)))
	}
	return d
}

func (t *treeImpl) GetV3() []float64 {
	cdata := C.CRoot_Tree_GetV3(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(cdata, C.int(j)))
	}
	return d
}

func (t *treeImpl) GetV4() []float64 {
	cdata := C.CRoot_Tree_GetV4(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(cdata, C.int(j)))
	}
	return d
}

func (t *treeImpl) GetW() []float64 {
	cdata := C.CRoot_Tree_GetW(t.c)
	sz := t.GetSelectedRows()
	d := make([]float64, sz)
	for j := int64(0); j != sz; sz++ {
		d[j] = float64(C._go_croot_double_at(cdata, C.int(j)))
	}
	return d
}

func (t *treeImpl) LoadTree(entry int64) int64 {
	return int64(C.CRoot_Tree_LoadTree(t.c, C.int64_t(entry)))
}

// func (t *treeImpl) MakeClass
// func (t *treeImpl) Notify

func (t *treeImpl) SetBranchAddress(name string, obj interface{}) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

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
	// register the type with ROOT's reflection type system
	gendict(typ)

	br.c = cmem.ValueOf(val.Interface())
	br.cptr = unsafe.Pointer(br.c.UnsafeAddr())
	br.addr = unsafe.Pointer(&br.cptr)
	//

	rc := C.CRoot_Tree_SetBranchAddress(t.c, cname, br.addr, nil)

	if t.branches == nil {
		t.branches = make(map[string]*gobranch)
	}

	t.branches[name] = br
	return int32(rc)
}

func (t *treeImpl) SetBranchStatus(name string, status bool) uint32 {
	cfound := C.uint32_t(0)
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.CRoot_Tree_SetBranchStatus(t.c, cname, bool2c(status), &cfound)
	return uint32(cfound)
}

func (t *treeImpl) Write(name string, option, bufsize int) int {
	if len(name) != 0 {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))
		return int(C.CRoot_Tree_Write(t.c, cname, C.int32_t(option), C.int32_t(bufsize)))
	}
	cname := (*C.char)(unsafe.Pointer(nil))
	return int(C.CRoot_Tree_Write(t.c, cname, C.int32_t(option), C.int32_t(bufsize)))
}
