// +build root6

package croot

func initROOTInterpreter() {
	code := `
#include <stdint.h> // for intXXX_t
`
	err := genROOTDict(code)
	if err != nil {
		panic("croot: could not initialize CLing interpreter")
	}
}
