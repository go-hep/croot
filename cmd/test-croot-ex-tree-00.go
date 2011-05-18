package main

import (
	"fmt"
	"flag"

	"bitbucket.org/binet/go-croot/pkg/croot"
)

type Det struct {
	e float64
	t float64
}

type Event struct {
	a Det
	b Det
}

var evtmax *int = flag.Int("evtmax", 10000, "number of events to generate")

func tree0(f *croot.File) {
	// create a tree
	tree := croot.NewTree("tree", "tree", 32)
	e := &Event{}
	
	bufsiz := 32000
	
	// create a branch with energy
	tree.Branch2("evt_a_e", &e.a.e, "evt_a_e/D", bufsiz)
	tree.Branch2("evt_a_t", &e.a.t, "evt_a_e/D", bufsiz)
	tree.Branch2("evt_b_e", &e.b.e, "evt_b_e/D", bufsiz)
	tree.Branch2("evt_b_t", &e.b.t, "evt_b_t/D", bufsiz)

	// fill some events with random numbers
	nevents := *evtmax
	for iev := 0; iev != nevents; iev++ {
		if iev%1000 == 0 {
			fmt.Printf(":: processing event %d...\n", iev)
		}

		// the two energies follow a gaussian distribution
		ea, eb := croot.GRandom.Rannord()
		e.a.e = ea
		e.b.e = eb

		e.a.t = croot.GRandom.Rndm(1)
		e.b.t = e.a.t * croot.GRandom.Gaus(0., 1.)

		if iev%1000 == 0 {
			fmt.Printf("evt.a.e= %8.3f\n", e.a.e)
			fmt.Printf("evt.a.t= %8.3f\n", e.a.t)
			fmt.Printf("evt.b.e= %8.3f\n", e.b.e)
			fmt.Printf("evt.b.t= %8.3f\n", e.b.t)
		}
		tree.Fill()
	}
	f.Write("",0,0)
}

func main() {
	flag.Parse()
	f := croot.OpenFile("event.root", "recreate", "my event file", 1, 0)
	tree0(f)
	f.Close("")
}
// EOF
