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


typedef signed char                _go_reflex_int8;
typedef struct { char *p; int n; } _go_reflex_String;
//extern void
//·_Cfunc_GoString(_go_reflex_int8 *p, _go_reflex_String s);

void
_go_reflex_gostring_ctor_stub(void *retaddr, void *mem, void *args, void *ctx)
 {
 if (retaddr) {
 } else {
 //_go_reflex_String *str = (_go_reflex_String*)realloc(mem, 16);
 void *dummy = (_go_reflex_String*)realloc(mem, 16);
 if (dummy) {}
 }
 
 }

 void*
 _get_go_reflex_gostring_ctor_stub() { return &_go_reflex_gostring_ctor_stub; }

*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type gostring struct {
	reflect.StringHeader
	Cap int
}

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
var reflexed_types map[reflect.Type]*ReflexType

func init() {
	reflexed_types = make(map[reflect.Type]*ReflexType)
}

func RegisterType(v interface{}) {
	t := reflect.ValueOf(v).Type()
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

// helper function to create a Reflex::Type from a go.reflect.Type
func genreflex(t reflect.Type) {
	fmt.Printf("::genreflex[%v]...\n", t)
	_,ok := reflexed_types[t]
	if ok {
		// already processed...
		return
	}

	var rflx_type *ReflexType = nil

	switch t.Kind() {
	case reflect.Bool:
		// no-op

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// no-op

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// no-op

	case reflect.Uintptr:
		// no-op

	case reflect.Float32, reflect.Float64:
		// noop

	case reflect.Complex64, reflect.Complex128:
		// no-op ?

	case reflect.Array:
		genreflex(t.Elem())
		//FIXME

	case reflect.Chan:
		panic(fmt.Sprintf("cannot handle Chan-kind [%s]", t.Name()))

	case reflect.Func:
		panic(fmt.Sprintf("cannot handle Func-kind [%s]", t.Name()))

	case reflect.Interface:
		panic(fmt.Sprintf("cannot handle Interface-kind [%s]", t.Name()))

	case reflect.Map:
		panic(fmt.Sprintf("cannot handle Map-kind [%s]", t.Name()))

	case reflect.Ptr:
		genreflex(t.Elem())

	case reflect.Slice:
		genreflex(t.Elem())

	case reflect.String:
		//FIXME

	case reflect.Struct:
		rflx_type = genreflex_struct(t)

	default:
		panic(fmt.Sprintf("unhandled type [%s]", t.Name()))
	}

	if rflx_type != nil {
		reflexed_types[t] = rflx_type
	}
	fmt.Printf("::genreflex[%v]...[done]\n", t)
}

// helper function to create a Reflex::Class-type from a go.struct
func genreflex_struct(t reflect.Type) *ReflexType {
	tname := t.Name()
	full_name := to_cxx_name(t)
	fmt.Printf("::genreflex_struct[%s]...\n", full_name)

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
	fmt.Printf(":: %s-size: %d\n", rt.Name(), rt.SizeOf())
	fmt.Printf(":: %s-mbrs: %d\n", rt.Name(), rt.DataMemberSize(Reflex_INHERITEDMEMBERS_NO))
	fmt.Printf("::genreflex_struct[%s]...[done]\n", full_name)
	return rt
}

// return a *croot.ReflexType from a reflect.Type one
func rflx_type_from(t reflect.Type) *ReflexType {
	var rflx *ReflexType = nil
	rflx,ok := reflexed_types[t]
	if ok {
		// already processed...
		return rflx
	}
	rflx = nil
	switch t.Kind() {
	case reflect.Bool:
		rflx = ReflexType_ByName("bool")

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

	case reflect.Complex64:
		rflx = ReflexType_ByName("float complex")

	case reflect.Complex128:
		rflx = ReflexType_ByName("double complex")

	case reflect.Array:
		rflx = NewReflexArrayBuilder(rflx_type_from(t.Elem()), t.Len())

	case reflect.Ptr:
		rflx = NewReflexPointerBuilder(rflx_type_from(t.Elem()))

		//case reflect.Slice:
	case reflect.String:
		ty_str := reflect.TypeOf(gostring{})
		genreflex_struct(ty_str)
		rflx = ReflexType_ByName(ty_str.Name())

	case reflect.Struct:
		genreflex_struct(t)
		rflx = ReflexType_ByName(t.Name())

	case reflect.UnsafePointer:
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