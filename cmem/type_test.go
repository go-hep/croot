package cmem_test

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
	"unsafe"

	"github.com/go-hep/croot/cmem"
)

func eq(t *testing.T, ref, chk interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = path.Base(file)
	if !reflect.DeepEqual(ref, chk) {
		err := fmt.Errorf("%s:%d: expected [%v], got [%v]", file, line, ref, chk)
		t.Errorf(err.Error())
	}
}

func TestBuiltinTypes(t *testing.T) {
	for _, table := range []struct {
		n  string
		t  cmem.Type
		rt reflect.Type
	}{
		{"unsigned char", cmem.C_uchar, reflect.TypeOf(byte(0))},
		{"char", cmem.C_char, reflect.TypeOf(byte(0))},

		{"int8", cmem.C_int8, reflect.TypeOf(int8(0))},
		{"uint8", cmem.C_uint8, reflect.TypeOf(uint8(0))},
		{"int16", cmem.C_int16, reflect.TypeOf(int16(0))},
		{"uint16", cmem.C_uint16, reflect.TypeOf(uint16(0))},
		{"int32", cmem.C_int32, reflect.TypeOf(int32(0))},
		{"uint32", cmem.C_uint32, reflect.TypeOf(uint32(0))},
		{"int64", cmem.C_int64, reflect.TypeOf(int64(0))},
		{"uint64", cmem.C_uint64, reflect.TypeOf(uint64(0))},

		{"float", cmem.C_float, reflect.TypeOf(float32(0))},
		{"double", cmem.C_double, reflect.TypeOf(float64(0))},
		//FIXME: use float128 when/if available
		//{"long double", cmem.C_longdouble, reflect.TypeOf(complex128(0))},

		{"*", cmem.C_pointer, reflect.TypeOf((*int)(nil))},
	} {
		if table.n != table.t.Name() {
			t.Errorf("expected [%s], got [%s]", table.n, table.t.Name())
		}
		if table.t.Size() != table.rt.Size() {
			t.Errorf("expected [%d], got [%d] (type=%q)", table.t.Size(), table.rt.Size(), table.n)
		}
	}
}

type struct_0 struct {
	A int
}

type struct_1 struct {
	A int
	B int
}

type struct_2 struct {
	F1 uint8
	F2 int16
	F3 int32
	F4 uint8
}

func TestNewStructType(t *testing.T) {

	arr10, err := cmem.NewArrayType(reflect.TypeOf([10]int32{}))
	if err != nil {
		t.Errorf(err.Error())
	}
	eq(t, int(10), arr10.Len())

	for _, rt := range []reflect.Type{
		reflect.TypeOf(struct_0{}),
		reflect.TypeOf(struct_1{}),
		reflect.TypeOf(struct_2{}),
		//FIXME: 32b/64b alignement differ!!
		// make 2 tests!
		// {"struct_3",
		// 	[]cmem.Field{
		// 		{"F1", cmem.C_uint8},
		// 		{"F2", arr10},
		// 		{"F3", cmem.C_int32},
		// 		{"F4", cmem.C_uint8},
		// 	},
		// 	56,
		// 	[]uintptr{0, 8, 48, 52},
		// },
	} {
		name := rt.Name()
		typ, err := cmem.NewStructType(rt)
		if err != nil {
			t.Errorf(err.Error())
		}
		eq(t, name, typ.Name())
		//eq(t, table.size, typ.Size())
		if rt.Size() != typ.Size() {
			t.Errorf("expected size [%d] got [%d] (type=%q)", rt.Size(), typ.Size(), name)
		}
		eq(t, rt.NumField(), typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			if rt.Field(i).Offset != typ.Field(i).Offset {
				t.Errorf("type=%q field=%d: expected offset [%d]. got [%d]", name, i, rt.Field(i).Offset, typ.Field(i).Offset)
			}
			//eq(t, table.offsets[i], typ.Field(i).Offset)
		}
		eq(t, cmem.Struct, typ.Kind())
	}
}

type s_0 struct {
	A int32
}

