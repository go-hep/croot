package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type ReflexType struct {
	t C.CRoot_Reflex_Type
}

type ReflexMember struct {
	m C.CRoot_Reflex_Member
}

type ReflexScope struct {
	s C.CRoot_Reflex_Scope
}

type ReflexClassBuilder struct {
	c C.CRoot_Reflex_ClassBuilder
}

type ReflexFunctionBuilder struct {
	f C.CRoot_Reflex_FunctionBuilder
}

type ReflexPropertyList struct {
	c C.CRoot_Reflex_PropertyList
}

type ReflexStubFunction C.CRoot_Reflex_StubFunction
type ReflexOffsetFunction C.CRoot_Reflex_OffsetFunction

type Reflex_ENTITY_DESCRIPTION C.CRoot_Reflex_ENTITY_DESCRIPTION

const (
	Reflex_PUBLIC          Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_PUBLIC
	Reflex_PROTECTED       Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_PROTECTED
	Reflex_PRIVATE         Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_PRIVATE
	Reflex_REGISTER        Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_REGISTER
	Reflex_STATIC          Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_STATIC
	Reflex_CONSTRUCTOR     Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_CONSTRUCTOR
	Reflex_DESTRUCTOR      Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_DESTRUCTOR
	Reflex_EXPLICIT        Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_EXPLICIT
	Reflex_EXTERN          Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_EXTERN
	Reflex_COPYCONSTRUCTOR Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_COPYCONSTRUCTOR
	Reflex_OPERATOR        Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_OPERATOR
	Reflex_INLINE          Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_INLINE
	Reflex_CONVERTER       Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_CONVERTER
	Reflex_AUTO            Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_AUTO
	Reflex_MUTABLE         Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_MUTABLE
	Reflex_CONST           Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_CONST
	Reflex_VOLATILE        Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_VOLATILE
	Reflex_REFERENCE       Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_REFERENCE
	Reflex_ABSTRACT        Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_ABSTRACT
	Reflex_VIRTUAL         Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_VIRTUAL
	Reflex_TRANSIENT       Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TRANSIENT
	Reflex_ARTIFICIAL      Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_ARTIFICIAL

	// the bits 31 - 28 are reserved for template default arguments

	Reflex_TEMPLATEDEFAULTS1  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS1
	Reflex_TEMPLATEDEFAULTS2  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS2
	Reflex_TEMPLATEDEFAULTS3  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS3
	Reflex_TEMPLATEDEFAULTS4  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS4
	Reflex_TEMPLATEDEFAULTS5  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS5
	Reflex_TEMPLATEDEFAULTS6  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS6
	Reflex_TEMPLATEDEFAULTS7  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS7
	Reflex_TEMPLATEDEFAULTS8  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS8
	Reflex_TEMPLATEDEFAULTS9  Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS9
	Reflex_TEMPLATEDEFAULTS10 Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS10
	Reflex_TEMPLATEDEFAULTS11 Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS11
	Reflex_TEMPLATEDEFAULTS12 Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS12
	Reflex_TEMPLATEDEFAULTS13 Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS13
	Reflex_TEMPLATEDEFAULTS14 Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS14
	Reflex_TEMPLATEDEFAULTS15 Reflex_ENTITY_DESCRIPTION = C.CRoot_Reflex_TEMPLATEDEFAULTS15
)

type Reflex_ENTITY_HANDLING C.CRoot_Reflex_ENTITY_HANDLING

const (
	Reflex_FINAL     Reflex_ENTITY_HANDLING = C.CRoot_Reflex_FINAL
	Reflex_QUALIFIED Reflex_ENTITY_HANDLING = C.CRoot_Reflex_QUALIFIED
	Reflex_SCOPED    Reflex_ENTITY_HANDLING = C.CRoot_Reflex_SCOPED
	Reflex_F         Reflex_ENTITY_HANDLING = C.CRoot_Reflex_F
	Reflex_Q         Reflex_ENTITY_HANDLING = C.CRoot_Reflex_Q
	Reflex_S         Reflex_ENTITY_HANDLING = C.CRoot_Reflex_S
)

