// +build ignore

package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("root-config", "--version")
	out, err := cmd.Output()
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

	case 6:
		err = copyFile("gendict_generated.go", "gendict_root6.go.tmpl")
		if err != nil {
			log.Fatalf("error generating dict: %v\n", err)
		}
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