func TestNewArrayType(t *testing.T) {

	s_t, err := cmem.NewStructType(reflect.TypeOf(s_0{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	p_s_t, err := cmem.NewPointerType(reflect.PtrTo(reflect.TypeOf(s_0{})))
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, table := range []struct {
		name  string
		n     int
		elem  cmem.Type
		rtype reflect.Type
	}{
		{"uint8[10]", 10, cmem.C_uint8, reflect.TypeOf([10]uint8{})},
		{"uint16[10]", 10, cmem.C_uint16, reflect.TypeOf([10]uint16{})},
		{"uint32[10]", 10, cmem.C_uint32, reflect.TypeOf([10]uint32{})},
		{"uint64[10]", 10, cmem.C_uint64, reflect.TypeOf([10]uint64{})},
		{"int8[10]", 10, cmem.C_int8, reflect.TypeOf([10]int8{})},
		{"int16[10]", 10, cmem.C_int16, reflect.TypeOf([10]int16{})},
		{"int32[10]", 10, cmem.C_int32, reflect.TypeOf([10]int32{})},
		{"int64[10]", 10, cmem.C_int64, reflect.TypeOf([10]int64{})},

		{"float[10]", 10, cmem.C_float, reflect.TypeOf([10]float32{})},
		{"double[10]", 10, cmem.C_double, reflect.TypeOf([10]float64{})},

		{"s_0[10]", 10, s_t, reflect.TypeOf([10]s_0{})},
		{"s_0*[10]", 10, p_s_t, reflect.TypeOf([10]*s_0{})},
	} {
		typ, err := cmem.NewArrayType(table.rtype)
		if err != nil {
			t.Errorf(err.Error())
		}
		eq(t, table.name, typ.Name())
		eq(t, table.elem, typ.Elem())
		eq(t, uintptr(table.n)*table.elem.Size(), typ.Size())
		eq(t, table.n, typ.Len())
		eq(t, cmem.Array, typ.Kind())
	}
}

func TestNewSliceType(t *testing.T) {

	capSize := 2 * unsafe.Sizeof(reflect.SliceHeader{}.Cap)

	s_t, err := cmem.NewStructType(reflect.TypeOf(s_0{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	p_s_t, err := cmem.NewPointerType(reflect.TypeOf(&s_0{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, table := range []struct {
		name  string
		elem  cmem.Type
		rtype reflect.Type
	}{
		{"uint8[]", cmem.C_uint8, reflect.TypeOf([]uint8{})},
		{"uint16[]", cmem.C_uint16, reflect.TypeOf([]uint16{})},
		{"uint32[]", cmem.C_uint32, reflect.TypeOf([]uint32{})},
		{"uint64[]", cmem.C_uint64, reflect.TypeOf([]uint64{})},
		{"int8[]", cmem.C_int8, reflect.TypeOf([]int8{})},
		{"int16[]", cmem.C_int16, reflect.TypeOf([]int16{})},
		{"int32[]", cmem.C_int32, reflect.TypeOf([]int32{})},
		{"int64[]", cmem.C_int64, reflect.TypeOf([]int64{})},

		{"float[]", cmem.C_float, reflect.TypeOf([]float32{})},
		{"double[]", cmem.C_double, reflect.TypeOf([]float64{})},

		{"s_0[]", s_t, reflect.TypeOf([]s_0{})},
		{"s_0*[]", p_s_t, reflect.TypeOf([]*s_0{})},
	} {
		typ, err := cmem.NewSliceType(table.rtype)
		if err != nil {
			t.Errorf(err.Error())
		}
		eq(t, table.name, typ.Name())
		eq(t, table.elem, typ.Elem())
		eq(t, capSize+cmem.C_pointer.Size(), typ.Size())
		//eq(t, table.n, typ.Len())
		eq(t, cmem.Slice, typ.Kind())
	}
}

func TestNewPointerType(t *testing.T) {
	s_t, err := cmem.NewStructType(reflect.TypeOf(s_0{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	p_s_t, err := cmem.NewPointerType(reflect.TypeOf(&s_0{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, table := range []struct {
		name  string
		elem  cmem.Type
		rtype reflect.Type
	}{
		{"int8*", cmem.C_int8, reflect.PtrTo(reflect.TypeOf(int8(0)))},
		{"int16*", cmem.C_int16, reflect.PtrTo(reflect.TypeOf(int16(0)))},
		{"int32*", cmem.C_int32, reflect.PtrTo(reflect.TypeOf(int32(0)))},
		{"int64*", cmem.C_int64, reflect.PtrTo(reflect.TypeOf(int64(0)))},
		{"uint8*", cmem.C_uint8, reflect.PtrTo(reflect.TypeOf(uint8(0)))},
		{"uint16*", cmem.C_uint16, reflect.PtrTo(reflect.TypeOf(uint16(0)))},
		{"uint32*", cmem.C_uint32, reflect.PtrTo(reflect.TypeOf(uint32(0)))},
		{"uint64*", cmem.C_uint64, reflect.PtrTo(reflect.TypeOf(uint64(0)))},

		{"float*", cmem.C_float, reflect.PtrTo(reflect.TypeOf(float32(0)))},
		{"double*", cmem.C_double, reflect.PtrTo(reflect.TypeOf(float64(0)))},

		{"s_0*", s_t, reflect.PtrTo(reflect.TypeOf(s_0{}))},
		{"s_0**", p_s_t, reflect.PtrTo(reflect.TypeOf(&s_0{}))},
	} {
		typ, err := cmem.NewPointerType(table.rtype)
		if err != nil {
			t.Errorf(err.Error())
		}
		eq(t, table.name, typ.Name())
		eq(t, table.elem, typ.Elem())
		eq(t, cmem.C_pointer.Size(), typ.Size())
		eq(t, cmem.Ptr, typ.Kind())
	}
}

// EOF
