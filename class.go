package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
//
import "C"

import (
	"unsafe"
)

type Version int16

type Class interface {
	Object

	GetBaseClass(class string) Class
	GetClassSize() int
	GetClassVersion() Version
	GetDataMember(name string) DataMember
}

// DataMember provides information about name of data member, its type, and
// comment field string.
type DataMember interface {
	Object

	/*
		GetArrayDim() int
		GetClass() Class
		GetDataType() DataType
		GetFullTypeName() string
		GetOffset() int
		GetTypeName() string

		IsaPointer() bool
		IsBasic() bool
		IsEnum() bool
		IsPersistent() bool
		IsSTLContainer() STLContainerType
	*/
}

// FIXME(sbinet) defined in TDictionary (kList, kVector, ...)
type STLContainerType int

// DataType provides basic data type description.
type DataType int

type classImpl struct {
	c C.CRoot_Class
}

func GetClass(name string) Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	c := C.CRoot_Class_GetClass(cname)
	if c == nil {
		return nil
	}
	return &classImpl{c: c}
}

func (c *classImpl) GetBaseClass(name string) Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	base := C.CRoot_Class_GetBaseClass(c.c, cname)
	if base == nil {
		return nil
	}

	return &classImpl{c: base}
}

func (c *classImpl) GetClassSize() int {
	return int(C.CRoot_Class_GetClassSize(c.c))
}

func (c *classImpl) GetClassVersion() Version {
	return Version(C.CRoot_Class_GetClassVersion(c.c))
}

func (c *classImpl) GetDataMember(name string) DataMember {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	mbr := C.CRoot_Class_GetDataMember(c.c, cname)
	if mbr == nil {
		return nil
	}

	return &dataMemberImpl{c: mbr}
}

type dataMemberImpl struct {
	c C.CRoot_DataMember
}

// EOF
