package croot

// #include "croot/croot.h"
// #include <string.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"go-hep.org/x/cgo/croot/cgentype"
)

// RegisterType declares the (equivalent) C-layout of value v to ROOT so
// values of the same type than v can be written out to ROOT files
func RegisterType(v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	t := rv.Type()
	//fmt.Printf("registering [%s] (sz:%d)...\n",t, t.Size())
	gendict(t)
}

func followPtr(v reflect.Value) reflect.Value {
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
}

func toCxxName(t reflect.Type) string {
	//return fmt.Sprintf("::golang::%s::%s", t.PkgPath(), t.Name())
	return t.Name()
}

// map of already translated-to-ROOT types
var dicts map[reflect.Type]struct{}

func init() {
	dicts = make(map[reflect.Type]struct{})
	initROOTInterpreter()
}

func gendict(t reflect.Type) {
	if _, ok := dicts[t]; ok {
		return
	}

	switch t.Kind() {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		// no-op
		return

	case reflect.Uintptr, reflect.UnsafePointer:
		// TODO(sbinet)
		panic(fmt.Errorf(
			"croot: cannot generate dictionary for uintptr/unsafe.Pointer [%s]",
			t.Name(),
		))

	case reflect.Array:
		gendict(t.Elem())
		return

	case reflect.Map:
		panic(fmt.Errorf(
			"croot: cannot generate dictionary for map [%s]",
			t.Name(),
		))

	case reflect.Ptr:
		gendict(t.Elem())
		return

	case reflect.Slice:
		gendict(t.Elem())
		return

	case reflect.String:
		// no-op
		return

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			gendict(t.Field(i).Type)
		}

	case reflect.Interface:
		panic(fmt.Errorf(
			"croot: cannot generate dictionary for interface [%s]",
			t.Name(),
		))

	case reflect.Chan:
		panic(fmt.Errorf(
			"croot: cannot generate dictionary for channel [%s]", t.Name(),
		))
	}

	code := cgentype.Generate(t)
	err := genROOTDict(code)
	if err != nil {
		panic(fmt.Errorf(
			"croot: failed to generate dictionary for [%s]",
			t.Name(),
		))
	}
	dicts[t] = struct{}{}
}

func genROOTDict(code string) error {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))

	ok := c2bool(C.CRoot_Interpreter_LoadText(C.CRoot_gInterpreter, ccode))
	if !ok {
		return fmt.Errorf("croot: could not generate dictionary")
	}
	return nil

}
