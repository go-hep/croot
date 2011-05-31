package croot

/*
 #cgo LDFLAGS: -lcroot
 #include "croot.h"

 #include <stdlib.h>
 #include <string.h>


 void
 _go_reflex_dummy_ctor_stub(void *retaddr, void *mem, void *args, void *ctx)
 {}

 void*
 _get_go_reflex_dummy_ctor_stub() { return &_go_reflex_dummy_ctor_stub; }

 void
 _go_reflex_dummy_dtor_stub(void *retaddr, void *mem, void *args, void *ctx)
 {}

 void*
 _get_go_reflex_dummy_dtor_stub() { return &_go_reflex_dummy_dtor_stub; }

*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"bitbucket.org/binet/go-ctypes/pkg/ctypes"
)

type ctor_fct func(retaddr, mem, args, ctx unsafe.Pointer)
var ctors []*ctor_fct
//export GoCRoot_make_ctor
func make_ctor(sz uintptr) *ctor_fct {
	fct := func(retaddr, mem, args, ctx unsafe.Pointer) {
		fmt.Printf("--ctor[%d] [%v] [%v] [%v] [%v]...\n",
			sz, retaddr, mem, args, ctx)
	}
	ctor := (*ctor_fct)(&fct)
	ctors = append(ctors, ctor)
	return ctor
}

// map of already translated-to-Reflex types
var reflexed_types map[ctypes.Type]*ReflexType

func init() {
	reflexed_types = make(map[ctypes.Type]*ReflexType)
}

func RegisterType(v interface{}) {
	t := ctypes.ValueOf(v).Type()
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

// helper function to create a Reflex::Type from a go.ctypes.Type
func genreflex(t ctypes.Type) {
	//fmt.Printf("::genreflex[%v]...\n", t)
	_,ok := reflexed_types[t]
	if ok {
		// already processed...
		return
	}

	var rflx_type *ReflexType = nil

	switch t.Kind() {
	case ctypes.Bool:
		// no-op

	case ctypes.Int, ctypes.Int8, ctypes.Int16, ctypes.Int32, ctypes.Int64:
		// no-op

	case ctypes.Uint, ctypes.Uint8, ctypes.Uint16, ctypes.Uint32, ctypes.Uint64:
		// no-op

	case ctypes.Uintptr:
		// no-op

	case ctypes.Float32, ctypes.Float64:
		// noop

	case ctypes.Complex64, ctypes.Complex128:
		// no-op ?

	case ctypes.Array:
		genreflex(t.Elem())

	// case ctypes.Chan:
	// 	panic(fmt.Sprintf("cannot handle Chan-kind [%s]", t.Name()))

	// case ctypes.Func:
	// 	panic(fmt.Sprintf("cannot handle Func-kind [%s]", t.Name()))

	// case ctypes.Interface:
	// 	panic(fmt.Sprintf("cannot handle Interface-kind [%s]", t.Name()))

	// case ctypes.Map:
	// 	panic(fmt.Sprintf("cannot handle Map-kind [%s]", t.Name()))

	case ctypes.Ptr:
		genreflex(t.Elem())

	case ctypes.Slice:
		genreflex(t.Elem())

	case ctypes.String:
		//FIXME

	case ctypes.Struct:
		rflx_type = genreflex_struct(t)

	default:
		panic(fmt.Sprintf("unhandled type [%s]", t.Name()))
	}

	if rflx_type != nil {
		reflexed_types[t] = rflx_type
	}
	//fmt.Printf("::genreflex[%v]...[done]\n", t)
}

// helper function to create a Reflex::Class-type from a go.struct
func genreflex_struct(t ctypes.Type) *ReflexType {
	tname := t.Name()
	full_name := to_cxx_name(t.GoType())
	//fmt.Printf("::genreflex_struct[%s]...\n", full_name)

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
		uint32(Reflex_PUBLIC | Reflex_CONSTRUCTOR))

	ty_dtor := NewReflexFunctionTypeBuilder(ty_void)
	stub_fct_dtor := (ReflexStubFunction)(C._get_go_reflex_dummy_dtor_stub())
	bldr.AddFunctionMember(
		ty_dtor,
		"~"+tname,
		stub_fct_dtor,
		nil,
		uint32(Reflex_PUBLIC | Reflex_DESTRUCTOR))

	bldr.Delete()
	rt := ReflexType_ByName(tname)
	//fmt.Printf(":: %s-size: %d\n", rt.Name(), rt.SizeOf())
	//fmt.Printf(":: %s-mbrs: %d\n", rt.Name(), rt.DataMemberSize(Reflex_INHERITEDMEMBERS_NO))
	//fmt.Printf("::genreflex_struct[%s]...[done]\n", full_name)
	return rt
}

// return a *croot.ReflexType from a ctypes.Type one
func rflx_type_from(t ctypes.Type) *ReflexType {
	var rflx *ReflexType = nil
	rflx,ok := reflexed_types[t]
	if ok {
		// already processed...
		return rflx
	}
	rflx = nil
	switch t.Kind() {
	case ctypes.Bool:
		rflx = ReflexType_ByName("bool")

	case ctypes.Int:
		rflx = ReflexType_ByName("int")

	case ctypes.Int8:
		rflx = ReflexType_ByName("int8_t")

	case ctypes.Int16:
		rflx = ReflexType_ByName("int16_t")

	case ctypes.Int32:
		rflx = ReflexType_ByName("int32_t")

	case ctypes.Int64:
		rflx = ReflexType_ByName("int64_t")

	case ctypes.Uint:
		rflx = ReflexType_ByName("unsigned int")

	case ctypes.Uint8:
		rflx = ReflexType_ByName("uint8_t")

	case ctypes.Uint16:
		rflx = ReflexType_ByName("uint16_t")

	case ctypes.Uint32:
		rflx = ReflexType_ByName("uint32_t")

	case ctypes.Uint64:
		rflx = ReflexType_ByName("uint64_t")

	case ctypes.Uintptr:
		rflx = ReflexType_ByName("uintptr_t")

	case ctypes.Float32:
		rflx = ReflexType_ByName("float")

	case ctypes.Float64:
		rflx = ReflexType_ByName("double")

	case ctypes.Complex64:
		rflx = ReflexType_ByName("float complex")

	case ctypes.Complex128:
		rflx = ReflexType_ByName("double complex")

	case ctypes.Array:
		rflx = NewReflexArrayBuilder(rflx_type_from(t.Elem()), t.Len())

	case ctypes.Ptr:
		rflx = NewReflexPointerBuilder(rflx_type_from(t.Elem()))

	// case ctypes.Slice:
	// 	rflx 

	case ctypes.String:
		rflx = ReflexType_ByName("char*")

	case ctypes.Struct:
		genreflex_struct(t)
		rflx = ReflexType_ByName(t.Name())

	case ctypes.UnsafePointer:
		rflx = NewReflexPointerBuilder(ReflexType_ByName("void"))

	default:
		panic(
			fmt.Sprintf(
			"no mapping reflex->reflect for type '%v' (kind=%s)",
			t,
			t.Kind().String()))
	}

	reflexed_types[t] = rflx
	return rflx
}