package main

const code_tmpl = `// automatically generated!

package {{.Package}}

import (
  "fmt"

  "github.com/go-hep/croot"
)

{{range .Defs}}
type {{.Name}} struct {
{{range .Fields}}	{{.Name}} {{.Type}}
{{end}}}
{{end}}

{{with .DataReader}}
type DataReader struct {
{{range .Fields}}	{{.Name}} {{.Type}}
{{end}}

 // branches
{{range .Fields}}	b_{{.Name}} croot.Branch
{{end}}

	Tree croot.Tree
}

func NewDataReader(tree croot.Tree) (*DataReader, error) {
	dr := &DataReader{}
	err := dr.Init(tree)
	if err != nil {
		return nil, err
	}
	return dr, nil
}

func (dr *DataReader) Init(tree croot.Tree) error {
	var err error
	var o int32
	dr.Tree = tree
{{range .Fields}}
	o = dr.Tree.SetBranchAddress("{{.BranchName}}", &dr.{{.Name}})
	if o < 0 {
		return fmt.Errorf("invalid branch: [{{.BranchName}}] (got %d)", o)
	}
{{end}}
	return err
}

func (dr *DataReader) GetEntry(entry int64) int {
	if dr.Tree == nil {
		return 0
	}
	return dr.Tree.GetEntry(entry, 1)
}
{{end}}


func init() {
	// register all generated types with CRoot
{{range .Defs}}	croot.RegisterType(&{{.Name}}{})
{{end}}
}
`

// EOF
