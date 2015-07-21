package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-hep/croot"
)

var (
	fname = flag.String("f", "", "ROOT file to inspect")
	tname = flag.String("t", "", "ROOT tree to inspect")
)

func main() {

	fmt.Printf(":: croot-ls ::\n")
	flag.Parse()

	if *fname == "" && len(os.Args) > 1 {
		*fname = os.Args[1]
	}

	if *fname == "" {
		fmt.Printf("**error** you have to give a (valid) path to a ROOT file\n")
		os.Exit(1)
	}

	if *tname == "" {
		fmt.Printf("**error** you have to give a TTree name to inspect off the ROOT file\n")
		os.Exit(1)
	}

	f, err := croot.OpenFile(*fname, "read", "go-croot-ls-file", 1, 0)
	if err != nil {
		fmt.Printf("**error** %v\n", err)
		os.Exit(1)
	}

	defer f.Close("")

	fmt.Printf(":: file: %s\n", f.GetName())
	tree := f.Get(*tname).(croot.Tree)
	if tree == nil {
		fmt.Printf("**error** no such tree [%s]\n", *tname)
		os.Exit(1)
	}
	fmt.Printf(":: tree: %s (entries=%v)\n", *tname, tree.GetEntries())
	branches := tree.GetListOfBranches()
	fmt.Printf(":: branches: %v\n", len(branches))
	for i, b := range branches {
		fmt.Printf("  br[%d]: %v\n", i, b.GetName())
		b.Print("")
	}

	os.Exit(0)
}
