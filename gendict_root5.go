// +build root5

package croot

func initROOTInterpreter() {
	// FIXME(sbinet): CINT chokes while parsing <stdint.h>
	/*
			code := `
		#include <stdint.h> // for intXXX_t
		`
			err := genROOTDict(code)
			if err != nil {
				panic("croot: could not initialize CINT interpreter")
			}
	*/
}
