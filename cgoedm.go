package croot

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/go-hep/croot/cmem"
)

// cgoslice_t is a binary compatible memory representation of a Go slice
type cgoslice_t struct {
	Len  int
	Cap  int
	Data unsafe.Pointer
}

// cgoslice_t is a binary compatible memory representation of a Go string
type cgostring_t struct {
	Len  int
	Data unsafe.Pointer
}

// go_converter translates back and forth b/w a Go value and its C counter-part
type go_converter interface {
	cnv_to_c(g reflect.Value, c cmem.Value) error
	cnv_from_c(g reflect.Value, c cmem.Value) error

	//get_c_ptr() unsafe.Pointer
	//get_go_ptr() reflect.Value

	//get_c_addr() unsafe.Pointer
	//get_go_addr() reflect.Value
}

func new_go_cnv_from_c(typename string) (go_converter, error) {
	var cnv go_converter
	var err error

	return cnv, err
}

type cnv_base_t struct {
	cptr unsafe.Pointer
	gptr reflect.Value
}

type cnv_builtins_t struct {
	//cnv_base_t
}

func (cnv cnv_builtins_t) cnv_to_c(gptr reflect.Value, cptr cmem.Value) error {
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

func (cnv *cnv_builtins_t) cnv_from_c(gptr reflect.Value, cptr cmem.Value) error {
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

type cnv_struct_t struct {
	cnv_base_t
	fields []go_converter
}

func (cnv *cnv_struct_t) cnv_to_c(gptr reflect.Value, cptr cmem.Value) error {
	for _, fcnv := range cnv.fields {
		err := fcnv.cnv_to_c(gptr, cptr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cnv *cnv_struct_t) cnv_from_c(gptr reflect.Value, cptr cmem.Value) error {

	for i, fcnv := range cnv.fields {
		g_field := gptr.Field(i)
		c_field := cptr.Field(i)
		err := fcnv.cnv_from_c(g_field, c_field)
		if err != nil {
			return err
		}
	}
	return nil
}

type cnv_slice_t struct {
	cnv_base_t
	elmt go_converter
}

func (cnv *cnv_slice_t) cnv_to_c(gptr reflect.Value, cptr cmem.Value) error {
	cptr.SetLen(gptr.Len())
	for i := 0; i < gptr.NumField(); i++ {
		g_elmt := gptr.Index(i)
		c_elmt := cptr.Index(i)
		err := cnv.elmt.cnv_to_c(g_elmt, c_elmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cnv *cnv_slice_t) cnv_from_c(gptr reflect.Value, cptr cmem.Value) error {
	gptr.SetLen(cptr.Len())
	for i := 0; i < gptr.NumField(); i++ {
		g_elmt := gptr.Index(i)
		c_elmt := cptr.Index(i)
		err := cnv.elmt.cnv_from_c(g_elmt, c_elmt)
		if err != nil {
			return err
		}
	}
	return nil
	return nil
}

// EOF