type Reflex_TYPE C.CRoot_Reflex_TYPE

const (
	Reflex_CLASS                  Reflex_TYPE = C.CRoot_Reflex_CLASS
	Reflex_STRUCT                 Reflex_TYPE = C.CRoot_Reflex_STRUCT
	Reflex_ENUM                   Reflex_TYPE = C.CRoot_Reflex_ENUM
	Reflex_FUNCTION               Reflex_TYPE = C.CRoot_Reflex_FUNCTION
	Reflex_ARRAY                  Reflex_TYPE = C.CRoot_Reflex_ARRAY
	Reflex_FUNDAMENTAL            Reflex_TYPE = C.CRoot_Reflex_FUNDAMENTAL
	Reflex_POINTER                Reflex_TYPE = C.CRoot_Reflex_POINTER
	Reflex_POINTERTOMEMBER        Reflex_TYPE = C.CRoot_Reflex_POINTERTOMEMBER
	Reflex_TYPEDEF                Reflex_TYPE = C.CRoot_Reflex_TYPEDEF
	Reflex_UNION                  Reflex_TYPE = C.CRoot_Reflex_UNION
	Reflex_TYPETEMPLATEINSTANCE   Reflex_TYPE = C.CRoot_Reflex_TYPETEMPLATEINSTANCE
	Reflex_MEMBERTEMPLATEINSTANCE Reflex_TYPE = C.CRoot_Reflex_MEMBERTEMPLATEINSTANCE
	Reflex_NAMESPACE              Reflex_TYPE = C.CRoot_Reflex_NAMESPACE
	Reflex_DATAMEMBER             Reflex_TYPE = C.CRoot_Reflex_DATAMEMBER
	Reflex_FUNCTIONMEMBER         Reflex_TYPE = C.CRoot_Reflex_FUNCTIONMEMBER
	Reflex_UNRESOLVED             Reflex_TYPE = C.CRoot_Reflex_UNRESOLVED
)

type Reflex_REPRESTYPE C.CRoot_Reflex_REPRESTYPE

const (
	Reflex_REPRES_POINTER            Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_POINTER
	Reflex_REPRES_CHAR               Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_CHAR
	Reflex_REPRES_SIGNED_CHAR        Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_SIGNED_CHAR
	Reflex_REPRES_SHORT_INT          Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_SHORT_INT
	Reflex_REPRES_INT                Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_INT
	Reflex_REPRES_LONG_INT           Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_LONG_INT
	Reflex_REPRES_UNSIGNED_CHAR      Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_UNSIGNED_CHAR
	Reflex_REPRES_UNSIGNED_SHORT_INT Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_UNSIGNED_SHORT_INT
	Reflex_REPRES_UNSIGNED_INT       Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_UNSIGNED_INT
	Reflex_REPRES_UNSIGNED_LONG_INT  Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_UNSIGNED_LONG_INT
	Reflex_REPRES_BOOL               Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_BOOL
	Reflex_REPRES_FLOAT              Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_FLOAT
	Reflex_REPRES_DOUBLE             Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_DOUBLE
	Reflex_REPRES_LONG_DOUBLE        Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_LONG_DOUBLE
	Reflex_REPRES_VOID               Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_VOID
	Reflex_REPRES_LONGLONG           Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_LONGLONG
	Reflex_REPRES_ULONGLONG          Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_ULONGLONG
	Reflex_REPRES_STRUCT             Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_STRUCT
	Reflex_REPRES_CLASS              Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_CLASS
	Reflex_REPRES_ENUM               Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_ENUM
	Reflex_REPRES_NOTYPE             Reflex_REPRESTYPE = C.CRoot_Reflex_REPRES_NOTYPE
)

type Reflex_EMEMBERQUERY C.CRoot_Reflex_EMEMBERQUERY

