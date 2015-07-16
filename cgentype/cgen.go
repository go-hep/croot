// Package cgentype generates equivalent C++ types from Go types.
package cgentype

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-hep/croot/cmem"
)

func Generate(t reflect.Type) string {
	buf := []string{}
	ct := cmem.TypeOf(t)
	switch ct.Kind() {
	case cmem.Struct:
		buf = append(buf, "struct ", ct.Name(), " {\n")
		for i := 0; i < ct.NumField(); i++ {
			field := ct.Field(i)
			name := cxxtypename(field.Type)
			buf = append(
				buf,
				"\t",
				name.Name, " ", field.Name, name.Dims, "; // ",
				field.Type.GoType().Name(), " ",
				field.Type.Kind().String(),
				"\n",
			)
		}
		buf = append(buf, "};\n")

	case cmem.Slice:
		name := cxxtypename(ct)
		buf = append(buf, name.String(), "; // ", ct.GoType().Name(), "\n")

	case cmem.String:
		name := cxxtypename(ct)
		buf = append(buf, name.String(), ";\n")

	default:
		name := cxxtypename(ct)
		buf = append(buf, name.String(), "; // ", ct.GoType().Name(), " ", ct.Kind().String(), "\n")
	}

	return strings.Join(buf, "")
}

type typename struct {
	Name string
	Dims string
}

func (tn typename) String() string {
	return tn.Name + tn.Dims
}

func cxxtypename(ct cmem.Type) typename {

	switch ct.Kind() {
	case cmem.Slice:
		elem := ct.Elem()
		name := cxxtypename(elem)
		return typename{
			Name: "::golang::goslice< " + name.Name + name.Dims + " >",
			Dims: "",
		}

	case cmem.String:
		return typename{
			Name: "::golang::gostring",
			Dims: "",
		}

	case cmem.Ptr:
		if ct == cmem.C_string {
			return typename{
				Name: "::golang::gostring",
				Dims: "",
			}
		}

	case cmem.Array:
		elem := ct.Elem()
		name := cxxtypename(elem)
		n := ct.Len()
		return typename{
			Name: name.Name,
			Dims: fmt.Sprintf("[%d]%s", n, name.Dims),
		}

	case cmem.Int8:
		return typename{
			Name: "int8_t",
			Dims: "",
		}

	case cmem.Int16:
		return typename{
			Name: "int16_t",
			Dims: "",
		}

	case cmem.Int32:
		return typename{
			Name: "int32_t",
			Dims: "",
		}

	case cmem.Int64:
		return typename{
			Name: "int64_t",
			Dims: "",
		}

	case cmem.Uint8:
		return typename{
			Name: "uint8_t",
			Dims: "",
		}

	case cmem.Uint16:
		return typename{
			Name: "uint16_t",
			Dims: "",
		}

	case cmem.Uint32:
		return typename{
			Name: "uint32_t",
			Dims: "",
		}

	case cmem.Uint64:
		return typename{
			Name: "uint64_t",
			Dims: "",
		}

	}

	return typename{
		Name: ct.Name(),
		Dims: "",
	}
}
