// go-croot-version prints the ROOT version against which go-hep/croot has been
// compiled.
package main

import (
	"fmt"

	"go-hep.org/x/cgo/croot"
)

func main() {
	fmt.Printf("%v\n", croot.GRoot.GetVersion())
}