const (
	Reflex_INHERITEDMEMBERS_DEFAULT Reflex_EMEMBERQUERY = C.CRoot_Reflex_INHERITEDMEMBERS_DEFAULT
	Reflex_INHERITEDMEMBERS_NO      Reflex_EMEMBERQUERY = C.CRoot_Reflex_INHERITEDMEMBERS_NO
	Reflex_INHERITEDMEMBERS_ALSO    Reflex_EMEMBERQUERY = C.CRoot_Reflex_INHERITEDMEMBERS_ALSO
)

type Reflex_EDELAYEDLOADSETTING C.CRoot_Reflex_EDELAYEDLOADSETTING

const (
	Reflex_DELAYEDLOAD_OFF Reflex_EDELAYEDLOADSETTING = C.CRoot_Reflex_DELAYEDLOAD_OFF
	Reflex_DELAYEDLOAD_ON  Reflex_EDELAYEDLOADSETTING = C.CRoot_Reflex_DELAYEDLOAD_ON
)

func Reflex_FireClassCallback(t ReflexType) {
	C.CRoot_Reflex_FireClassCallback(t.t)
}

func Reflex_FireFunctionCallback(m ReflexMember) {
	C.CRoot_Reflex_FireFunctionCallback(m.m)
}

func new_reflex_type(t C.CRoot_Reflex_Type) *ReflexType {
	o := &ReflexType{t: t}
	runtime.SetFinalizer(o, (*ReflexType).Delete)
	return o
}

func NewReflexType(name string, modifiers uint32) *ReflexType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	t := C.CRoot_Reflex_Type_new(c_name, C.uint(modifiers))
	return new_reflex_type(t)
}

func (t *ReflexType) Delete() {
	C.CRoot_Reflex_Type_delete(t.t)
}

func (t *ReflexType) Id() uintptr {
	c_id := C.CRoot_Reflex_Type_Id(t.t)
	return uintptr(c_id)
}

func (t *ReflexType) ArrayLength() int {
	return int(C.CRoot_Reflex_Type_ArrayLength(t.t))
}

func ReflexType_ByName(name string) *ReflexType {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_t := C.CRoot_Reflex_Type_ByName(c_name)
	t := &ReflexType{t: c_t}
	runtime.SetFinalizer(t, (*ReflexType).Delete)
	return t
}

func (t *ReflexType) DataMemberAt(nth int, query Reflex_EMEMBERQUERY) *ReflexMember {
	c_query := C.CRoot_Reflex_EMEMBERQUERY(query)
	c_mbr := C.CRoot_Reflex_Type_DataMemberAt(t.t, C.size_t(nth), c_query)
	mbr := new_reflex_member(c_mbr)
	return mbr
}

func (t *ReflexType) DataMemberSize(query Reflex_EMEMBERQUERY) int {
	c_query := C.CRoot_Reflex_EMEMBERQUERY(query)
	sz := C.CRoot_Reflex_Type_DataMemberSize(t.t, c_query)
	return int(sz)
}

func (t *ReflexType) IsAbstract() bool {
	return bool(C.CRoot_Reflex_Type_IsAbstract(t.t))
}

func (t *ReflexType) IsArray() bool {
	return bool(C.CRoot_Reflex_Type_IsArray(t.t))
}

func (t *ReflexType) IsClass() bool {
	return bool(C.CRoot_Reflex_Type_IsClass(t.t))
}

func (t *ReflexType) IsComplete() bool {
	return bool(C.CRoot_Reflex_Type_IsComplete(t.t))
}

func (t *ReflexType) IsEquivalentTo(other *ReflexType, modifiers_mask uint) bool {
	return bool(C.CRoot_Reflex_Type_IsEquivalentTo(t.t, other.t, C.uint(modifiers_mask)))
}

func (t *ReflexType) IsFunction() bool {
	return bool(C.CRoot_Reflex_Type_IsFunction(t.t))
}

