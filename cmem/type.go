package cmem

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Kind uint

const (
	Void Kind = 0 + iota
	Uint
	Int
	Uint8
	Int8
	Uint16
	Int16
	Uint32
	Int32
	Uint64
	Int64
	Float
	Double
	LongDouble
	Struct
	Ptr
)

const (
	Array Kind = 255 + iota
	Slice
	String
)

func (k Kind) String() string {
	switch k {
	case Void:
		return "Void"
	case Uint:
		return "Uint"
	case Int:
		return "Int"
	case Float:
		return "Float"
	case Double:
		return "Double"
	case LongDouble:
		return "LongDouble"
	case Uint8:
		return "Uint8"
	case Int8:
		return "Int8"
	case Uint16:
		return "Uint16"
	case Int16:
		return "Int16"
	case Uint32:
		return "Uint32"
	case Int32:
		return "Int32"
	case Uint64:
		return "Uint64"
	case Int64:
		return "Int64"
	case Struct:
		return "Struct"
	case Ptr:
		return "Ptr"
	case Array:
		return "Array"
	case Slice:
		return "Slice"
	case String:
		return "String"
	}
	panic("unreachable")
}

// Type is a C-mem type, describing functions' type arguments
type Type interface {
	// Name returns the type's name.
	Name() string

	// Size returns the number of bytes needed to store
	// a value of the given type.
	Size() uintptr

	// String returns a string representation of the type.
	String() string

	// Kind returns the specific kind of this type
	Kind() Kind

	// Align returns the alignment in bytes of a value of this type.
	Align() int

	// Len returns an array type's length
	// It panics if the type's Kind is not Array.
	Len() int

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array or Ptr
	Elem() Type

	// Field returns a struct type's i'th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(i int) StructField

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField() int

	// GoType returns the reflect.Type this ffi.Type is mirroring
	// It returns nil if there is no such equivalent go type.
	GoType() reflect.Type

	// set_gotype sets the reflect.Type associated with this ffi.Type
	set_gotype(t reflect.Type)
}

type cmem_type struct {
	n    string
	kind Kind
	rt   reflect.Type
}

func (t *cmem_type) Name() string {
	return t.n
}

func (t *cmem_type) Size() uintptr {
	return t.rt.Size()
}

func (t *cmem_type) String() string {
	// fixme:
	return t.n
}

func (t *cmem_type) Kind() Kind {
	return t.kind
}

func (t *cmem_type) Align() int {
	return t.rt.Align()
}

func (t *cmem_type) Len() int {
	if t.Kind() != Array {
		panic("cmem: Len of non-array type")
	}
	tt := (*cmem_array_type)(unsafe.Pointer(&t))
	return tt.Len()
}

func (t *cmem_type) Elem() Type {
	switch t.Kind() {
	case Array:
		tt := (*cmem_array_type)(unsafe.Pointer(&t))
		return tt.Elem()
	case Ptr:
		tt := (*cmem_ptr_type)(unsafe.Pointer(&t))
		return tt.Elem()
	case Slice:
		tt := (*cmem_slice_type)(unsafe.Pointer(&t))
		return tt.Elem()
	}
	panic("cmem: Elem of invalid type")
}

func (t *cmem_type) NumField() int {
	if t.Kind() != Struct {
		panic("cmem: NumField of non-struct type")
	}
	tt := (*cmem_struct_type)(unsafe.Pointer(&t))
	return tt.NumField()
}

func (t *cmem_type) Field(i int) StructField {
	if t.Kind() != Struct {
		panic("cmem: Field of non-struct type")
	}
	tt := (*cmem_struct_type)(unsafe.Pointer(&t))
	return tt.Field(i)
}

func (t *cmem_type) GoType() reflect.Type {
	return t.rt
}

func (t *cmem_type) set_gotype(rt reflect.Type) {
	t.rt = rt
}

var (
	C_void       Type = &cmem_type{"void", Void, nil}
	C_uchar           = &cmem_type{"unsigned char", Uint8, reflect.TypeOf(uint8(0))}
	C_char            = &cmem_type{"char", Int8, reflect.TypeOf(int8(0))}
	C_ushort          = &cmem_type{"unsigned short", Uint16, reflect.TypeOf(uint16(0))}
	C_short           = &cmem_type{"short", Int16, reflect.TypeOf(int16(0))}
	C_uint            = &cmem_type{"unsigned int", Uint, reflect.TypeOf(uint(0))}
	C_int             = &cmem_type{"int", Int, reflect.TypeOf(int(0))}
	C_ulong           = &cmem_type{"unsigned long", Uint16, reflect.TypeOf(uint64(0))}
	C_long            = &cmem_type{"long", Int16, reflect.TypeOf(int64(0))}
	C_uint8           = &cmem_type{"uint8", Uint8, reflect.TypeOf(uint8(0))}
	C_int8            = &cmem_type{"int8", Int8, reflect.TypeOf(int8(0))}
	C_uint16          = &cmem_type{"uint16", Uint16, reflect.TypeOf(uint16(0))}
	C_int16           = &cmem_type{"int16", Int16, reflect.TypeOf(int16(0))}
	C_uint32          = &cmem_type{"uint32", Uint32, reflect.TypeOf(uint32(0))}
	C_int32           = &cmem_type{"int32", Int32, reflect.TypeOf(int32(0))}
	C_uint64          = &cmem_type{"uint64", Uint64, reflect.TypeOf(uint64(0))}
	C_int64           = &cmem_type{"int64", Int64, reflect.TypeOf(int64(0))}
	C_float           = &cmem_type{"float", Float, reflect.TypeOf(float32(0.))}
	C_double          = &cmem_type{"double", Double, reflect.TypeOf(float64(0.))}
	C_longdouble      = &cmem_type{"long double", LongDouble, nil}
	C_pointer         = &cmem_ptr_type{
		cmem_type: cmem_type{"*", Ptr, reflect.TypeOf(nil)},
		elem:      nil,
	}

	C_string = &cmem_string_type{
		cmem_type: cmem_type{"char*", Ptr, reflect.TypeOf("")},
		elem:      C_char,
	}
)

