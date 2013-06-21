package cmem_test

import (
	//"fmt"
	"reflect"
	"testing"
	//"unsafe"

	"github.com/go-hep/croot/cmem"
)

func TestGetSetBuiltinValue(t *testing.T) {

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"int", cmem.C_int, int64(val)},
			{"int8", cmem.C_int8, int64(val)},
			{"int16", cmem.C_int16, int64(val)},
			{"int32", cmem.C_int32, int64(val)},
			{"int64", cmem.C_int64, int64(val)},
		} {
			cval := cmem.New(tt.t)
			eq(t, tt.n, cval.Type().Name())
			eq(t, tt.t.Kind(), cval.Kind())
			eq(t, reflect.Zero(reflect.TypeOf(tt.val)).Int(), cval.Int())
			cval.SetInt(val)
			eq(t, tt.val, cval.Int())
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"unsigned int", cmem.C_uint, uint64(val)},
			{"uint8", cmem.C_uint8, uint64(val)},
			{"uint16", cmem.C_uint16, uint64(val)},
			{"uint32", cmem.C_uint32, uint64(val)},
			{"uint64", cmem.C_uint64, uint64(val)},
		} {
			cval := cmem.New(tt.t)
			eq(t, tt.n, cval.Type().Name())
			eq(t, tt.t.Kind(), cval.Kind())
			eq(t, reflect.Zero(reflect.TypeOf(tt.val)).Uint(), cval.Uint())
			cval.SetUint(val)
			eq(t, tt.val, cval.Uint())
		}
	}

	{
		const val = -66.0
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"float", cmem.C_float, float64(val)},
			{"double", cmem.C_double, float64(val)},
			//FIXME: Go has no equivalent for long double...
			//{"long double", cmem.C_longdouble, float128(val)},
		} {
			cval := cmem.New(tt.t)
			eq(t, tt.n, cval.Type().Name())
			eq(t, tt.t.Kind(), cval.Kind())
			eq(t, reflect.Zero(reflect.TypeOf(tt.val)).Float(), cval.Float())
			cval.SetFloat(val)
			eq(t, tt.val, cval.Float())
		}
	}

	{
		const val = -66
		cval := cmem.New(cmem.C_int64)
		cptr := cval.Addr()
		cval.SetInt(val)
		eq(t, int64(val), cval.Int())
		eq(t, int64(val), cptr.Elem().Int())
		cval.SetInt(0)
		eq(t, int64(0), cptr.Elem().Int())
		cptr.Elem().SetInt(val)
		eq(t, int64(val), cval.Int())
		eq(t, int64(val), cptr.Elem().Int())

	}
}

func TestGetSetArrayValue(t *testing.T) {

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			len int
			t   cmem.Type
			val interface{}
		}{
			{"uint8[10]", 10, cmem.C_uint8, [10]uint8{}},
			{"uint16[10]", 10, cmem.C_uint16, [10]uint16{}},
			{"uint32[10]", 10, cmem.C_uint32, [10]uint32{}},
			{"uint64[10]", 10, cmem.C_uint64, [10]uint64{}},
		} {
			ctyp, err := cmem.NewArrayType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.New(ctyp)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.New(gtyp).Elem()
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
				gval.Index(i).SetUint(val)
				cval.Index(i).SetUint(val)
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
			}
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			len int
			t   cmem.Type
			val interface{}
		}{
			{"int8[10]", 10, cmem.C_int8, [10]int8{}},
			{"int16[10]", 10, cmem.C_int16, [10]int16{}},
			{"int32[10]", 10, cmem.C_int32, [10]int32{}},
			{"int64[10]", 10, cmem.C_int64, [10]int64{}},
		} {
			ctyp, err := cmem.NewArrayType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.New(ctyp)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.New(gtyp).Elem()
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
				gval.Index(i).SetInt(val)
				cval.Index(i).SetInt(val)
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
			}
		}
	}

	{
		const val = -66.2
		for _, tt := range []struct {
			n   string
			len int
			t   cmem.Type
			val interface{}
		}{
			{"float[10]", 10, cmem.C_float, [10]float32{}},
			{"double[10]", 10, cmem.C_double, [10]float64{}},
			// FIXME: go has no long double equivalent
			//{"long double[10]", 10, cmem.C_longdouble, [10]float128{}},
		} {
			ctyp, err := cmem.NewArrayType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.New(ctyp)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.New(gtyp).Elem()
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
				gval.Index(i).SetFloat(val)
				cval.Index(i).SetFloat(val)
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
			}
		}
	}

}

type struct_ssv struct {
	F1 uint16
	F2 [10]int32
	F3 int32
	F4 uint16
}

