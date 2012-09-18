package croot_test

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"

	"github.com/sbinet/go-croot/pkg/croot"
	"github.com/sbinet/go-ffi/pkg/ffi"
)

type Det struct {
	E float64
	T float64
	Ts []float64 //FIXME: not yet...
}

type Event struct {
	I int64
	A Det
	B Det
}

func TestTreeBuiltinsRW(t *testing.T) {
	const fname = "simple-event.root"
	const evtmax = 10000
	const splitlevel = 32
	const bufsiz = 32000
	const compress = 1
	const netopt = 0

	// write
	ref := make([]string, 0, 50)
	{
		add := func(str string) {
			ref = append(ref, str)
		}

		f := croot.OpenFile(fname, "recreate", "croot event file", compress, netopt)

		// create a tree
		tree := croot.NewTree("tree", "tree", splitlevel)

		e := Event{}

		// create a branch with energy
		tree.Branch2("evt_i", &e.I, "evt_i/L", bufsiz)
		tree.Branch2("evt_a_e", &e.A.E, "evt_a_e/D", bufsiz)
		tree.Branch2("evt_a_t", &e.A.T, "evt_a_t/D", bufsiz)
		tree.Branch2("evt_b_e", &e.B.E, "evt_b_e/D", bufsiz)
		tree.Branch2("evt_b_t", &e.B.T, "evt_b_t/D", bufsiz)

		// initialize our source of random numbers...
		src := rand.New(rand.NewSource(1))

		// fill some events with random numbers
		for iev := int64(0); iev != evtmax; iev++ {
			if iev%1000 == 0 {
				add(fmt.Sprintf(":: processing event %d...\n", iev))
			}

			e.I = iev
			// the two energies follow a gaussian distribution
			e.A.E = src.NormFloat64()
			e.B.E = src.NormFloat64()

			e.A.T = src.Float64()
			e.B.T = e.A.T * (src.NormFloat64()*1. + 0.)

			if iev%1000 == 0 {
				add(fmt.Sprintf("evt.i=   %8d\n", e.I))
				add(fmt.Sprintf("evt.a.e= %8.3f\n", e.A.E))
				add(fmt.Sprintf("evt.a.t= %8.3f\n", e.A.T))
				add(fmt.Sprintf("evt.b.e= %8.3f\n", e.B.E))
				add(fmt.Sprintf("evt.b.t= %8.3f\n", e.B.T))
			}
			tree.Fill()
		}
		f.Write("", 0, 0)
		f.Close("")
	}

	// read back
	chk := make([]string, 0, 50)
	{
		add := func(str string) {
			chk = append(chk, str)
		}

		f := croot.OpenFile(fname, "read", "croot event file", compress, netopt)
		tree := f.GetTree("tree")
		if tree.GetEntries() != evtmax {
			t.Errorf("expected [%v] entries, got %v\n", evtmax, tree.GetEntries())
		}

		e := Event{}

		tree.SetBranchAddress("evt_i", &e.I)
		tree.SetBranchAddress("evt_a_e", &e.A.E)
		tree.SetBranchAddress("evt_a_t", &e.A.T)
		tree.SetBranchAddress("evt_b_e", &e.B.E)
		tree.SetBranchAddress("evt_b_t", &e.B.T)

		// read events
		for iev := int64(0); iev != evtmax; iev++ {
			if iev%1000 == 0 {
				add(fmt.Sprintf(":: processing event %d...\n", iev))
			}
			if tree.GetEntry(iev, 1) <= 0 {
				panic("error")
			}
			if iev%1000 == 0 {
				add(fmt.Sprintf("evt.i=   %8d\n", e.I))
				add(fmt.Sprintf("evt.a.e= %8.3f\n", e.A.E))
				add(fmt.Sprintf("evt.a.t= %8.3f\n", e.A.T))
				add(fmt.Sprintf("evt.b.e= %8.3f\n", e.B.E))
				add(fmt.Sprintf("evt.b.t= %8.3f\n", e.B.T))
			}

			if iev != e.I {
				t.Errorf("invalid event number. expected %v, got %v", iev, e.I)
			}
		}
		f.Close("")
	}

	if !reflect.DeepEqual(ref, chk) {
		t.Errorf("log files do not match")
	}

	err := os.Remove(fname)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestTreeStructRW(t *testing.T) {
	const fname = "struct-event.root"
	const evtmax = 10000
	const splitlevel = 32
	const bufsiz = 32000
	const compress = 1
	const netopt = 0

	// write
	ref := make([]string, 0, 50)
	{
		add := func(str string) {
			ref = append(ref, str)
		}

		f := croot.OpenFile(fname, "recreate", "croot event file", compress, netopt)

		// create a tree
		tree := croot.NewTree("tree", "tree", splitlevel)

		e := Event{}

		tree.Branch("evt", &e, bufsiz, 0)

		// initialize our source of random numbers...
		src := rand.New(rand.NewSource(1))

		// fill some events with random numbers
		for iev := int64(0); iev != evtmax; iev++ {
			if iev%1000 == 0 {
				add(fmt.Sprintf(":: processing event %d...\n", iev))
			}

			e.I = iev
			// the two energies follow a gaussian distribution
			e.A.E = src.NormFloat64()
			e.B.E = src.NormFloat64()

			e.A.T = src.Float64()
			e.B.T = e.A.T * (src.NormFloat64()*1. + 0.)

			if iev%1000 == 0 {
				add(fmt.Sprintf("evt.i=   %8d\n", e.I))
				add(fmt.Sprintf("evt.a.e= %8.3f\n", e.A.E))
				add(fmt.Sprintf("evt.a.t= %8.3f\n", e.A.T))
				add(fmt.Sprintf("evt.b.e= %8.3f\n", e.B.E))
				add(fmt.Sprintf("evt.b.t= %8.3f\n", e.B.T))
			}
			tree.Fill()
		}
		f.Write("", 0, 0)
		f.Close("")
	}

	// read back
	chk := make([]string, 0, 50)
	{
		add := func(str string) {
			chk = append(chk, str)
		}

		f := croot.OpenFile(fname, "read", "croot event file", compress, netopt)
		tree := f.GetTree("tree")
		if tree.GetEntries() != evtmax {
			t.Errorf("expected [%v] entries, got %v\n", evtmax, tree.GetEntries())
		}

		e := Event{}
		tree.SetBranchAddress("evt", &e)

		// read events
		for iev := int64(0); iev != evtmax; iev++ {
			if iev%1000 == 0 {
				add(fmt.Sprintf(":: processing event %d...\n", iev))
			}
			if tree.GetEntry(iev, 1) <= 0 {
				panic("error")
			}
			if iev%1000 == 0 {
				add(fmt.Sprintf("evt.i=   %8d\n", e.I))
				add(fmt.Sprintf("evt.a.e= %8.3f\n", e.A.E))
				add(fmt.Sprintf("evt.a.t= %8.3f\n", e.A.T))
				add(fmt.Sprintf("evt.b.e= %8.3f\n", e.B.E))
				add(fmt.Sprintf("evt.b.t= %8.3f\n", e.B.T))
			}

			if iev != e.I {
				t.Errorf("invalid event number. expected %v, got %v", iev, e.I)
			}
		}
		f.Close("")
	}

	if !reflect.DeepEqual(ref, chk) {
		t.Errorf("log files do not match")
	}

	err := os.Remove(fname)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestTreeBuiltinsRWViaFFI(t *testing.T) {
	const fname = "simple-event-via-ffi.root"
	const evtmax = 10000
	const splitlevel = 32
	const bufsiz = 32000
	const compress = 1
	const netopt = 0

	ctype, err := ffi.NewStructType(
		"cffi_event",
		[]ffi.Field{
			{"evt_i", ffi.C_int64},
			{"evt_a_e", ffi.C_double},
			{"evt_a_t", ffi.C_double},
			{"evt_b_e", ffi.C_double},
			{"evt_b_t", ffi.C_double},
		})
	if err != nil {
		t.Errorf("could not create ffi.Event struct type: %v\n", err)
	}

	// write
	ref := make([]string, 0, 50)
	{
		add := func(str string) {
			ref = append(ref, str)
		}

		f := croot.OpenFile(fname, "recreate", "croot event file", compress, netopt)

		// create a tree
		tree := croot.NewTree("tree", "tree", splitlevel)

		e := Event{}

		// create a branch with energy
		tree.Branch2("evt_i", &e.I, "evt_i/L", bufsiz)
		tree.Branch2("evt_a_e", &e.A.E, "evt_a_e/D", bufsiz)
		tree.Branch2("evt_a_t", &e.A.T, "evt_a_t/D", bufsiz)
		tree.Branch2("evt_b_e", &e.B.E, "evt_b_e/D", bufsiz)
		tree.Branch2("evt_b_t", &e.B.T, "evt_b_t/D", bufsiz)

		// initialize our source of random numbers...
		src := rand.New(rand.NewSource(1))

		// fill some events with random numbers
		for iev := int64(0); iev != evtmax; iev++ {
			if iev%1000 == 0 {
				add(fmt.Sprintf(":: processing event %d...\n", iev))
			}

			e.I = iev
			// the two energies follow a gaussian distribution
			e.A.E = src.NormFloat64()
			e.B.E = src.NormFloat64()

			e.A.T = src.Float64()
			e.B.T = e.A.T * (src.NormFloat64()*1. + 0.)

			if iev%1000 == 0 {
				add(fmt.Sprintf("evt.i=   %8d\n", e.I))
				add(fmt.Sprintf("evt.a.e= %8.3f\n", e.A.E))
				add(fmt.Sprintf("evt.a.t= %8.3f\n", e.A.T))
				add(fmt.Sprintf("evt.b.e= %8.3f\n", e.B.E))
				add(fmt.Sprintf("evt.b.t= %8.3f\n", e.B.T))
			}
			tree.Fill()
		}
		f.Write("", 0, 0)
		f.Close("")
	}

	// read back
	chk := make([]string, 0, 50)
	{
		add := func(str string) {
			chk = append(chk, str)
		}

		f := croot.OpenFile(fname, "read", "croot event file", compress, netopt)
		tree := f.GetTree("tree")
		if tree.GetEntries() != evtmax {
			t.Errorf("expected [%v] entries, got %v\n", evtmax, tree.GetEntries())
		}

		evt_i := ffi.New(ctype.Field(0).Type)
		evt_a_e := ffi.New(ctype.Field(1).Type)
		evt_a_t := ffi.New(ctype.Field(2).Type)
		evt_b_e := ffi.New(ctype.Field(3).Type)
		evt_b_t := ffi.New(ctype.Field(4).Type)

		tree.SetBranchAddress("evt_i", evt_i)
		tree.SetBranchAddress("evt_a_e", evt_a_e)
		tree.SetBranchAddress("evt_a_t", evt_a_t)
		tree.SetBranchAddress("evt_b_e", evt_b_e)
		tree.SetBranchAddress("evt_b_t", evt_b_t)

		// read events
		for iev := int64(0); iev != evtmax; iev++ {
			if iev%1000 == 0 {
				add(fmt.Sprintf(":: processing event %d...\n", iev))
			}
			if tree.GetEntry(iev, 1) <= 0 {
				panic("error")
			}
			if iev%1000 == 0 {
				add(fmt.Sprintf("evt.i=   %8d\n", evt_i.Int()))
				add(fmt.Sprintf("evt.a.e= %8.3f\n", evt_a_e.Float()))
				add(fmt.Sprintf("evt.a.t= %8.3f\n", evt_a_t.Float()))
				add(fmt.Sprintf("evt.b.e= %8.3f\n", evt_b_e.Float()))
				add(fmt.Sprintf("evt.b.t= %8.3f\n", evt_b_t.Float()))
			}

			if iev != evt_i.Int() {
				t.Errorf("invalid event number. expected %v, got %v", iev, evt_i.Int())
			}
		}
		f.Close("")
	}

	if !reflect.DeepEqual(ref, chk) {
		t.Errorf("log files do not match")
	}

	err = os.Remove(fname)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func init() {
	croot.RegisterType(&Event{})
}

// EOF