type cmem_string_type struct {
	cmem_type
	elem Type
}

func (t *cmem_string_type) Elem() Type {
	return t.elem
}

type StructField struct {
	Name   string  // Name is the field name
	Type   Type    // field type
	Offset uintptr // offset within struct, in bytes
}

type cmem_struct_type struct {
	cmem_type
	fields []StructField
}

func (t *cmem_struct_type) NumField() int {
	return len(t.fields)
}

func (t *cmem_struct_type) Field(i int) StructField {
	if i < 0 || i >= len(t.fields) {
		panic("cmem: field index out of range")
	}
	return t.fields[i]
}

func (t *cmem_struct_type) set_gotype(rt reflect.Type) {
	t.cmem_type.rt = rt
}

var g_id_ch chan int

// NewStructType creates a new cmem_type describing a C-struct
func NewStructType(typ reflect.Type) (Type, error) {
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("cmem.NewStructType: type isn't of Kind reflect.Struct")
	}

	name := typ.Name()
	if name == "" {
		// anonymous type...
		// generate some id.
		name = fmt.Sprintf("_cmem_anon_type_%d", <-g_id_ch)
	}
	nfields := typ.NumField()
	if t := TypeByName(name); t != nil {
		// check the definitions are the same
		if t.GoType() != typ {
			return nil, fmt.Errorf("cmem.NewStructType: inconsistent re-declaration of [%s]", name)
		}
		return t, nil
	}
	t := &cmem_struct_type{
		cmem_type: cmem_type{n: name, kind: Struct, rt: typ},
		fields:    make([]StructField, typ.NumField()),
	}

	for i := 0; i < nfields; i++ {
		rf := typ.Field(i)
		t.fields[i] = StructField{
			Name:   rf.Name,
			Type:   ctype_from_gotype(rf.Type),
			Offset: rf.Offset,
		}
		if t.fields[i].Type == nil {
			return nil, fmt.Errorf("cmem.NewStructType: no cmem.Type for reflect.Type %q", rf.Type.Name())
		}
	}
	register_type(t)
	return t, nil
}

type cmem_array_type struct {
	cmem_type
	len  int
	elem Type
}

func (t *cmem_array_type) Len() int {
	return t.len
}

func (t *cmem_array_type) Elem() Type {
	return t.elem
}

// NewArrayType creates a new cmem_type with the given size and element type.
func NewArrayType(typ reflect.Type) (Type, error) {
	if typ.Kind() != reflect.Array {
		return nil, fmt.Errorf("cmem: expected a reflect.Array kind")
	}

	elmt := ctype_from_gotype(typ.Elem())
	n := fmt.Sprintf("%s[%d]", elmt.Name(), typ.Len())
	if t := TypeByName(n); t != nil {
		return t, nil
	}
	t := &cmem_array_type{
		cmem_type: cmem_type{n: n, kind: Array, rt: typ},
		len:       typ.Len(),
		elem:      elmt,
	}

	register_type(t)
	return t, nil
}

type cmem_ptr_type struct {
	cmem_type
	elem Type
}

func (t *cmem_ptr_type) Size() uintptr {
	return ptrSize
}

func (t *cmem_ptr_type) Elem() Type {
	return t.elem
}

// NewPointerType creates a new cmem_type from the corresponding reflect.Type
func NewPointerType(typ reflect.Type) (Type, error) {
	if typ.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("cmem: expected a reflect.Ptr kind")
	}

	elmt := ctype_from_gotype(typ.Elem())
	n := elmt.Name() + "*"
	if t := TypeByName(n); t != nil {
		return t, nil
	}
	t := &cmem_ptr_type{
		cmem_type: cmem_type{n: n, kind: Ptr, rt: typ},
		elem:      elmt,
	}

	register_type(t)
	return t, nil
}

type cmem_slice_header struct {
	Len  croot_int
	Cap  croot_int
	Data unsafe.Pointer
}

