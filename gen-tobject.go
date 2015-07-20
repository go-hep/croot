// +build ignore

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

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

func ({{.Var}} *{{.Name}}) cptr() {{.Var}}.CRoot_Object {
	return ({{.Var}}.CRoot_Object)({{.Var}}.c)
}

func ({{.Var}} *{{.Name}}) as_tobject() *object_impl {
	return &object_impl{ {{.Var}}.cptr() }
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

func ({{.Var}} *{{.Name}}) Print(option Option) {
	{{.Var}}.as_tobject().Print(option)
}

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

var (
	TObjects = []Type{
		{Name: "branchImpl", Var: "br"},
		{Name: "branchElementImpl", Var: "be"},
		{Name: "classImpl", Var: "c"},
		{Name: "dataMemberImpl", Var: "mbr"},
		{Name: "fileImpl", Var: "f"},
		{Name: "h1fImpl", Var: "h"},
		{Name: "leafImpl", Var: "leaf"},
		{Name: "leafDImpl", Var: "leaf"},
		{Name: "leafFImpl", Var: "leaf"},
		{Name: "leafIImpl", Var: "leaf"},
		{Name: "leafOImpl", Var: "leaf"},
		{Name: "objArrayImpl", Var: "obj"},
		{Name: "randomImpl", Var: "rndm"},
		{Name: "treeImpl", Var: "tree"},
	}
)

type Type struct {
	Name string
	Var  string
}

func (t Type) ROOTType() string {
	return "T" + strings.Title(t.Name)[:len(t.Name)-len("impl")]
}

func (t Type) CRootType() string {
	return "CRoot_" + strings.Title(t.Name)[:len(t.Name)-len("impl")]
}

func main() {
	tmpl, err := template.New("tobjects").Parse(tmplTObject)
	if err != nil {
		log.Fatalf("error parsing template: %v\n", err)
	}

	err = tmpl.Execute(os.Stdout, TObjects)
	if err != nil {
		log.Fatalf("error generating TObject implementation: %v\n",
			err,
		)
	}
}
