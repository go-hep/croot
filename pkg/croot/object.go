package croot

// #include "croot/croot.h"
import "C"

type Object struct {
	o C.CRoot_Object
}

// EOF