type cmem_slice_type struct {
	cmem_type
	elem Type
}

func (t *cmem_slice_type) Elem() Type {
	return t.elem
}

// NewSliceType creates a new cmem_type slice from the corresponding reflect.Type
func NewSliceType(typ reflect.Type) (Type, error) {
	if typ.Kind() != reflect.Slice {
		return nil, fmt.Errorf("cmem: expected a reflect.Slice kind")
	}

	elmt := ctype_from_gotype(typ.Elem())
	n := elmt.Name() + "[]"
	if t := TypeByName(n); t != nil {
		return t, nil
	}
	t := &cmem_slice_type{
		cmem_type: cmem_type{n: n, kind: Slice, rt: typ},
		elem:      elmt,
	}

	register_type(t)
	return t, nil
}

// PtrTo returns the pointer type with element t.
// For example, if t represents type Foo, PtrTo(t) represents *Foo.
func PtrTo(t Type) Type {
	rt := t.GoType()
	typ, err := NewPointerType(reflect.PtrTo(rt))
	if err != nil {
		return nil
	}
	return typ
}

// TypeOf returns the cmem Type of the value in the interface{}.
// TypeOf(nil) returns nil
// TypeOf(reflect.Type) returns the cmem Type corresponding to the reflected value
func TypeOf(i interface{}) Type {
	switch typ := i.(type) {
	case reflect.Type:
		return ctype_from_gotype(typ)
	case reflect.Value:
		return ctype_from_gotype(typ.Type())
	default:
		rt := reflect.TypeOf(i)
		return ctype_from_gotype(rt)
	}
	panic("unreachable")
}

// the global map of types
var g_types map[string]Type

// TypeByName returns a ffi.Type by name.
// Returns nil if no such type exists
func TypeByName(n string) Type {
	t, ok := g_types[n]
	if ok {
		return t
	}
	return nil
}

func register_type(t Type) {
	g_types[t.Name()] = t
}

func ctype_from_gotype(rt reflect.Type) Type {
	var t Type

	switch rt.Kind() {
	case reflect.Int:
		t = C_int

	case reflect.Int8:
		t = C_int8

	case reflect.Int16:
		t = C_int16

	case reflect.Int32:
		t = C_int32

	case reflect.Int64:
		t = C_int64

	case reflect.Uint:
		t = C_uint

	case reflect.Uint8:
		t = C_uint8

	case reflect.Uint16:
		t = C_uint16

	case reflect.Uint32:
		t = C_uint32

	case reflect.Uint64:
		t = C_uint64

	case reflect.Float32:
		t = C_float

	case reflect.Float64:
		t = C_double

	case reflect.Array:
		ct, err := NewArrayType(rt)
		if err != nil {
			panic("cmem: " + err.Error())
		}
		t = ct

	case reflect.Ptr:
		ct, err := NewPointerType(rt)
		if err != nil {
			panic("cmem: " + err.Error())
		}
		t = ct

	case reflect.Slice:
		ct, err := NewSliceType(rt)
		if err != nil {
			panic("cmem: " + err.Error())
		}
		t = ct

	case reflect.Struct:
		ct, err := NewStructType(rt)
		if err != nil {
			panic("cmem: " + err.Error())
		}
		t = ct

	case reflect.String:
		t = C_string

	default:
		panic("unhandled kind [" + rt.Kind().String() + "]")
	}

	return t
}

func init() {
	// init out id counter channel
	g_id_ch = make(chan int, 1)
	go func() {
		i := 0
		for {
			g_id_ch <- i
			i++
		}
	}()

	g_types = make(map[string]Type)

	// initialize all builtin types
	init_type := func(t Type) {
		n := t.Name()
		//fmt.Printf("ctype [%s] - size: %v...\n", n, t.Size())
		if _, ok := g_types[n]; ok {
			//fmt.Printf("cmem type [%s] already registered\n", n)
			return
		}
		//fmt.Printf("ctype [%s] - size: %v\n", n, t.Size())
		g_types[n] = t
	}

	init_type(C_void)
	init_type(C_uchar)
	init_type(C_char)
	init_type(C_ushort)
	init_type(C_short)
	init_type(C_uint)
	init_type(C_int)
	init_type(C_ulong)
	init_type(C_long)
	init_type(C_uint8)
	init_type(C_int8)
	init_type(C_uint16)
	init_type(C_int16)
	init_type(C_uint32)
	init_type(C_int32)
	init_type(C_uint64)
	init_type(C_int64)
	init_type(C_float)
	init_type(C_double)
	init_type(C_longdouble)
	init_type(C_pointer)

}

// make sure cmem_types satisfy cmem.Type interface
var _ Type = (*cmem_type)(nil)
var _ Type = (*cmem_array_type)(nil)
var _ Type = (*cmem_ptr_type)(nil)
var _ Type = (*cmem_slice_type)(nil)
var _ Type = (*cmem_string_type)(nil)
var _ Type = (*cmem_struct_type)(nil)

// EOF
