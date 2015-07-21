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
	err := cling_gendict(code)
	if err != nil {
		panic(fmt.Errorf(
			"croot: failed to generate dictionary for [%s]",
			t.Name(),
		))
	}
	clinged_types[t] = struct{}{}
}

func cling_gendict(code string) error {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))

	ok := c2bool(C.CRoot_Interpreter_LoadText(C.CRoot_gInterpreter, ccode))
	if !ok {
		return fmt.Errorf("croot: could not generate dictionary")
	}
	return nil

}

func init_cling() {
	code := `
#include <stdint.h> // for intXXX_t
`
	err := cling_gendict(code)
	if err != nil {
		panic("croot: could not initialize CLing interpreter")
	}
}
