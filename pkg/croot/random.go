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

type random_impl struct {
	c C.CRoot_Random
}

func (r *random_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(r.c)
}

func (r *random_impl) as_tobject() *object_impl {
	return &object_impl{r.cptr()}
}

func (r *random_impl) ClassName() string {
	return r.as_tobject().ClassName()
}

func (r *random_impl) Clone(opt Option) Object {
	return r.as_tobject().Clone(opt)
}

func (r *random_impl) FindObject(name string) Object {
	return r.as_tobject().FindObject(name)
}

func (r *random_impl) GetName() string {
	return r.as_tobject().GetName()
}

func (r *random_impl) GetTitle() string {
	return r.as_tobject().GetTitle()
}

func (r *random_impl) InheritsFrom(clsname string) bool {
	return r.as_tobject().InheritsFrom(clsname)
}

func (r *random_impl) Print(option Option) {
	r.as_tobject().Print(option)
}

var GRandom Random = nil

func (r *random_impl) Gaus(mean, sigma float64) float64 {
	val := C.CRoot_Random_Gaus(r.c, C.double(mean), C.double(sigma))
	return float64(val)
}

func (r *random_impl) Rannorf() (a, b float32) {
	c_a := (*C.float)(unsafe.Pointer(&a))
	c_b := (*C.float)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannorf(r.c, c_a, c_b)
	return
}

func (r *random_impl) Rannord() (a, b float64) {
	c_a := (*C.double)(unsafe.Pointer(&a))
	c_b := (*C.double)(unsafe.Pointer(&b))
	C.CRoot_Random_Rannord(r.c, c_a, c_b)
	return
}

func (r *random_impl) Rndm(i int) float64 {
	val := C.CRoot_Random_Rndm(r.c, C.int32_t(i))
	return float64(val)
}

func init() {
	GRandom = &random_impl{c: C.CRoot_gRandom}
	cnvmap["TRandom"] = func(o c_object) Object {
		return &random_impl{c: (C.CRoot_Random)(o.cptr())}
	}
}

// EOF
