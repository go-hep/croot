// +build root6

package croot

import "reflect"

func to_cxx_name(t reflect.Type) string {
	//return fmt.Sprintf("::golang::%s::%s", t.PkgPath(), t.Name())
	return t.Name()
}

func genreflex(t reflect.Type) {
	panic("genreflex not implemented for ROOT-6")
}
