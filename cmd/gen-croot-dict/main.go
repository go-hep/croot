// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command gen-croot-dict generates a ROOT dictionary shared object file usable from ROOT and PyROOT.
//
// Example:
//  $> gen-croot-dict -o libgo-croot-dict.so dict.json
//  $> root -l
//  [root] gSystem->Load("./libgo-croot-dict.so");
//  [root] auto f = TFile::Open("event.root");
//  [root] auto t = (TTree*)f->Get("tree");
//  [root] auto evt = new Event;
//  [root] t->SetBranchAddress("evt", &evt);
//  [root] t->GetEntry(42);
//  [root] std::cout << "event: " << evt->I << "\n";
//
// where "dict.json" is:
//  [
//    {
//      "import": "github.com/go-hep/croot/testdata/edm",
//      "types": ["Event", "Det"]
//    }
//  ]
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var (
	fname = flag.String("o", "libgo-croot-dict.so", "path to output go-croot dictionary shared object")
)

func main() {
	flag.Parse()
	in, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	var pkgs []Package
	err = json.NewDecoder(in).Decode(&pkgs)
	if err != nil {
		log.Fatal(err)
	}

	massage(pkgs)

	data := struct {
		Packages []Package
	}{pkgs}

	tmpl := template.Must(template.New("gen-dict").Parse(srcTmpl))

	dir, err := ioutil.TempDir("", "go-croot-dict-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "main.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("go", "build", "-o", *fname, "-buildmode=c-shared", f.Name()) //"main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

type Package struct {
	Alias  string `json:"-"`
	Import string `json:"import"`
	Types  []Type `json:"types"`
}

func (pkg *Package) UnmarshalJSON(data []byte) error {
	var raw struct {
		Import string   `json:"import"`
		Types  []string `json:"types"`
	}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	pkg.Alias = ""
	pkg.Import = raw.Import
	pkg.Types = pkg.Types[:0]
	for _, typ := range raw.Types {
		pkg.Types = append(pkg.Types, Type{
			Package: "",
			Name:    typ,
		})
	}

	return nil
}

type Type struct {
	Package string `json:"-"`
	Name    string `json:"name"`
}

func massage(pkgs []Package) {
	for i, pkg := range pkgs {
		pkg.Alias = fmt.Sprintf("_pkg_%03d", i)
		for j := range pkg.Types {
			pkg.Types[j].Package = pkg.Alias
		}
		pkgs[i] = pkg
	}
}

const srcTmpl = `// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-hep/croot"

	{{- range .Packages}}
	{{.Alias}} "{{.Import}}"
	{{- end}}
)

func init() {
	{{- range .Packages}}
	{{- range .Types}}
	croot.RegisterType({{.Package}}.{{.Name}}{})
	{{- end}}
	{{- end}}
}

func main() {}
`
