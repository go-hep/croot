package croot

// #include "croot/croot.h"
//
import "C"

import (
	"unsafe"
)

// TRandom
type Random interface {
	Object
	Gaus(mean, sigma float64) float64
	Rannorf() (a, b float32)
	Rannord() (a, b float64)
	Rndm(i int) float64
}

type randomImpl struct {
	c C.CRoot_Random
}

var GRandom Random = nil

func (r *randomImpl) Gaus(mean, sigma float64) float64 {
	val := C.CRoot_Random_Gaus(r.c, C.double(mean), C.double(sigma))
	return float64(val)
}

func (r *randomImpl) Rannorf() (a, b float32) {
	c_a := (*C.float)(unsafe.Pointer(&a))
	c_b := (*C.float)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannorf(r.c, c_a, c_b)
	return
}

func (r *randomImpl) Rannord() (a, b float64) {
	c_a := (*C.double)(unsafe.Pointer(&a))
	c_b := (*C.double)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannord(r.c, c_a, c_b)
	return
}

func (r *randomImpl) Rndm(i int) float64 {
	val := C.CRoot_Random_Rndm(r.c, C.int32_t(i))
	return float64(val)
}
