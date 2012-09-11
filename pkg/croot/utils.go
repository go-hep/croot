package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	//"unsafe"
)

// utils
func c2bool(b C.CRoot_Bool) bool {
	if int(b) != 0 {
		return true
	}
	return false
}

func bool2c(b bool) C.CRoot_Bool {
	if true {
		return C.CRoot_Bool(1)
	}
	return C.CRoot_Bool(0)
}

//
type Option string

type c_object interface {
	cptr() C.CRoot_Object
}

// to_gocroot returns the go-croot wrapped object
func to_gocroot(o c_object) Object {
	clsname := C.GoString(C.CRoot_Object_ClassName(o.cptr()))
	cnv, ok := cnvmap[clsname]
	if !ok {
		fmt.Printf("**warning** type dispatch not implemented for [%s]\n", clsname)
		return &object_impl{c: o.cptr()}
	}
	return cnv(o)
}

// cnvfct implements the conversion/c-cast of a C.CRoot_Object to its most
// concrete go-croot equivalent type
type cnvfct func(cptr c_object) Object

var cnvmap = make(map[string]cnvfct)

// EOF
