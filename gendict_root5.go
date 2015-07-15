// +build root5

package croot

/*
 #include "croot/croot.h"

 #include <stdlib.h>
 #include <string.h>
#include <stdio.h>
 static
 void
 _go_reflex_dummy_ctor_stub(void *retaddr, void *mem, void *args, void *ctx)
 {
 //printf("::go-reflex-dummy-ctor %p %p %p %p\n", retaddr, mem, args, ctx);
 //abort();
 }

 static
 void*
 _get_go_reflex_dummy_ctor_stub() { return &_go_reflex_dummy_ctor_stub; }

 static
 void
 _go_reflex_dummy_dtor_stub(void *retaddr, void *mem, void *args, void *ctx)
 {
 //printf("::go-reflex-dummy-dtor %p %p %p %p\n", retaddr, mem, args, ctx);
 //abort();
 }

 static
 void*
 _get_go_reflex_dummy_dtor_stub() { return &_go_reflex_dummy_dtor_stub; }

*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// the int ROOT uses for [Len] of C-arrays
// FIXME: this should be just int or int64 with ROOT-6.
type croot_int int32

var (
	_c_pointer_sz   = reflect.TypeOf(uintptr(0)).Size()
	_c_croot_int_sz = reflect.TypeOf(croot_int(0)).Size()
)

// map of already translated-to-Reflex types
var reflexed_types map[string]*ReflexType

func init() {
	reflexed_types = make(map[string]*ReflexType)
}

// helper function to create a Reflex::Type from a go.reflect.Type
func gendict(t reflect.Type) {
	//fmt.Printf("::gendict[%v]...\n", t)
	_, ok := reflexed_types[reflect_name2rflx(t)]
	if ok {
		// already processed...
		//fmt.Printf("::gendict[%v]... (already processed)\n", t)
		return
	}

	var rflx_type *ReflexType = nil

	switch t.Kind() {
	// case reflect.Bool:
	// 	// no-op

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// no-op

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// no-op

	// case reflect.Uintptr:
	// 	// no-op

	case reflect.Float32, reflect.Float64 /*, reflect.Float128*/ :
		// noop

	// case reflect.Complex64, reflect.Complex128:
	// 	// no-op ?

	case reflect.Array:
		gendict(t.Elem())
		rflx_type = rflx_type_from(t)

	// case reflect.Chan:
	// 	panic(fmt.Sprintf("cannot handle Chan-kind [%s]", t.Name()))

	// case reflect.Func:
	// 	panic(fmt.Sprintf("cannot handle Func-kind [%s]", t.Name()))

	// case reflect.Interface:
	// 	panic(fmt.Sprintf("cannot handle Interface-kind [%s]", t.Name()))

	// case reflect.Map:
	// 	panic(fmt.Sprintf("cannot handle Map-kind [%s]", t.Name()))

	case reflect.Ptr:
		//fmt.Printf("gendict-ptr...\n")
		gendict(t.Elem())

	case reflect.Slice:
		gendict(t.Elem())
		rflx_type = gendict_slice(t)

	case reflect.String:
		rflx_type = ReflexType_ByName("golang::gostring")

	case reflect.Struct:
		rflx_type = gendict_struct(t)

	default:
		panic(fmt.Sprintf("unhandled type [%s]", t.Name()))
	}

	if rflx_type != nil {
		reflexed_types[reflect_name2rflx(t)] = rflx_type
	}
	//fmt.Printf("::gendict[%v]...[done]\n", t)
}

