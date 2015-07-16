package croot

import "reflect"

// RegisterType declares the (equivalent) C-layout of value v to ROOT so
// values of the same type than v can be written out to ROOT files
func RegisterType(v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	t := rv.Type()
	//fmt.Printf("registering [%s] (sz:%d)...\n",t, t.Size())
	gendict(t)
}

func follow_ptr(v reflect.Value) reflect.Value {
	for {
		switch v.Kind() {
		case reflect.Ptr:
			if v.Elem().Kind() == reflect.Ptr {
				v = v.Elem()
			} else {
				return v
			}
		default:
			return v
		}
	}
}

func to_cxx_name(t reflect.Type) string {
	//return fmt.Sprintf("::golang::%s::%s", t.PkgPath(), t.Name())
	return t.Name()
}
