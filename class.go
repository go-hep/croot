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
	IsSTLContainer() STLType
}

// STLType describes STL collections and some std classes.
type STLType int

// STLType values defined in ESTLType.h
const (
	STLT_NotSTL               STLType = 0
	STLT_STLvector                    = 1
	STLT_STLlist                      = 2
	STLT_STLdeque                     = 3
	STLT_STLmap                       = 4
	STLT_STLmultimap                  = 5
	STLT_STLset                       = 6
	STLT_STLmultiset                  = 7
	STLT_STLbitset                    = 8
	STLT_STLforwardlist               = 9
	STLT_STLunorderedset              = 10
	STLT_STLunorderedmultiset         = 11
	STLT_STLunorderedmap              = 12
	STLT_STLunorderedmultimap         = 13
	STLT_STLend                       = 14
	STLT_STLany                       = 300 /* TVirtualStreamerInfo::kSTL */
	STLT_STLstring                    = 365 /* TVirtualStreamerInfo::kSTLstring */
)

// DataType provides basic data type description.
type DataType interface {
	Object

	GetFullTypeName() string
	GetType() DataTypeKind
	GetTypeName() string
	Size() int
	Property() int64
}

// DataTypeKind describes basic data types
type DataTypeKind int

// DataTypeKind defined in TDataType.h -- EDataType
const (
	DTK_Char                    DataTypeKind = 1
	DTK_UChar                                = 11
	DTK_Short                                = 2
	DTK_UShort                               = 12
	DTK_Int                                  = 3
	DTK_UInt                                 = 13
	DTK_Long                                 = 4
	DTK_ULong                                = 14
	DTK_Float                                = 5
	DTK_Double                               = 8
	DTK_Double32                             = 9
	DTK_char                                 = 10
	DTK_Bool                                 = 18
	DTK_Long64                               = 16
	DTK_ULong64                              = 17
	DTK_Other                                = -1
	DTK_NoType                               = 0
	DTK_Float16                              = 19
	DTK_Counter                              = 6
	DTK_CharStar                             = 7
	DTK_Bits                                 = 15 /* for compatibility with TStreamerInfo */
	DTK_Void                                 = 20
	DTK_DataTypeAliasUnsigned                = 21
	DTK_DataTypeAliasSignedChar              = 22

	// could add "long int" etc

	DTK_NumDataTypes DataTypeKind = 23
)

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

func (dm *dataMemberImpl) GetArrayDim() int {
	return int(C.CRoot_DataMember_GetArrayDim(dm.c))
}

func (dm *dataMemberImpl) GetClass() Class {
	cls := C.CRoot_DataMember_GetClass(dm.c)
	if cls == nil {
		return nil
	}
	return &classImpl{c: cls}
}

func (dm *dataMemberImpl) GetDataType() DataType {
	dt := C.CRoot_DataMember_GetDataType(dm.c)
	if dt == nil {
		return nil
	}
	return &dataTypeImpl{c: dt}
}

func (dm *dataMemberImpl) GetFullTypeName() string {
	cname := C.CRoot_DataMember_GetFullTypeName(dm.c)
	// we dont own cname
	// defer C.free(unsafe.Pointer(cname))
	return C.GoString(cname)
}

func (dm *dataMemberImpl) GetOffset() int {
	o := C.CRoot_DataMember_GetOffset(dm.c)
	return int(o)
}

func (dm *dataMemberImpl) GetTypeName() string {
	cname := C.CRoot_DataMember_GetTypeName(dm.c)
	// we dont own cname
	// defer C.free(unsafe.Pointer(cname))
	return C.GoString(cname)
}

func (dm *dataMemberImpl) IsaPointer() bool {
	o := C.CRoot_DataMember_IsaPointer(dm.c)
	return c2bool(o)
}

func (dm *dataMemberImpl) IsBasic() bool {
	o := C.CRoot_DataMember_IsBasic(dm.c)
	return c2bool(o)
}

func (dm *dataMemberImpl) IsEnum() bool {
	o := C.CRoot_DataMember_IsEnum(dm.c)
	return c2bool(o)
}

func (dm *dataMemberImpl) IsPersistent() bool {
	o := C.CRoot_DataMember_IsPersistent(dm.c)
	return c2bool(o)
}

func (dm *dataMemberImpl) IsSTLContainer() STLType {
	o := C.CRoot_DataMember_IsSTLContainer(dm.c)
	return STLType(o)
}

type dataTypeImpl struct {
	c C.CRoot_DataType
}

func (dt *dataTypeImpl) GetFullTypeName() string {
	cname := C.CRoot_DataType_GetFullTypeName(dt.c)
	// we dont own cname
	// defer C.free(unsafe.Pointer(cname))
	return C.GoString(cname)
}

func (dt *dataTypeImpl) GetType() DataTypeKind {
	kind := C.CRoot_DataType_GetType(dt.c)
	return DataTypeKind(kind)
}

func (dt *dataTypeImpl) GetTypeName() string {
	cname := C.CRoot_DataType_GetTypeName(dt.c)
	// we dont own cname
	// defer C.free(unsafe.Pointer(cname))
	return C.GoString(cname)
}

func (dt *dataTypeImpl) Size() int {
	o := C.CRoot_DataType_Size(dt.c)
	return int(o)
}

func (dt *dataTypeImpl) Property() int64 {
	o := C.CRoot_DataType_Property(dt.c)
	return int64(o)
}
