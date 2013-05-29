package croot

// #include "croot/croot.h"
//
import "C"

import (
//"unsafe"
)

// Leaf
type Leaf interface {
	Object
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

func init() {
	cnvmap["TLeaf"] = func(o c_object) Object {
		return &leaf_impl{c: (C.CRoot_Leaf)(o.cptr())}
	}
}

// EOF
