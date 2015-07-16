// +build root6

package croot

// #include "croot/croot.h"
// #include <string.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/go-hep/croot/cgentype"
)

// map of already translated-to-Cling types
var clinged_types map[reflect.Type]struct{}

func init() {
	clinged_types = make(map[reflect.Type]struct{})
	init_cling()
}

func gendict(t reflect.Type) {
	if _, ok := clinged_types[t]; ok {
		return
	}

	switch t.Kind() {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Uintptr:
		// no-op
		return

	case reflect.Array:
		gendict(t.Elem())
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
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))

	ok := c2bool(C.CRoot_Interpreter_LoadText(C.CRoot_gInterpreter, c_code))
	if !ok {
		panic(fmt.Errorf(
			"croot: failed to generate dictionary for [%s]",
			t.Name(),
		))
	}
	clinged_types[t] = struct{}{}
}

func init_cling() {
	c_code := C.CString(`
#include <stdint.h>
`)
	defer C.free(unsafe.Pointer(c_code))

	ok := c2bool(C.CRoot_Interpreter_LoadText(C.CRoot_gInterpreter, c_code))
	if !ok {
		panic("croot: could not initialize CLing interpreter")
	}
}
