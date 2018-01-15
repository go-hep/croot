package croot

import (
	"fmt"
	"reflect"
	"unsafe"

	"go-hep.org/x/cgo/croot/cmem"
)

// cgoSliceType is a binary compatible memory representation of a Go slice
type cgoSliceType struct {
	Len  int
	Cap  int
	Data unsafe.Pointer
}

// cgoStrinType is a binary compatible memory representation of a Go string
type cgoStringType struct {
	Len  int
	Data unsafe.Pointer
}

// goConverter translates back and forth b/w a Go value and its C counter-part
type goConverter interface {
	cnvToC(g reflect.Value, c cmem.Value) error
	cnvFromC(g reflect.Value, c cmem.Value) error

	//get_c_ptr() unsafe.Pointer
	//get_go_ptr() reflect.Value

	//get_c_addr() unsafe.Pointer
	//get_go_addr() reflect.Value
}

func newGoCnvFromC(typename string) (goConverter, error) {
	var cnv goConverter
	var err error

	return cnv, err
}

type cnvBaseType struct {
	cptr unsafe.Pointer
	gptr reflect.Value
}

type cnvBuiltinsType struct {
	//cnv_base_t
}

func (cnv cnvBuiltinsType) cnvToC(gptr reflect.Value, cptr cmem.Value) error {
	rt := gptr.Type()

	switch rt.Kind() {
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		cptr.SetValue(gptr.Elem())

	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		cptr.SetValue(gptr.Elem())

	case reflect.Float32, reflect.Float64:
		cptr.SetValue(gptr.Elem())

	default:
		return fmt.Errorf("croot.converter: conversion to C for kind=%v cannot be handled", rt.Kind())
	}
	return nil
}

func (cnv *cnvBuiltinsType) cnvFromC(gptr reflect.Value, cptr cmem.Value) error {
	rt := gptr.Type()

	switch rt.Kind() {
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		gptr.Set(cptr.GoValue())

	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		gptr.Set(cptr.GoValue())

	case reflect.Float32, reflect.Float64:
		gptr.Set(cptr.GoValue())

	default:
		return fmt.Errorf("croot.converter: conversion from C for kind=%v cannot be handled", rt.Kind())
	}
	return nil
}

type cnvStructType struct {
	cnvBaseType
	fields []goConverter
}

func (cnv *cnvStructType) cnvToC(gptr reflect.Value, cptr cmem.Value) error {
	for _, fcnv := range cnv.fields {
		err := fcnv.cnvToC(gptr, cptr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cnv *cnvStructType) cnvFromC(gptr reflect.Value, cptr cmem.Value) error {

	for i, fcnv := range cnv.fields {
		gField := gptr.Field(i)
		cfield := cptr.Field(i)
		err := fcnv.cnvFromC(gField, cfield)
		if err != nil {
			return err
		}
	}
	return nil
}

type cnvSliceType struct {
	cnvBaseType
	elmt goConverter
}

func (cnv *cnvSliceType) cnvToC(gptr reflect.Value, cptr cmem.Value) error {
	cptr.SetLen(gptr.Len())
	for i := 0; i < gptr.NumField(); i++ {
		gElmt := gptr.Index(i)
		celmt := cptr.Index(i)
		err := cnv.elmt.cnvToC(gElmt, celmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cnv *cnvSliceType) cnvFromC(gptr reflect.Value, cptr cmem.Value) error {
	gptr.SetLen(cptr.Len())
	for i := 0; i < gptr.NumField(); i++ {
		gElmt := gptr.Index(i)
		celmt := cptr.Index(i)
		err := cnv.elmt.cnvFromC(gElmt, celmt)
		if err != nil {
			return err
		}
	}
	return nil
}

// EOF
