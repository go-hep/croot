package croot

// #include "croot/croot.h"
import "C"

// ObjArray
type ObjArray struct {
	c C.CRoot_ObjArray
}

// EOF
