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

func gendict(t reflect.Type) {
	code := cgentype.Generate(t)
	fmt.Printf("code=%v\n===\n", code)
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))

	ok := c2bool(C.CRoot_Interpreter_LoadText(C.CRoot_gInterpreter, c_code))
	if !ok {
		panic("oops")
	}
	panic("gendict not implemented for ROOT-6")
}
