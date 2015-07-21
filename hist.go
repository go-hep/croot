package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
//
import "C"

import (
	"reflect"
	"unsafe"
)

// H1F
type H1F interface {
	Object

	AddBinContent(bin int, weight float64)
	GetBinContent(bin int) float64
	SetBinContent(bin int, value float64)

	Fill(value, weight float64) int
	FillN(data [][2]float64)

	GetBin(bin int) float64
	GetBinCenter(bin int) float64
	GetBinError(bin int) float64
	GetBinErrorLow(bin int) float64
	GetBinErrorUp(bin int) float64
	GetBinWidth(bin int) float64

	GetEntries() float64
	GetMean() float64
	GetMeanError() float64
	GetRMS() float64
	GetRMSError() float64
}

func NewH1F(name, title string, nbins int, xlow, xup float64) H1F {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c := C.CRoot_H1F_new(
		c_name, c_title,
		C.int32_t(nbins), C.double(xlow), C.double(xup),
	)
	if c == nil {
		return nil
	}

	h := &h1FImpl{c: c}
	return h
}

func NewH1FFrom(name, title string, data []float64) H1F {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	val := reflect.ValueOf(&data)
	slice := (*reflect.SliceHeader)(unsafe.Pointer(val.Pointer()))

	nbins := C.int32_t(len(data))
	c_data := (*C.double)(unsafe.Pointer(slice.Data))
	c := C.CRoot_H1F_new2(c_name, c_title, nbins, c_data)
	if c == nil {
		return nil
	}
	h := &h1FImpl{c: c}
	return h
}

type h1FImpl struct {
	c C.CRoot_H1F
}

// -- H1F interface impl --

func (h *h1FImpl) AddBinContent(bin int, weight float64) {
	C.CRoot_H1F_AddBinContent(h.c, C.int32_t(bin), C.double(weight))
}

func (h *h1FImpl) GetBinContent(bin int) float64 {
	o := C.CRoot_H1F_GetBinContent(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) SetBinContent(bin int, value float64) {
	C.CRoot_H1F_SetBinContent(h.c, C.int32_t(bin), C.double(value))
}

func (h *h1FImpl) Fill(value, weight float64) int {
	o := C.CRoot_H1F_Fill(h.c, C.double(value), C.double(weight))
	return int(o)
}

func (h *h1FImpl) FillN(data [][2]float64) {
	x := make([]float64, len(data))
	w := make([]float64, len(data))
	for i := range data {
		x[i] = data[i][0]
		w[i] = data[i][1]
	}

	x_val := reflect.ValueOf(&x)
	x_slice := (*reflect.SliceHeader)(unsafe.Pointer(x_val.Pointer()))
	c_x := (*C.double)(unsafe.Pointer(x_slice.Data))

	w_val := reflect.ValueOf(&w)
	w_slice := (*reflect.SliceHeader)(unsafe.Pointer(w_val.Pointer()))
	c_w := (*C.double)(unsafe.Pointer(w_slice.Data))

	const stride = 1
	C.CRoot_H1F_FillN(h.c, C.int32_t(len(data)), c_x, c_w, stride)
}

func (h *h1FImpl) GetBin(bin int) float64 {
	o := C.CRoot_H1F_GetBin(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) GetBinCenter(bin int) float64 {
	o := C.CRoot_H1F_GetBinCenter(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) GetBinError(bin int) float64 {
	o := C.CRoot_H1F_GetBinError(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) GetBinErrorLow(bin int) float64 {
	o := C.CRoot_H1F_GetBinErrorLow(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) GetBinErrorUp(bin int) float64 {
	o := C.CRoot_H1F_GetBinErrorUp(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) GetBinWidth(bin int) float64 {
	o := C.CRoot_H1F_GetBinWidth(h.c, C.int32_t(bin))
	return float64(o)
}

func (h *h1FImpl) GetEntries() float64 {
	o := C.CRoot_H1F_GetEntries(h.c)
	return float64(o)
}

func (h *h1FImpl) GetMean() float64 {
	o := C.CRoot_H1F_GetMean(h.c)
	return float64(o)
}

func (h *h1FImpl) GetMeanError() float64 {
	o := C.CRoot_H1F_GetMeanError(h.c)
	return float64(o)
}

func (h *h1FImpl) GetRMS() float64 {
	o := C.CRoot_H1F_GetRMS(h.c)
	return float64(o)
}

func (h *h1FImpl) GetRMSError() float64 {
	o := C.CRoot_H1F_GetRMSError(h.c)
	return float64(o)
}

// EOF
