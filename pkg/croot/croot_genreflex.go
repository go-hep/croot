package croot

/*
 #cgo LDFLAGS: -lcroot
 #include "croot.h"

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

	"github.com/sbinet/go-ffi/pkg/ffi"
)

type ctor_fct func(retaddr, mem, args, ctx unsafe.Pointer)

var ctors []*ctor_fct

//export GoCRoot_make_ctor
func GoCRoot_make_ctor(sz uintptr) *ctor_fct {
	fct := func(retaddr, mem, args, ctx unsafe.Pointer) {
		fmt.Printf("--ctor[%d] [%v] [%v] [%v] [%v]...\n",
			sz, retaddr, mem, args, ctx)
	}
	ctor := (*ctor_fct)(&fct)
	ctors = append(ctors, ctor)
	return ctor
}

// map of already translated-to-Reflex types
var reflexed_types map[string]*ReflexType

func init() {
	reflexed_types = make(map[string]*ReflexType)
}

// RegisterType declares the (equivalent) C-layout of value v to ROOT so
// values of the same type than v can be written out to ROOT files
func RegisterType(v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	t := ffi.ValueOf(rv.Interface()).Type()
	if t.Kind() == ffi.Ptr {
		t = t.Elem()
	}
	//fmt.Printf("registering [%s] (sz:%d)...\n",t, t.Size())
	genreflex(t)
}

func follow_ptr(v reflect.Value) reflect.Value {
	for {
		switch v.Kind() {
		case reflect.Ptr:
			if v.Elem().Kind() == reflect.Ptr {
				v = v.Elem()
			} else {
				return v
			}
		default:
			return v
		}
	}
	return v
}

func to_cxx_name(t reflect.Type) string {
	//return fmt.Sprintf("::golang::%s::%s", t.PkgPath(), t.Name())
	return t.Name()
}

// helper function to create a Reflex::Type from a go.ffi.Type
func genreflex(t ffi.Type) {
	//fmt.Printf("::genreflex[%v]...\n", t)
	_, ok := reflexed_types[t.Name()]
	if ok {
		// already processed...
		return
	}

	var rflx_type *ReflexType = nil

	switch t.Kind() {
	// case ffi.Bool:
	// 	// no-op

	case ffi.Int, ffi.Int8, ffi.Int16, ffi.Int32, ffi.Int64:
		// no-op

	case /*ffi.Uint,*/ ffi.Uint8, ffi.Uint16, ffi.Uint32, ffi.Uint64:
		// no-op

	// case ffi.Uintptr:
	// 	// no-op

	case ffi.Float, ffi.Double, ffi.LongDouble:
		// noop

	// case ffi.Complex64, ffi.Complex128:
	// 	// no-op ?

	case ffi.Array:
		genreflex(t.Elem())

	// case ffi.Chan:
	// 	panic(fmt.Sprintf("cannot handle Chan-kind [%s]", t.Name()))

	// case ffi.Func:
	// 	panic(fmt.Sprintf("cannot handle Func-kind [%s]", t.Name()))

	// case ffi.Interface:
	// 	panic(fmt.Sprintf("cannot handle Interface-kind [%s]", t.Name()))

	// case ffi.Map:
	// 	panic(fmt.Sprintf("cannot handle Map-kind [%s]", t.Name()))

	case ffi.Ptr:
		genreflex(t.Elem())

	case ffi.Slice:
		genreflex(t.Elem())

	case ffi.String:
		//FIXME

	case ffi.Struct:
		rflx_type = genreflex_struct(t)

	default:
		panic(fmt.Sprintf("unhandled type [%s]", t.Name()))
	}

	if rflx_type != nil {
		reflexed_types[t.Name()] = rflx_type
	}
	//fmt.Printf("::genreflex[%v]...[done]\n", t)
}

// helper function to create a Reflex::Class-type from a go.struct
func genreflex_struct(t ffi.Type) *ReflexType {
	tname := t.Name()
	if t.GoType() == nil {
		panic("no go-type for ffi.Type [" + t.Name() + "]")
	}
	full_name := to_cxx_name(t.GoType())
	// fmt.Printf("::genreflex_struct[%s]...\n", full_name)

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
	// fmt.Printf(":: ffi-siz: %d\n", t.Size())
	// fmt.Printf(":: %s-size: %d\n", rt.Name(), rt.SizeOf())
	// fmt.Printf(":: %s-mbrs: %d\n", rt.Name(), rt.DataMemberSize(Reflex_INHERITEDMEMBERS_NO))
	// fmt.Printf("::genreflex_struct[%s]...[done]\n", full_name)
	return rt
}

// return a *croot.ReflexType from a ffi.Type one
func rflx_type_from(t ffi.Type) *ReflexType {
	var rflx *ReflexType = nil
	rflx, ok := reflexed_types[t.Name()]
	if ok {
		// already processed...
		return rflx
	}
	rflx = nil
	switch t.Kind() {
	// case ffi.Bool:
	// 	rflx = ReflexType_ByName("bool")

	case ffi.Int:
		rflx = ReflexType_ByName("int")

	case ffi.Int8:
		rflx = ReflexType_ByName("int8_t")

	case ffi.Int16:
		rflx = ReflexType_ByName("int16_t")

	case ffi.Int32:
		rflx = ReflexType_ByName("int32_t")

	case ffi.Int64:
		rflx = ReflexType_ByName("int64_t")

	// case ffi.Uint:
	// 	rflx = ReflexType_ByName("unsigned int")

	case ffi.Uint8:
		rflx = ReflexType_ByName("uint8_t")

	case ffi.Uint16:
		rflx = ReflexType_ByName("uint16_t")

	case ffi.Uint32:
		rflx = ReflexType_ByName("uint32_t")

	case ffi.Uint64:
		rflx = ReflexType_ByName("uint64_t")

	// case ffi.Uintptr:
	// 	rflx = ReflexType_ByName("uintptr_t")

	case ffi.Float:
		rflx = ReflexType_ByName("float")

	case ffi.Double:
		rflx = ReflexType_ByName("double")

	case ffi.LongDouble:
		rflx = ReflexType_ByName("long double")

	// case ffi.Complex64:
	// 	rflx = ReflexType_ByName("float complex")

	// case ffi.Complex128:
	// 	rflx = ReflexType_ByName("double complex")

	case ffi.Array:
		rflx = NewReflexArrayBuilder(rflx_type_from(t.Elem()), t.Len())

	case ffi.Ptr:
		rflx = NewReflexPointerBuilder(rflx_type_from(t.Elem()))

	// case ffi.Slice:
	// 	rflx 

	case ffi.String:
		rflx = ReflexType_ByName("char*")

	case ffi.Struct:
		genreflex_struct(t)
		rflx = ReflexType_ByName(t.Name())

	// case ffi.UnsafePointer:
	// 	rflx = NewReflexPointerBuilder(ReflexType_ByName("void"))

	default:
		panic(
			fmt.Sprintf(
				"no mapping reflex->reflect for type '%v' (kind=%s)",
				t,
				t.Kind().String()))
	}

	reflexed_types[t.Name()] = rflx
	return rflx
}