func (t *ReflexType) IsFundamental() bool {
	return bool(C.CRoot_Reflex_Type_IsFundamental(t.t))
}

func (t *ReflexType) IsPrivate() bool {
	return bool(C.CRoot_Reflex_Type_IsPrivate(t.t))
}

func (t *ReflexType) IsPublic() bool {
	return bool(C.CRoot_Reflex_Type_IsPublic(t.t))
}

func (t *ReflexType) IsProtected() bool {
	return bool(C.CRoot_Reflex_Type_IsProtected(t.t))
}

func (t *ReflexType) IsPointer() bool {
	return bool(C.CRoot_Reflex_Type_IsPointer(t.t))
}

func (t *ReflexType) IsPointerToMember() bool {
	return bool(C.CRoot_Reflex_Type_IsPointerToMember(t.t))
}

func (t *ReflexType) IsReference() bool {
	return bool(C.CRoot_Reflex_Type_IsReference(t.t))
}

func (t *ReflexType) IsStruct() bool {
	return bool(C.CRoot_Reflex_Type_IsStruct(t.t))
}

func (t *ReflexType) IsVirtual() bool {
	return bool(C.CRoot_Reflex_Type_IsVirtual(t.t))
}

func (t *ReflexType) MemberAt(nth int, query Reflex_EMEMBERQUERY) *ReflexMember {
	c_query := C.CRoot_Reflex_EMEMBERQUERY(query)
	c_mbr := C.CRoot_Reflex_Type_MemberAt(t.t, C.size_t(nth), c_query)
	mbr := new_reflex_member(c_mbr)
	return mbr
}

func (t *ReflexType) MemberSize(query Reflex_EMEMBERQUERY) int {
	c_query := C.CRoot_Reflex_EMEMBERQUERY(query)
	sz := C.CRoot_Reflex_Type_MemberSize(t.t, c_query)
	return int(sz)
}

func (t *ReflexType) Name() string {
	c_name := C.CRoot_Reflex_Type_Name(t.t)
	name := C.GoString(c_name)
	return name
}

func (t *ReflexType) Properties() ReflexPropertyList {
	plist := C.CRoot_Reflex_Type_Properties(t.t)
	return ReflexPropertyList{c: plist}
}

func (t *ReflexType) RawType() *ReflexType {
	r := C.CRoot_Reflex_Type_RawType(t.t)
	return new_reflex_type(r)
}

func (t *ReflexType) SizeOf() uintptr {
	sz := C.CRoot_Reflex_Type_SizeOf(t.t)
	return uintptr(sz)
}

func (t *ReflexType) ToType() *ReflexType {
	r := C.CRoot_Reflex_Type_ToType(t.t)
	return new_reflex_type(r)
}

func ReflexType_TypeAt(nth int) *ReflexType {
	t := C.CRoot_Reflex_Type_TypeAt(C.size_t(nth))
	return new_reflex_type(t)
}

func ReflexType_TypeSize() int {
	sz := C.CRoot_Reflex_Type_TypeSize()
	return int(sz)
}

func (t *ReflexType) TypeType() Reflex_TYPE {
	tt := C.CRoot_Reflex_Type_TypeType(t.t)
	return Reflex_TYPE(tt)
}

func (t *ReflexType) Unload() {
	C.CRoot_Reflex_Type_Unload(t.t)
}

func (t *ReflexType) UpdateMembers() {
	C.CRoot_Reflex_Type_UpdateMembers(t.t)
}

func (t *ReflexType) AddDataMember(dm *ReflexMember) {
	C.CRoot_Reflex_Type_AddDataMember(t.t, dm.m)
}

//FIXME
//func (t *ReflexType) AddDataMember2() {}

func (t *ReflexType) RemoveDataMember(dm *ReflexMember) {
	C.CRoot_Reflex_Type_RemoveDataMember(t.t, dm.m)
}

