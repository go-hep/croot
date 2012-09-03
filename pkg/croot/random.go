package croot

// #include "croot/croot.h"
//
import "C"

import (
	"unsafe"
)

// TRandom
type Random struct {
	r C.CRoot_Random
}

var GRandom *Random = nil

func (r *Random) Gaus(mean, sigma float64) float64 {
	val := C.CRoot_Random_Gaus(r.r, C.double(mean), C.double(sigma))
	return float64(val)
}

func (r *Random) Rannorf() (a, b float32) {
	c_a := (*C.float)(unsafe.Pointer(&a))
	c_b := (*C.float)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannorf(r.r, c_a, c_b)
	return
}

func (r *Random) Rannord() (a, b float64) {
	c_a := (*C.double)(unsafe.Pointer(&a))
	c_b := (*C.double)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannord(r.r, c_a, c_b)
	return
}

func (r *Random) Rndm(i int) float64 {
	val := C.CRoot_Random_Rndm(r.r, C.int32_t(i))
	return float64(val)
}

func init() {
	GRandom = &Random{r: C.CRoot_gRandom}
}

// EOF
