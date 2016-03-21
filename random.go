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

var GRandom Random = &randomImpl{C.CRoot_gRandom}

func (r *randomImpl) Gaus(mean, sigma float64) float64 {
	val := C.CRoot_Random_Gaus(r.c, C.double(mean), C.double(sigma))
	return float64(val)
}

func (r *randomImpl) Rannorf() (a, b float32) {
	ca := (*C.float)(unsafe.Pointer(&a))
	cb := (*C.float)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannorf(r.c, ca, cb)
	return
}

func (r *randomImpl) Rannord() (a, b float64) {
	ca := (*C.double)(unsafe.Pointer(&a))
	cb := (*C.double)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannord(r.c, ca, cb)
	return
}

func (r *randomImpl) Rndm(i int) float64 {
	val := C.CRoot_Random_Rndm(r.c, C.int32_t(i))
	return float64(val)
}