func (t *ReflexType) SetSize(sz uintptr) {
	C.CRoot_Reflex_Type_SetSize(t.t, C.size_t(sz))
}

func (t *ReflexType) RepresType() Reflex_REPRESTYPE {
	r := C.CRoot_Reflex_Type_RepresType(t.t)
	return Reflex_REPRESTYPE(r)
}

func new_reflex_member(m C.CRoot_Reflex_Member) *ReflexMember {
	o := &ReflexMember{m: m}
	runtime.SetFinalizer(o, (*ReflexMember).Delete)
	return o
}

func NewReflexMember() *ReflexMember {
	m := C.CRoot_Reflex_Member_new()
	return new_reflex_member(m)
}

func (m *ReflexMember) Delete() {
	C.CRoot_Reflex_Member_delete(m.m)
}

func (m *ReflexMember) IsDataMember() bool {
	return bool(C.CRoot_Reflex_Member_IsDataMember(m.m))
}

func (m *ReflexMember) IsPrivate() bool {
	return bool(C.CRoot_Reflex_Member_IsPrivate(m.m))
}

func (m *ReflexMember) IsProtected() bool {
	return bool(C.CRoot_Reflex_Member_IsProtected(m.m))
}

func (m *ReflexMember) IsPublic() bool {
	return bool(C.CRoot_Reflex_Member_IsPublic(m.m))
}

func (m *ReflexMember) IsTransient() bool {
	return bool(C.CRoot_Reflex_Member_IsTransient(m.m))
}

func (m *ReflexMember) IsVirtual() bool {
	return bool(C.CRoot_Reflex_Member_IsVirtual(m.m))
}

func (m *ReflexMember) MemberType() Reflex_TYPE {
	r := C.CRoot_Reflex_Member_MemberType(m.m)
	return Reflex_TYPE(r)
}

func (m *ReflexMember) Name() string {
	n := C.CRoot_Reflex_Member_Name(m.m)
	return C.GoString(n)
}

func (m *ReflexMember) Offset() uintptr {
	o := C.CRoot_Reflex_Member_Offset(m.m)
	return uintptr(o)
}

//FIXME
//func (m *ReflexMember) InterpreterOffset() {}

func (m *ReflexMember) TypeOf() *ReflexType {
	t := C.CRoot_Reflex_Member_TypeOf(m.m)
	return new_reflex_type(t)
}

func (m *ReflexMember) Stubcontext() uintptr {
	ctx := C.CRoot_Reflex_Member_Stubcontext(m.m)
	return uintptr(ctx)
}

func (m *ReflexMember) Stubfunction() ReflexStubFunction {
	fct := C.CRoot_Reflex_Member_Stubfunction(m.m)
	return (ReflexStubFunction)(fct)
}

func (m *ReflexMember) Properties() ReflexPropertyList {
	plist := C.CRoot_Reflex_Member_Properties(m.m)
	return ReflexPropertyList{c: plist}
}

// propertylist api

func (list *ReflexPropertyList) Count() int {
	sz := C.CRoot_Reflex_PropertyList_PropertyCount(list.c)
	return int(sz)
}