// helper function to create a Reflex::Class-type from a go.struct
func gendict_struct(t reflect.Type) *ReflexType {
	tname := t.Name()
	full_name := to_cxx_name(t)
	//fmt.Printf("::gendict_struct[%s]...\n", full_name)

	bldr := NewReflexClassBuilder(
		//FIXME: generate namespaces for each containing package
		//       mentionned in 'full_name'
		full_name,
		t.Size(),
		uint32(Reflex_PUBLIC|Reflex_ARTIFICIAL),
		Reflex_STRUCT)

	nfields := t.NumField()
	for i := 0; i < nfields; i++ {
		f := t.Field(i)
		offset := f.Offset
		f_name := f.Name
		gendict(f.Type)
		bldr.AddDataMember(
			rflx_type_from(f.Type),
			f_name,
			offset,
			uint32(Reflex_PUBLIC))
	}
	ty_void := ReflexType_ByName("void")
	sz := C.size_t(t.Size())

	ty_ctor := NewReflexFunctionTypeBuilder(ty_void)
	stub_fct_ctor := (ReflexStubFunction)(C._get_go_reflex_dummy_ctor_stub())
	bldr.AddFunctionMember(
		ty_ctor,
		tname,
		stub_fct_ctor,
		unsafe.Pointer(&sz),
		uint32(Reflex_PUBLIC|Reflex_CONSTRUCTOR))

	ty_dtor := NewReflexFunctionTypeBuilder(ty_void)
	stub_fct_dtor := (ReflexStubFunction)(C._get_go_reflex_dummy_dtor_stub())
	bldr.AddFunctionMember(
		ty_dtor,
		"~"+tname,
		stub_fct_dtor,
		nil,
		uint32(Reflex_PUBLIC|Reflex_DESTRUCTOR))

	bldr.Delete()
	rt := ReflexType_ByName(tname)
	// fmt.Printf(":: reflect-siz: %d\n", t.Size())
	// fmt.Printf(":: %s-size: %d\n", rt.Name(), rt.SizeOf())
	// fmt.Printf(":: %s-mbrs: %d\n", rt.Name(), rt.DataMemberSize(Reflex_INHERITEDMEMBERS_NO))
	//fmt.Printf("::gendict_struct[%s]...[done]\n", full_name)
	return rt
}

// helper function to create a Reflex::Class-type from a go.struct
func gendict_slice(t reflect.Type) *ReflexType {
	tname := reflect_name2rflx(t)
	full_name := tname
	// full_name = "golang::goslice<double>"
	// fmt.Printf("::gendict_slice[%s]...\n", full_name)
	// {
	// 	rt := ReflexType_ByName(full_name)
	// 	fmt.Printf(":: reflect-siz: %d\n", t.Size())
	// 	fmt.Printf(":: %s-size: %d\n", rt.Name(), rt.SizeOf())
	// 	fmt.Printf(":: %s-mbrs: %d\n", rt.Name(), rt.DataMemberSize(Reflex_INHERITEDMEMBERS_NO))
	// 	fmt.Printf("::gendict_slice[%s]...[done]\n", full_name)
	// 	return rt
	// }
	bldr := NewReflexClassBuilder(
		//FIXME: generate namespaces for each containing package
		//       mentionned in 'full_name'
		full_name,
		t.Size(),
		uint32(Reflex_PUBLIC|Reflex_ARTIFICIAL),
		Reflex_STRUCT)

	ty_int32_t := ReflexType_ByName("int32_t")
	offset := uintptr(0)
	bldr.AddDataMember(
		ty_int32_t,
		"Len",
		offset,
		uint32(Reflex_PUBLIC),
	)
	offset += _c_croot_int_sz

	bldr.AddDataMember(
		ty_int32_t,
		"Cap",
		offset,
		uint32(Reflex_PUBLIC),
	)
	offset += _c_croot_int_sz

	bldr.AddDataMember(
		rflx_type_from(reflect.PtrTo(t.Elem())),
		"Data",
		offset,
		uint32(Reflex_PUBLIC),
	)
	bldr.AddProperty("comment", "[Len]")

	ty_void := ReflexType_ByName("void")
	sz := C.size_t(t.Size())

	ty_ctor := NewReflexFunctionTypeBuilder(ty_void)
	stub_fct_ctor := (ReflexStubFunction)(C._get_go_reflex_dummy_ctor_stub())
	bldr.AddFunctionMember(
		ty_ctor,
		tname,
		stub_fct_ctor,
		unsafe.Pointer(&sz),
		uint32(Reflex_PUBLIC|Reflex_CONSTRUCTOR))

	ty_dtor := NewReflexFunctionTypeBuilder(ty_void)
	stub_fct_dtor := (ReflexStubFunction)(C._get_go_reflex_dummy_dtor_stub())
	bldr.AddFunctionMember(
		ty_dtor,
		"~"+tname,
		stub_fct_dtor,
		nil,
		uint32(Reflex_PUBLIC|Reflex_DESTRUCTOR))

	bldr.Delete()
	rt := ReflexType_ByName(tname)
	// fmt.Printf(":: reflect-siz: %d\n", t.Size())
	// fmt.Printf(":: %s-size: %d\n", rt.Name(), rt.SizeOf())
	// fmt.Printf(":: %s-mbrs: %d\n", rt.Name(), rt.DataMemberSize(Reflex_INHERITEDMEMBERS_NO))
	// fmt.Printf("::gendict_slice[%s]...[done]\n", full_name)
	return rt
}

