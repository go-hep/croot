croot
=====

[![GoDoc](https://godoc.org/github.com/go-hep/croot?status.svg)](https://godoc.org/github.com/go-hep/croot)

Go bindings to the C-API of ROOT (CRoot).

`croot` supports now both `ROOT-5` and `ROOT-6`.

## Installation

Unfortunately, ``go-hep/croot`` isn't completely ``go-get`` able.
You **NEED** to run `go generate` (or `make gen`) once the code has been retrieved from github:

```sh
$ go get -d github.com/go-hep/croot
$ cd $GOPATH/src/github.com/go-hep/croot
$ make clean gen install
$ make test
ok  	github.com/go-hep/croot	2.862s
```    


## Documentation

 http://godoc.org/github.com/go-hep/croot


## ``TTree`` Limitations/Status

- read/write of C-builtins: **WORKS**
- read/write of arrays: **WORKS**
- read/write of slices: **WORKS**
- read/write of structs: **WORKS**
- read/write of strings: **WORKS**

## Example

`croot` can now (correctly) write and read `go` structs which have
an equivalent `C` representation.
**Except** for structs embedding pointers.

A typical write program would look like:

``` go
package main

import (
	"fmt"
	"flag"
	"math/rand"

	"github.com/go-hep/croot"
)

type Det struct {
	E float64
	T float64
}

type Event struct {
	A Det
	B Det
}

var evtmax *int = flag.Int("evtmax", 10000, "number of events to generate")
var fname *string = flag.String("fname", "event.root", "file to create")

func main() {
	flag.Parse()
	f, err := croot.OpenFile(*fname, "recreate", "my event file", 1, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close("")
	tree0(f)
}

func tree0(f croot.File) {
	// create a tree
	tree := croot.NewTree("tree", "tree", 32)
	e := Event{}
	const bufsiz = 32000

	_, err := tree.Branch("evt", &e, bufsiz, 0)
	if err != nil {
	   panic(err)
	}

	// fill some events with random numbers
	nevents := *evtmax
	for iev := 0; iev != nevents; iev++ {
		if iev%1000 == 0 {
			fmt.Printf(":: processing event %d...\n", iev)
		}


		// the two energies follow a gaussian distribution
		e.A.E = rand.NormFloat64() //ea
		e.B.E = rand.NormFloat64() //eb

		e.A.T = croot.GRandom.Rndm(1)
		e.B.T = e.A.T * croot.GRandom.Gaus(0., 1.)
		if iev%1000 == 0 {
			fmt.Printf("ievt: %d\n", iev)
			fmt.Printf("evt.a.e= %8.3f\n", e.A.E)
			fmt.Printf("evt.a.t= %8.3f\n", e.A.T)
			fmt.Printf("evt.b.e= %8.3f\n", e.B.E)
			fmt.Printf("evt.b.t= %8.3f\n", e.B.T)
		}
		_, err = tree.Fill()
		if err != nil {
		   panic(err)
        }
	}
	f.Write("", 0, 0)
}
// EOF
```

and a typical reading program:
``` go
package main

import (
	"fmt"
	"flag"

	"github.com/go-hep/croot"
)

type Det struct {
	E float64
	T float64
}

type Event struct {
	A Det
	B Det
}

var evtmax *int64 = flag.Int64("evtmax", 10000, "number of events to read")
var fname *string = flag.String("fname", "event.root", "file to read back")

func main() {
	flag.Parse()
	
	croot.RegisterType(&Event{})

	fmt.Printf(":: opening [%s]...\n", *fname)
	f, err := croot.OpenFile(*fname, "read", "my event file", 1, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close("")
	tree0(f)

}

func tree0(f croot.File) {
	t := f.Get("tree").(croot.Tree)
	e := Event{}

	t.SetBranchAddress("evt", &e)

	// fill some events with random numbers
	nevents := int64(*evtmax)
	if nevents < 0 || nevents > int64(t.GetEntries()) {
		nevents = int64(t.GetEntries())
	}
	for iev := int64(0); iev != nevents; iev++ {
		if iev%1000 == 0 {
			fmt.Printf(":: processing event %d...\n", iev)
		}
		if t.GetEntry(iev, 1) <= 0 {
			panic("error")
		}
		if iev%1000 == 0 {
			fmt.Printf("ievt: %d\n", iev)
			fmt.Printf("evt.a.e= %8.3f\n", e.A.E)
			fmt.Printf("evt.a.t= %8.3f\n", e.A.T)
			fmt.Printf("evt.b.e= %8.3f\n", e.B.E)
			fmt.Printf("evt.b.t= %8.3f\n", e.B.T)
		}
	}
}

// EOF
```
