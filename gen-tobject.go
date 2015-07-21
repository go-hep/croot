// +build ignore

package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	TObjects = []Type{
		{Name: "branchImpl", Var: "br"},
		{Name: "branchElementImpl", Var: "be"},
		{Name: "classImpl", Var: "c"},
		{Name: "dataMemberImpl", Var: "dm"},
		{Name: "dataTypeImpl", Var: "dt"},
		{Name: "fileImpl", Var: "f"},
		{Name: "h1FImpl", Var: "h"},
		{Name: "leafImpl", Var: "leaf"},
		{Name: "leafDImpl", Var: "leaf"},
		{Name: "leafFImpl", Var: "leaf"},
		{Name: "leafIImpl", Var: "leaf"},
		{Name: "leafOImpl", Var: "leaf"},
		{Name: "objArrayImpl", Var: "obj"},
		{Name: "randomImpl", Var: "rndm"},
		{
			Name: "treeImpl",
			Var:  "tree",
			PrintImpl: `func (t *treeImpl) Print(option Option) {
	coption := C.CString(string(option))
	defer C.free(unsafe.Pointer(coption))

	C.CRoot_Tree_Print(t.c, (*C.CRoot_Option)(coption))
}`,
		},
	}
)

type Type struct {
	Name string
	Var  string

	PrintImpl string
}

func (t Type) ROOTType() string {
	return "T" + strings.Title(t.Name)[:len(t.Name)-len("impl")]
}

func (t Type) CRootType() string {
	return "CRoot_" + strings.Title(t.Name)[:len(t.Name)-len("impl")]
}

func main() {
	f, err := os.Create("object_impl.go")
	if err != nil {
		log.Fatalf("error creating [object_impl.go]: %v\n", err)
	}
	defer f.Close()

	tmpl, err := template.New("tobjects").Parse(tmplTObject)
	if err != nil {
		log.Fatalf("error parsing template: %v\n", err)
	}

	err = tmpl.Execute(f, TObjects)
	if err != nil {
		log.Fatalf("error generating TObject implementation: %v\n",
			err,
		)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("error closing [%s]: %v\n", f.Name(), err)
	}

	cmd := exec.Command("gofmt", "-w", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error gofmt-ing [%s]: %v\n", f.Name(), err)
	}
}

const tmplTObject = `// automatically generated.
// DO NOT EDIT

package croot

// #include "croot/croot.h"
//
// #include <stdlib.h>
// #include <string.h>
//
import "C"

import (
	"unsafe"
)

{{range .}}
//
// --- TObject impl for {{.Name}} ({{.ROOTType}}) ---
//

func ({{.Var}} *{{.Name}}) cptr() C.CRoot_Object {
	return (C.CRoot_Object)({{.Var}}.c)
}

func ({{.Var}} *{{.Name}}) as_tobject() *objectImpl {
	return &objectImpl{ {{.Var}}.cptr() }
}

func ({{.Var}} *{{.Name}}) ClassName() string {
	return {{.Var}}.as_tobject().ClassName()
}

func ({{.Var}} *{{.Name}}) Clone(opt Option) Object {
	return {{.Var}}.as_tobject().Clone(opt)
}

func ({{.Var}} *{{.Name}}) FindObject(name string) Object {
	return {{.Var}}.as_tobject().FindObject(name)
}

func ({{.Var}} *{{.Name}}) GetName() string {
	return {{.Var}}.as_tobject().GetName()
}

func ({{.Var}} *{{.Name}}) GetTitle() string {
	return {{.Var}}.as_tobject().GetTitle()
}

func ({{.Var}} *{{.Name}}) InheritsFrom(clsname string) bool {
	return {{.Var}}.as_tobject().InheritsFrom(clsname)
}

{{if .PrintImpl}}{{.PrintImpl}}{{else}}func ({{.Var}} *{{.Name}}) Print(option Option) {
	{{.Var}}.as_tobject().Print(option)
}
{{end}}

{{end}}

func init() {
	// register conversions
	{{range .}}
	cnvmap["{{.ROOTType}}"] = func(o c_object) Object {
		return &{{.Name}}{ c: (C.{{.CRootType}})(o.cptr())}
	}
	{{end}}
}
`
