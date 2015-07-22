// go-croot-version prints the ROOT version against which go-hep/croot has been
// compiled.
package main

import (
	"fmt"

	"github.com/go-hep/croot"
)

func main() {
	fmt.Printf("%v\n", croot.GRoot.GetVersion())
}
