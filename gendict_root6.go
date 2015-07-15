// +build root6

package croot

import "reflect"

func to_cxx_name(t reflect.Type) string {
	//return fmt.Sprintf("::golang::%s::%s", t.PkgPath(), t.Name())
	return t.Name()
}

func gendict(t reflect.Type) {
	panic("gendict not implemented for ROOT-6")
}