func TestGetSetStructValue(t *testing.T) {

	const val = 42
	arr10, err := cmem.NewArrayType(reflect.TypeOf([10]int32{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	ctyp, err := cmem.NewStructType(reflect.TypeOf(struct_ssv{}))
	eq(t, "struct_ssv", ctyp.Name())
	eq(t, cmem.Struct, ctyp.Kind())
	eq(t, 4, ctyp.NumField())

	cval := cmem.New(ctyp)
	eq(t, ctyp.Kind(), cval.Kind())
	eq(t, ctyp.NumField(), cval.NumField())
	eq(t, uint64(0), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(0), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(0), cval.Field(2).Int())
	eq(t, uint64(0), cval.Field(3).Uint())

	// set everything to 'val'
	cval.Field(0).SetUint(val)
	for i := 0; i < arr10.Len(); i++ {
		cval.Field(1).Index(i).SetInt(val)
	}
	cval.Field(2).SetInt(val)
	cval.Field(3).SetUint(val)

	// test values back
	eq(t, uint64(val), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(val), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(val), cval.Field(2).Int())
	eq(t, uint64(val), cval.Field(3).Uint())

	// test values back - by field name
	eq(t, uint64(val), cval.FieldByName("F1").Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(val), cval.FieldByName("F2").Index(i).Int())
	}
	eq(t, int64(val), cval.FieldByName("F3").Int())
	eq(t, uint64(val), cval.FieldByName("F4").Uint())
}

type struct_sswsv struct {
	F1 uint16
	F2 [10]int32
	F3 int32
	F4 uint16
	F5 []int32
}

func TestGetSetStructWithSliceValue(t *testing.T) {

	const val = 42
	arr10, err := cmem.NewArrayType(reflect.TypeOf([10]int32{}))
	if err != nil {
		t.Errorf(err.Error())
	}

	ctyp, err := cmem.NewStructType(reflect.TypeOf(struct_sswsv{}))
	eq(t, "struct_sswsv", ctyp.Name())
	eq(t, cmem.Struct, ctyp.Kind())
	eq(t, 5, ctyp.NumField())

	cval := cmem.New(ctyp)
	eq(t, ctyp.Kind(), cval.Kind())
	eq(t, ctyp.NumField(), cval.NumField())
	eq(t, uint64(0), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(0), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(0), cval.Field(2).Int())
	eq(t, uint64(0), cval.Field(3).Uint())
	eq(t, int(0), cval.Field(4).Len())
	eq(t, int(0), cval.Field(4).Len())

	goval := struct_sswsv{
		F1: val,
		F2: [10]int32{val, val, val, val, val,
			val, val, val, val, val},
		F3: val,
		F4: val,
		F5: make([]int32, 2, 3),
	}
	goval.F5[0] = val
	goval.F5[1] = val

	cval.SetValue(reflect.ValueOf(goval))

	eq(t, uint64(val), cval.Field(0).Uint())
	for i := 0; i < arr10.Len(); i++ {
		eq(t, int64(val), cval.Field(1).Index(i).Int())
	}
	eq(t, int64(val), cval.Field(2).Int())
	eq(t, uint64(val), cval.Field(3).Uint())
	eq(t, int(2), cval.Field(4).Len())
	// FIXME: should we get the 'cap' from go ?
	eq(t, int( /*3*/ 2), cval.Field(4).Cap())
	eq(t, int64(val), cval.Field(4).Index(0).Int())
	eq(t, int64(val), cval.Field(4).Index(1).Int())
}

func TestGetSetSliceValue(t *testing.T) {

	const sz = 10
	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"uint8[]", cmem.C_uint8, make([]uint8, sz)},
			{"uint16[]", cmem.C_uint16, make([]uint16, sz)},
			{"uint32[]", cmem.C_uint32, make([]uint32, sz)},
			{"uint64[]", cmem.C_uint64, make([]uint64, sz)},
		} {
			ctyp, err := cmem.NewSliceType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.MakeSlice(ctyp, sz, sz)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, gval.Len(), cval.Len())
			eq(t, int(sz), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
				gval.Index(i).SetUint(val)
				cval.Index(i).SetUint(val)
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
			}
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"int8[]", cmem.C_int8, make([]int8, sz)},
			{"int16[]", cmem.C_int16, make([]int16, sz)},
			{"int32[]", cmem.C_int32, make([]int32, sz)},
			{"int64[]", cmem.C_int64, make([]int64, sz)},
		} {
			ctyp, err := cmem.NewSliceType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.MakeSlice(ctyp, sz, sz)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, gval.Len(), cval.Len())
			eq(t, int(sz), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
				gval.Index(i).SetInt(val)
				cval.Index(i).SetInt(val)
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
			}
		}
	}

	{
		const val = -66.2
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"float[]", cmem.C_float, make([]float32, sz)},
			{"double[]", cmem.C_double, make([]float64, sz)},
			// FIXME: go has no long double equivalent
			//{"long double[]", cmem.C_longdouble, make([]float128, sz)}
		} {
			ctyp, err := cmem.NewSliceType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.MakeSlice(ctyp, sz, sz)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, gval.Len(), cval.Len())
			eq(t, int(sz), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
				gval.Index(i).SetFloat(val)
				cval.Index(i).SetFloat(val)
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
			}
		}
	}

	// now test if slices can automatically grow...
	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"uint8[]", cmem.C_uint8, make([]uint8, sz)},
			{"uint16[]", cmem.C_uint16, make([]uint16, sz)},
			{"uint32[]", cmem.C_uint32, make([]uint32, sz)},
			{"uint64[]", cmem.C_uint64, make([]uint64, sz)},
		} {
			ctyp, err := cmem.NewSliceType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.MakeSlice(ctyp, 0, 0)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, int(0), cval.Len())
			cval.SetValue(gval) // <---------
			eq(t, int(sz), cval.Len())
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
				gval.Index(i).SetUint(val)
				cval.Index(i).SetUint(val)
				eq(t, gval.Index(i).Uint(), cval.Index(i).Uint())
			}
		}
	}

	{
		const val = 42
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"int8[]", cmem.C_int8, make([]int8, sz)},
			{"int16[]", cmem.C_int16, make([]int16, sz)},
			{"int32[]", cmem.C_int32, make([]int32, sz)},
			{"int64[]", cmem.C_int64, make([]int64, sz)},
		} {
			ctyp, err := cmem.NewSliceType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.MakeSlice(ctyp, 0, 0)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, int(0), cval.Len())
			cval.SetValue(gval) // <---------
			eq(t, int(sz), cval.Len())
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
				gval.Index(i).SetInt(val)
				cval.Index(i).SetInt(val)
				eq(t, gval.Index(i).Int(), cval.Index(i).Int())
			}
		}
	}

	{
		const val = -66.2
		for _, tt := range []struct {
			n   string
			t   cmem.Type
			val interface{}
		}{
			{"float[]", cmem.C_float, make([]float32, sz)},
			{"double[]", cmem.C_double, make([]float64, sz)},
			// FIXME: go has no long double equivalent
			//{"long double[]", cmem.C_longdouble, make([]float128, sz)}
		} {
			ctyp, err := cmem.NewSliceType(reflect.ValueOf(tt.val).Type())
			if err != nil {
				t.Errorf(err.Error())
			}
			cval := cmem.MakeSlice(ctyp, 0, 0)
			eq(t, tt.n, cval.Type().Name())
			eq(t, ctyp.Kind(), cval.Kind())
			gtyp := reflect.TypeOf(tt.val)
			gval := reflect.MakeSlice(gtyp, sz, sz)
			eq(t, int(0), cval.Len())
			cval.SetValue(gval) // <---------
			eq(t, int(sz), cval.Len())
			eq(t, gval.Len(), cval.Len())
			for i := 0; i < gval.Len(); i++ {
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
				gval.Index(i).SetFloat(val)
				cval.Index(i).SetFloat(val)
				eq(t, gval.Index(i).Float(), cval.Index(i).Float())
			}
		}
	}
}

