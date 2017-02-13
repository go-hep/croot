// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	out, err := exec.Command("root-config", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}
	vers := 0
	switch {
	case bytes.HasPrefix(out, []byte("5")):
		vers = 5
	case bytes.HasPrefix(out, []byte("6")):
		vers = 6
	default:
		log.Fatalf("invalid ROOT version: %q\n", string(vers))
	}

	cflags, err := exec.Command("root-config", "--cflags").Output()
	if err != nil {
		log.Fatalf("error getting ROOT cflags: %v\n", err)
	}
	cflags = trim(cflags)

	ldflags, err := exec.Command("root-config", "--ldflags", "--libs").Output()
	if err != nil {
		log.Fatalf("error getting ROOT ldflags: %v\n", err)
	}
	ldflags = trim(ldflags)

	cxxSources := []string{
		"bindings/src/croot.cxx",
		"bindings/src/croot_class.cxx",
		"bindings/src/croot_go_schema.cxx",
		"bindings/src/croot_goobject.cxx",
		"bindings/src/croot_hist.cxx",
		"bindings/src/croot_interpreter.cxx",
		"bindings/src/croot_leaf.cxx",
		// dict
		"bindings/src/goedm_dict.cxx",
	}

	cxxObjects := make([]string, len(cxxSources))
	for i, src := range cxxSources {
		obj := strings.Replace(src, ".cxx", ".o", -1)
		cxxObjects[i] = obj
	}

	var (
		cxxCRootCXXFlags []string
		cxxCRootLDFlags  []string
	)

	genDictFile := fmt.Sprintf("gen-goedm-dict-root%d.go", vers)

	switch vers {
	case 5:
		err = copyFile("gendict_generated.go", "gendict_root6.go.tmpl")
		if err != nil {
			log.Fatalf("error generating dict: %v\n", err)
		}

		err = copyFile("reflex_generated.go", "reflex_root5.go.tmpl")
		if err != nil {
			log.Fatalf("error generating reflex: %v\n", err)
		}

		err = copyFile("cintex_generated.go", "cintex_root5.go.tmpl")
		if err != nil {
			log.Fatalf("error generating cintex: %v\n", err)
		}
		ldflags = append(ldflags, []byte(" -lReflex -lCintex")...)

	case 6:
		err = copyFile("gendict_generated.go", "gendict_root6.go.tmpl")
		if err != nil {
			log.Fatalf("error generating dict: %v\n", err)
		}
	}

	ctx := build.Default
	gopath := strings.Split(ctx.GOPATH, string(os.PathListSeparator))[0]
	libdir := filepath.Join(gopath, "pkg", ctx.GOOS+"_"+ctx.GOARCH, "github.com", "go-hep", "croot", "_lib")

	var (
		cgoLDFlags string
		cgoCFlags  string
	)

	switch runtime.GOOS {
	case "linux":
		cgoLDFlags = fmt.Sprintf("-Wl,-rpath,%s -L%s -lcxx-croot %s",
			libdir, libdir, string(ldflags),
		)
		cgoCFlags = fmt.Sprintf("-Ibindings/inc -I. %s", string(cflags))
	default:
		cgoLDFlags = fmt.Sprintf("-L%s -lcxx-croot %s", libdir, string(ldflags))
		cgoCFlags = fmt.Sprintf("-Ibindings/inc -I. %s", string(cflags))
	}

	err = ioutil.WriteFile(
		"cgoflags_generated.go",
		[]byte(fmt.Sprintf(
			`// Automatically generated.
// Do NOT edit.

package croot

// #include "croot/croot.h"
// #cgo CFLAGS: %[1]s
// #cgo CXXFLAGS: %[1]s
// #cgo LDFLAGS: %[2]s
import "C"
`,
			"-Ibindings/inc -I.",
			cgoLDFlags,
		)),
		0644,
	)
	if err != nil {
		log.Fatalf("error generating cgo flags: %v\n", err)
	}

	err = os.MkdirAll(libdir, 0755)
	if err != nil {
		log.Fatalf("error creating libcxx-croot install directory: %v\n", err)
	}

	crootDir := filepath.Join(gopath, "src", "github.com", "go-hep", "croot")
	// build C++ -> C ROOT shim library
	// 1. generate the dict code
	cmd := exec.Command("go", "generate", genDictFile)
	cmd.Dir = filepath.Join(crootDir, "bindings", "src")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error generating goedm dict file: %v\n", err)
	}

	cxx := os.Getenv("CXX")
	if cxx == "" {
		cxx = "c++"
	}

	os.Remove(filepath.Join(libdir, "libcxx-croot.so"))

	cxxCRootCXXFlags = append(cxxCRootCXXFlags, strings.Split(cgoCFlags, " ")...)
	cxxCRootCXXFlags = append(cxxCRootCXXFlags, "-fPIC")

	cxxCRootLDFlags = append(cxxCRootLDFlags, strings.Split(string(ldflags), " ")...)

	// 2. actually build+install lib-croot.so
	for i := range cxxSources {
		obj := cxxObjects[i]
		src := cxxSources[i]
		os.Remove(obj)
		var args []string
		args = append(args, cxxCRootCXXFlags...)
		args = append(args, "-o", obj, "-c", src)
		cmd = exec.Command(cxx, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("error compiling [%s]: %v\n", src, err)
		}
	}

	cmd = exec.Command(cxx, "-shared",
		"-o", filepath.Join(libdir, "libcxx-croot.so"),
	)
	cmd.Args = append(cmd.Args, cxxCRootCXXFlags...)
	cmd.Args = append(cmd.Args, cxxCRootLDFlags...)
	cmd.Args = append(cmd.Args, cxxObjects...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error build libcxx-croot: %v\n", err)
	}

	// 3. install the PCM dict file.
	err = copyFile(
		filepath.Join(libdir, "goedm_dict_rdict.pcm"),
		filepath.Join(crootDir, "bindings", "src", "goedm_dict_rdict.pcm"),
	)
	if err != nil {
		log.Fatalf("error installing PCM dict: %v\n", err)
	}

}

func copyFile(dst, src string) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()

	dstf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstf.Close()

	_, err = io.Copy(dstf, srcf)
	if err != nil {
		return err
	}

	return dstf.Close()
}

func trim(b []byte) []byte {
	return bytes.TrimRight(b, "\r\n")
}
