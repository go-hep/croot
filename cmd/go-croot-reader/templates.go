package main

const code_tmpl = `// automatically generated!

package {{.Package}}

import (
  "fmt"

  "github.com/go-hep/croot"
)

{{range .Defs}}
type {{.Name}} struct {
{{range .Fields}} {{.Name}} {{.Type}}
{{end}}}
{{end}}

{{with .Event}}
type Event struct {
{{range .Fields}} {{.Name}} {{.Type}}
{{end}}

 // branches
{{range .Fields}} b_{{.Name}} croot.Branch
{{end}}

 Tree croot.Tree
}

func (e *Event) Init(tree croot.Tree) error {
 var err error
 var o int32
 e.Tree = tree
{{range .Fields}}
 o = e.Tree.SetBranchAddress("{{.BranchName}}", &e.{{.Name}})
 if o < 0 {
   return fmt.Errorf("invalid branch: [{{.BranchName}}] (got %d)", o)
 }
{{end}}
 return err
}

func (e *Event) GetEntry(entry int64) int {
 if e.Tree == nil {
   return 0
 }
 return e.Tree.GetEntry(entry, 1)
}
{{end}}


func init() {
  // register all generated types with CRoot
  {{range .Defs}}croot.RegisterType(&{{.Name}}{})
  {{end}}
}
`

// EOF
