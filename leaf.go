package croot

// #include "croot/croot.h"
//
import "C"

import (
	"unsafe"
)

// Leaf
type Leaf interface {
	Object
	GetBranch() Branch
	GetLenStatic() int
	GetLeafCount() Leaf
	GetTypeName() string
	GetValuePointer() uintptr
}

type leaf_impl struct {
	c C.CRoot_Leaf
}

func (l *leaf_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(l.c)
}

func (l *leaf_impl) as_tobject() *object_impl {
	return &object_impl{l.cptr()}
}

func (l *leaf_impl) ClassName() string {
	return l.as_tobject().ClassName()
}

func (l *leaf_impl) Clone(opt Option) Object {
	return l.as_tobject().Clone(opt)
}

func (l *leaf_impl) FindObject(name string) Object {
	return l.as_tobject().FindObject(name)
}

func (l *leaf_impl) GetBranch() Branch {
	c := C.CRoot_Leaf_GetBranch(l.c)
	if c == nil {
		return nil
	}
	return &branch_impl{c: c}
}

func (l *leaf_impl) GetName() string {
	return l.as_tobject().GetName()
}

func (l *leaf_impl) GetTitle() string {
	return l.as_tobject().GetTitle()
}

func (l *leaf_impl) InheritsFrom(clsname string) bool {
	return l.as_tobject().InheritsFrom(clsname)
}

func (l *leaf_impl) Print(option Option) {
	l.as_tobject().Print(option)
}

func (l *leaf_impl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.c))
}

func (l *leaf_impl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.c)
	obj := object_impl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leaf_impl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.c)
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leaf_impl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.c)
	return uintptr(ptr)
}

// LeafI
type LeafI interface {
	Leaf
	GetValue(idx int) float64
}

type leaf_i_impl struct {
	c C.CRoot_LeafI
}

func (l *leaf_i_impl) GetValue(idx int) float64 {
	o := C.CRoot_LeafI_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leaf_i_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(l.c)
}

func (l *leaf_i_impl) as_tobject() *object_impl {
	return &object_impl{l.cptr()}
}

func (l *leaf_i_impl) ClassName() string {
	return l.as_tobject().ClassName()
}

func (l *leaf_i_impl) Clone(opt Option) Object {
	return l.as_tobject().Clone(opt)
}

func (l *leaf_i_impl) FindObject(name string) Object {
	return l.as_tobject().FindObject(name)
}

func (l *leaf_i_impl) GetName() string {
	return l.as_tobject().GetName()
}

func (l *leaf_i_impl) GetTitle() string {
	return l.as_tobject().GetTitle()
}

func (l *leaf_i_impl) InheritsFrom(clsname string) bool {
	return l.as_tobject().InheritsFrom(clsname)
}

func (l *leaf_i_impl) Print(option Option) {
	l.as_tobject().Print(option)
}

func (l *leaf_i_impl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leaf_i_impl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := object_impl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leaf_i_impl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leaf_i_impl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leaf_i_impl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}

// LeafF
type LeafF interface {
	Leaf
	GetValue(idx int) float64
}

type leaf_f_impl struct {
	c C.CRoot_LeafF
}

func (l *leaf_f_impl) GetValue(idx int) float64 {
	o := C.CRoot_LeafF_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leaf_f_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(l.c)
}

func (l *leaf_f_impl) as_tobject() *object_impl {
	return &object_impl{l.cptr()}
}

func (l *leaf_f_impl) ClassName() string {
	return l.as_tobject().ClassName()
}

func (l *leaf_f_impl) Clone(opt Option) Object {
	return l.as_tobject().Clone(opt)
}

func (l *leaf_f_impl) FindObject(name string) Object {
	return l.as_tobject().FindObject(name)
}

func (l *leaf_f_impl) GetName() string {
	return l.as_tobject().GetName()
}

func (l *leaf_f_impl) GetTitle() string {
	return l.as_tobject().GetTitle()
}

func (l *leaf_f_impl) InheritsFrom(clsname string) bool {
	return l.as_tobject().InheritsFrom(clsname)
}

func (l *leaf_f_impl) Print(option Option) {
	l.as_tobject().Print(option)
}

func (l *leaf_f_impl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leaf_f_impl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := object_impl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leaf_f_impl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leaf_f_impl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leaf_f_impl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}

// LeafD
type LeafD interface {
	Leaf
	GetValue(idx int) float64
}

type leaf_d_impl struct {
	c C.CRoot_LeafD
}

func (l *leaf_d_impl) GetValue(idx int) float64 {
	o := C.CRoot_LeafD_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leaf_d_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(l.c)
}

func (l *leaf_d_impl) as_tobject() *object_impl {
	return &object_impl{l.cptr()}
}

func (l *leaf_d_impl) ClassName() string {
	return l.as_tobject().ClassName()
}

func (l *leaf_d_impl) Clone(opt Option) Object {
	return l.as_tobject().Clone(opt)
}

func (l *leaf_d_impl) FindObject(name string) Object {
	return l.as_tobject().FindObject(name)
}

func (l *leaf_d_impl) GetName() string {
	return l.as_tobject().GetName()
}

func (l *leaf_d_impl) GetTitle() string {
	return l.as_tobject().GetTitle()
}

func (l *leaf_d_impl) InheritsFrom(clsname string) bool {
	return l.as_tobject().InheritsFrom(clsname)
}

func (l *leaf_d_impl) Print(option Option) {
	l.as_tobject().Print(option)
}

func (l *leaf_d_impl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leaf_d_impl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := object_impl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leaf_d_impl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leaf_d_impl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leaf_d_impl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}

func init() {
	cnvmap["TLeaf"] = func(o c_object) Object {
		return &leaf_impl{c: (C.CRoot_Leaf)(o.cptr())}
	}
	cnvmap["TLeafI"] = func(o c_object) Object {
		return &leaf_i_impl{c: (C.CRoot_LeafI)(o.cptr())}
	}
	cnvmap["TLeafF"] = func(o c_object) Object {
		return &leaf_f_impl{c: (C.CRoot_LeafF)(o.cptr())}
	}
	cnvmap["TLeafD"] = func(o c_object) Object {
		return &leaf_d_impl{c: (C.CRoot_LeafD)(o.cptr())}
	}
}

// EOF