// return a *croot.ReflexType from a reflect.Type one
func rflx_type_from(t reflect.Type) *ReflexType {
	var rflx *ReflexType = nil
	rflx, ok := reflexed_types[reflect_name2rflx(t)]
	if ok {
		// already processed...
		return rflx
	}
	rflx = nil
	switch t.Kind() {
	// case reflect.Bool:
	// 	rflx = ReflexType_ByName("bool")

	case reflect.Int:
		rflx = ReflexType_ByName("int")

	case reflect.Int8:
		rflx = ReflexType_ByName("int8_t")

	case reflect.Int16:
		rflx = ReflexType_ByName("int16_t")

	case reflect.Int32:
		rflx = ReflexType_ByName("int32_t")

	case reflect.Int64:
		rflx = ReflexType_ByName("int64_t")

	case reflect.Uint:
		rflx = ReflexType_ByName("unsigned int")

	case reflect.Uint8:
		rflx = ReflexType_ByName("uint8_t")

	case reflect.Uint16:
		rflx = ReflexType_ByName("uint16_t")

	case reflect.Uint32:
		rflx = ReflexType_ByName("uint32_t")

	case reflect.Uint64:
		rflx = ReflexType_ByName("uint64_t")

	case reflect.Uintptr:
		rflx = ReflexType_ByName("uintptr_t")

	case reflect.Float32:
		rflx = ReflexType_ByName("float")

	case reflect.Float64:
		rflx = ReflexType_ByName("double")

	// case reflect.Float128:
	// 	rflx = ReflexType_ByName("long double")

	case reflect.Complex64:
		rflx = ReflexType_ByName("float complex")

	case reflect.Complex128:
		rflx = ReflexType_ByName("double complex")

	case reflect.Array:
		rflx = NewReflexArrayBuilder(rflx_type_from(t.Elem()), t.Len())

	case reflect.Ptr:
		rflx = NewReflexPointerBuilder(rflx_type_from(t.Elem()))

	case reflect.Slice:
		gendict_slice(t)
		rflx = ReflexType_ByName(reflect_name2rflx(t))

	case reflect.String:
		rflx = ReflexType_ByName("golang::gostring")

	case reflect.Struct:
		gendict_struct(t)
		rflx = ReflexType_ByName(t.Name())

	// case reflect.UnsafePointer:
	// 	rflx = NewReflexPointerBuilder(ReflexType_ByName("void"))

	default:
		panic(
			fmt.Sprintf(
				"no mapping reflex->reflect for type '%v' (kind=%s)",
				t,
				t.Kind().String()))
	}

	reflexed_types[reflect_name2rflx(t)] = rflx
	return rflx
}

func reflect_name2rflx(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Slice:
		return "golang::goslice<" + t.Elem().Name() + ">"
	case reflect.String:
		return "golang::gostring"
	case reflect.Array:
		return fmt.Sprintf("%v[%v]", reflect_name2rflx(t.Elem()), t.Len())
	default:
		return t.Name()
	}
	panic("unreachable")
}
