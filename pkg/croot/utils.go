package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
import "C"

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

// EOF
