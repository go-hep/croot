package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
//
import "C"

import (
	"unsafe"
)

type Class interface {
	Object
}

type class_impl struct {
	c C.CRoot_Class
}

func (c *class_impl) cptr() C.CRoot_Object {
	return (C.CRoot_Object)(c.c)
}

func (c *class_impl) as_tobject() *object_impl {
	return &object_impl{c.cptr()}
}

func (c *class_impl) ClassName() string {
	return c.as_tobject().ClassName()
}

func (c *class_impl) Clone(opt Option) Object {
	return c.as_tobject().Clone(opt)
}

func (c *class_impl) FindObject(name string) Object {
	return c.as_tobject().FindObject(name)
}

func (c *class_impl) GetName() string {
	return c.as_tobject().GetName()
}

func (c *class_impl) GetTitle() string {
	return c.as_tobject().GetTitle()
}

func (c *class_impl) InheritsFrom(clsname string) bool {
	return c.as_tobject().InheritsFrom(clsname)
}

func (c *class_impl) Print(option Option) {
	c.as_tobject().Print(option)
}

func GetClass(name string) Class {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c := C.CRoot_Class_GetClass(c_name)
	if c == nil {
		return nil
	}
	return &class_impl{c: c}
}

// EOF
