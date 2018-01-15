package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"go-hep.org/x/cgo/croot"
)

var (
	fname = flag.String("f", "", "ROOT file to inspect")
	tname = flag.String("t", "", "ROOT tree to inspect")
	oname = flag.String("o", "event.go", "path to file to fill with Go code")
	pname = flag.String("p", "main", "name of the Go package to generate")
)

type StructDef struct {
	Name   string
	Fields []FieldDef
}

type FieldDef struct {
	Name       string
	Type       string
	VarName    string
	BranchName string
}

type Context struct {
	Package    string
	DataReader *StructDef
	Defs       map[string]*StructDef
}

var to_go_name = strings.Title

func gen_code(w io.Writer, ctx Context) error {
	var err error
	t := template.New("top")
	template.Must(t.Parse(code_tmpl))
	err = t.Execute(w, ctx)
	return err
}

func main() {

	fmt.Printf(":: croot-reader ::\n")
	flag.Parse()

	if *fname == "" && len(os.Args) > 1 {
		*fname = os.Args[1]
	}

	if *fname == "" {
		fmt.Printf("**error** you have to give a (valid) path to a ROOT file\n")
		os.Exit(1)
	}

	if *tname == "" && len(os.Args) > 2 {
		*tname = os.Args[2]
	}

	if *tname == "" {
		fmt.Printf("**error** you have to give a TTree name to inspect off the ROOT file\n")
		os.Exit(1)
	}

	if *oname == "" {
		fmt.Printf("**error** you have to give a path to an output file (which will hold the generated Go code.)\n")
		os.Exit(1)
	}

	if *pname == "" {
		fmt.Printf("**error** you have to give a (Go) package name to generate\n")
		os.Exit(1)
	}

	o, err := os.Create(*oname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "**error** %v\n", err)
		os.Exit(1)
	}
	defer func(o *os.File) {
		o.Sync()
		o.Close()
	}(o)

	ctx := Context{
		Package: *pname,
		Defs: map[string]*StructDef{
			"DataReader": {
				Name:   "DataReader",
				Fields: nil,
			},
		},
	}

	f, err := croot.OpenFile(*fname, "read", "go-croot-gen-datareader", 1, 0)
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

	defs := ctx.Defs

	fmt.Printf(":: tree: %s (entries=%v)\n", *tname, tree.GetEntries())
	branches := tree.GetListOfBranches()
	fmt.Printf(":: branches: %v\n", len(branches))
	for i, branch := range branches {
		n := branch.GetName()
		go_name := to_go_name(n)
		fmt.Printf(":: branch[%3d]=%s (=> %s)\n", i, n, go_name)
		leaves := branch.GetListOfLeaves()
		fmt.Printf(":: leaves: %v\n", len(leaves))
		br_struct := StructDef{Name: go_name, Fields: nil}
		for j, leaf := range leaves {
			fmt.Printf("  [%03d] leaf: %v\n", j, n)
			//leaf.Print("")
			nn := to_go_name(leaf.GetName())
			typename := leaf.GetTypeName()
			gotype, ok := go_typemap[typename]
			if !ok {
				gotype = "Undefined"
			}
			br_field := FieldDef{
				Name:       nn,
				BranchName: leaf.GetName(),
				VarName:    nn,
				Type:       gotype,
			}
			if j == 0 {
				br_field.VarName = n + "." + nn
			}
			br_struct.Fields = append(br_struct.Fields, br_field)

		}
		if len(br_struct.Fields) > 1 {
			defs[n] = &br_struct
			defs["DataReader"].Fields = append(
				defs["DataReader"].Fields,
				FieldDef{
					Name:       go_name,
					BranchName: n,
					VarName:    br_struct.Fields[0].VarName,
					Type:       go_name,
				},
			)
		} else {
			// lump into DataReader.
			defs["DataReader"].Fields = append(
				defs["DataReader"].Fields,
				br_struct.Fields...,
			)
		}
	}

	ctx.DataReader = defs["DataReader"]
	delete(defs, "DataReader")

	err = gen_code(o, ctx)
	os.Exit(0)
}
