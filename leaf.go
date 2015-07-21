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
	SetAddress(addr unsafe.Pointer)
}

type leafImpl struct {
	c C.CRoot_Leaf
}

func (l *leafImpl) GetBranch() Branch {
	c := C.CRoot_Leaf_GetBranch(l.c)
	if c == nil {
		return nil
	}
	return &branchImpl{c: c}
}

func (l *leafImpl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.c))
}

func (l *leafImpl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.c)
	obj := objectImpl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leafImpl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.c)
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leafImpl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.c)
	return uintptr(ptr)
}

func (l *leafImpl) SetAddress(addr unsafe.Pointer) {
	C.CRoot_Leaf_SetAddress(l.c, addr)
}

// LeafI
type LeafI interface {
	Leaf
	GetValue(idx int) float64
}

type leafIImpl struct {
	c C.CRoot_LeafI
}

func (l *leafIImpl) GetValue(idx int) float64 {
	o := C.CRoot_LeafI_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leafIImpl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leafIImpl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := objectImpl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leafIImpl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leafIImpl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leafIImpl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}

// LeafF
type LeafF interface {
	Leaf
	GetValue(idx int) float64
}

type leafFImpl struct {
	c C.CRoot_LeafF
}

func (l *leafFImpl) GetValue(idx int) float64 {
	o := C.CRoot_LeafF_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leafFImpl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leafFImpl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := objectImpl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leafFImpl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leafFImpl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leafFImpl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}

// LeafD
type LeafD interface {
	Leaf
	GetValue(idx int) float64
}

type leafDImpl struct {
	c C.CRoot_LeafD
}

func (l *leafDImpl) GetValue(idx int) float64 {
	o := C.CRoot_LeafD_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leafDImpl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leafDImpl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := objectImpl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leafDImpl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leafDImpl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leafDImpl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}

// LeafO
type LeafO interface {
	Leaf
	GetValue(idx int) float64
}

type leafOImpl struct {
	c C.CRoot_LeafO
}

func (l *leafOImpl) GetValue(idx int) float64 {
	o := C.CRoot_LeafO_GetValue(l.c, C.int(idx))
	return float64(o)
}

func (l *leafOImpl) GetLenStatic() int {
	return int(C.CRoot_Leaf_GetLenStatic(l.as_tleaf()))
}

func (l *leafOImpl) GetLeafCount() Leaf {
	c := C.CRoot_Leaf_GetLeafCount(l.as_tleaf())
	obj := objectImpl{c: (C.CRoot_Object)(c)}
	return to_gocroot(&obj).(Leaf)
}

func (l *leafOImpl) GetTypeName() string {
	c_str := C.CRoot_Leaf_GetTypeName(l.as_tleaf())
	// we do NOT own c_str
	// defer C.free(unsafe.Point(c_str))
	return C.GoString(c_str)
}

func (l *leafOImpl) GetValuePointer() uintptr {
	ptr := C.CRoot_Leaf_GetValuePointer(l.as_tleaf())
	return uintptr(ptr)
}

func (l *leafOImpl) as_tleaf() C.CRoot_Leaf {
	return (C.CRoot_Leaf)(unsafe.Pointer(l.c))
}