func (list *ReflexPropertyList) AsString(idx int) string {
	c_str := C.CRoot_Reflex_PropertyList_PropertyAsString(list.c, C.size_t(idx))
	defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

func (list *ReflexPropertyList) Keys() string {
	c_str := C.CRoot_Reflex_PropertyList_PropertyKeys(list.c)
	defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

func NewReflexPointerBuilder(r *ReflexType) *ReflexType {
	ptr := C.CRoot_Reflex_PointerBuilder_new(r.t)
	return new_reflex_type(ptr)
}

func NewReflexArrayBuilder(r *ReflexType, n int) *ReflexType {
	arr := C.CRoot_Reflex_ArrayBuilder_new(r.t, C.size_t(n))
	return new_reflex_type(arr)
}

func NewReflexFunctionTypeBuilder(r *ReflexType) *ReflexType {
	f := C.CRoot_Reflex_FunctionTypeBuilder_new(r.t)
	return new_reflex_type(f)
}

func NewReflexFunctionTypeBuilder1(r, t0 *ReflexType) *ReflexType {
	f := C.CRoot_Reflex_FunctionTypeBuilder_new1(r.t, t0.t)
	return new_reflex_type(f)
}

func NewReflexFunctionTypeBuilder2(r, t0, t1 *ReflexType) *ReflexType {
	f := C.CRoot_Reflex_FunctionTypeBuilder_new2(r.t, t0.t, t1.t)
	return new_reflex_type(f)
}

func NewReflexFunctionTypeBuilder3(r, t0, t1, t2 *ReflexType) *ReflexType {
	f := C.CRoot_Reflex_FunctionTypeBuilder_new3(r.t, t0.t, t1.t, t2.t)
	return new_reflex_type(f)
}

func NewReflexClassBuilder(name string, size uintptr, modifiers uint32, typ Reflex_TYPE) *ReflexClassBuilder {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_null := unsafe.Pointer(nil) // no typeinfo available from Go...

	c_cb := C.CRoot_Reflex_ClassBuilder_new(c_name, c_null, C.size_t(size), C.uint(modifiers), C.CRoot_Reflex_TYPE(typ))
	cb := &ReflexClassBuilder{c: c_cb}
	runtime.SetFinalizer(cb, (*ReflexClassBuilder).Delete)
	return cb
}

func (c *ReflexClassBuilder) Delete() {
	if c.c != nil {
		C.CRoot_Reflex_ClassBuilder_delete(c.c)
		c.c = nil
	}
}

func (c *ReflexClassBuilder) AddDataMember(t *ReflexType, name string, offset uintptr, modifiers uint32) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_offset := C.size_t(offset)
	c_modifiers := C.uint(modifiers)
	C.CRoot_Reflex_ClassBuilder_AddDataMember(c.c, t.t, c_name, c_offset, c_modifiers)
}

func (c *ReflexClassBuilder) AddFunctionMember(t *ReflexType, name string, stubFct ReflexStubFunction, ctx unsafe.Pointer, modifiers uint32) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_null := unsafe.Pointer(nil)
	c_ctx := ctx
	c_params := (*C.char)(c_null)
	c_stubFct := C.CRoot_Reflex_StubFunction(stubFct)
	C.CRoot_Reflex_ClassBuilder_AddFunctionMember(c.c, t.t, c_name, c_stubFct, c_ctx, c_params, C.uint(modifiers))
}

func (c *ReflexClassBuilder) AddProperty(key, value string) {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))
	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))
	C.CRoot_Reflex_ClassBuilder_AddProperty(c.c, c_key, c_value)
}

func (c *ReflexClassBuilder) ToType() *ReflexType {
	t := C.CRoot_Reflex_ClassBuilder_ToType(c.c)
	return new_reflex_type(t)
}

func NewReflexFunctionBuilder(t *ReflexType, name string, stubFct ReflexStubFunction, modifiers uint32) *ReflexFunctionBuilder {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_null := unsafe.Pointer(nil)
	c_params := (*C.char)(c_null)
	c_stubFct := C.CRoot_Reflex_StubFunction(stubFct)
	c_f := C.CRoot_Reflex_FunctionBuilder_new(t.t, c_name, c_stubFct, c_null, c_params, C.uchar(modifiers))
	f := &ReflexFunctionBuilder{f: c_f}
	runtime.SetFinalizer(f, (*ReflexFunctionBuilder).Delete)
	return f
}

func (f *ReflexFunctionBuilder) Delete() {
	C.CRoot_Reflex_FunctionBuilder_delete(f.f)
}

func (f *ReflexFunctionBuilder) ToMember() *ReflexMember {
	m := C.CRoot_Reflex_FunctionBuilder_ToMember(f.f)
	return new_reflex_member(m)
}

//eof