type struct_ints struct {
	F1 int8
	F2 int16
	F3 int32
	F4 int64
}

func TestValueOf(t *testing.T) {
	{
		const val = 42
		for _, v := range []interface{}{
			int(val),
			int8(val),
			int16(val),
			int32(val),
			int64(val),
		} {
			eq(t, int64(val), cmem.ValueOf(v).Int())
		}
	}

	{
		const val = 42
		for _, v := range []interface{}{
			uint(val),
			uint8(val),
			uint16(val),
			uint32(val),
			uint64(val),
		} {
			eq(t, uint64(val), cmem.ValueOf(v).Uint())
		}
	}
	{
		const val = 42.0
		for _, v := range []interface{}{
			float32(val),
			float64(val),
		} {
			eq(t, float64(val), cmem.ValueOf(v).Float())
		}
	}
	{
		const val = 42
		ctyp, err := cmem.NewStructType(reflect.TypeOf(struct_ints{}))
		if err != nil {
			t.Errorf(err.Error())
		}
		cval := cmem.New(ctyp)
		for i := 0; i < ctyp.NumField(); i++ {
			cval.Field(i).SetInt(int64(val))
			eq(t, int64(val), cval.Field(i).Int())
		}
		gval := struct_ints{val + 1, val + 1, val + 1, val + 1}
		rval := reflect.ValueOf(gval)
		eq(t, rval.NumField(), cval.NumField())
		for i := 0; i < ctyp.NumField(); i++ {
			eq(t, rval.Field(i).Int()-1, cval.Field(i).Int())
		}
		cval = cmem.ValueOf(gval)
		for i := 0; i < ctyp.NumField(); i++ {
			eq(t, rval.Field(i).Int(), cval.Field(i).Int())
		}
	}
}

// EOf
