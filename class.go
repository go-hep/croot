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

func (c *classImpl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(c.c)
}

func (c *classImpl) as_tobject() *object_impl {
	return &object_impl{c.cptr()}
}

func (c *classImpl) ClassName() string {
	return c.as_tobject().ClassName()
}

func (c *classImpl) Clone(opt Option) Object {
	return c.as_tobject().Clone(opt)
}

func (c *classImpl) FindObject(name string) Object {
	return c.as_tobject().FindObject(name)
}

func (c *classImpl) GetName() string {
	return c.as_tobject().GetName()
}

func (c *classImpl) GetTitle() string {
	return c.as_tobject().GetTitle()
}

func (c *classImpl) InheritsFrom(clsname string) bool {
	return c.as_tobject().InheritsFrom(clsname)
}

func (c *classImpl) Print(option Option) {
	c.as_tobject().Print(option)
}

func GetClass(name string) Class {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c := C.CRoot_Class_GetClass(c_name)
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

func (c *dataMemberImpl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(c.c)
}

func (c *dataMemberImpl) as_tobject() *object_impl {
	return &object_impl{c.cptr()}
}

func (c *dataMemberImpl) ClassName() string {
	return c.as_tobject().ClassName()
}

func (c *dataMemberImpl) Clone(opt Option) Object {
	return c.as_tobject().Clone(opt)
}

func (c *dataMemberImpl) FindObject(name string) Object {
	return c.as_tobject().FindObject(name)
}

func (c *dataMemberImpl) GetName() string {
	return c.as_tobject().GetName()
}

func (c *dataMemberImpl) GetTitle() string {
	return c.as_tobject().GetTitle()
}

func (c *dataMemberImpl) InheritsFrom(clsname string) bool {
	return c.as_tobject().InheritsFrom(clsname)
}

func (c *dataMemberImpl) Print(option Option) {
	c.as_tobject().Print(option)
}

// EOF
