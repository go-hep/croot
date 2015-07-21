package croot_test

import (
	"testing"

	"github.com/go-hep/croot"
)

func TestClass(t *testing.T) {
	cls := croot.GetClass("golang::gostring")
	if cls == nil {
		t.Fatalf("could not retrieve croot.Class for 'golang::gostring'\n")
	}

	cls.Print("")

	for _, table := range []struct {
		Name         string
		Offset       int
		TypeName     string
		FullTypeName string
		ArrayDim     int
	}{
		{
			Name:         "Len",
			Offset:       0,
			TypeName:     "int",
			FullTypeName: "int",
			ArrayDim:     0,
		},
		{
			Name:         "Data",
			Offset:       8,
			TypeName:     "char",
			FullTypeName: "char*",
			ArrayDim:     0,
		},
	} {

		dm := cls.GetDataMember(table.Name)
		if dm == nil {
			t.Fatalf(
				"could not retrieve %q data member from 'golang::gostring'\n",
				table.Name,
			)
		}

		offset := dm.GetOffset()
		if table.Offset != offset {
			t.Fatalf(
				"golang::gostring.%s: offset error. want=%d. got=%d\n",
				table.Name, table.Offset,
				offset,
			)
		}

		tname := dm.GetTypeName()
		if table.TypeName != tname {
			t.Fatalf(
				"golang::gostring.%s: typename error. want=%q. got=%q\n",
				table.Name, table.TypeName,
				tname,
			)
		}

		ftname := dm.GetFullTypeName()
		if table.FullTypeName != ftname {
			t.Fatalf(
				"golang::gostring.%s: full-typename error. want=%q. got=%q\n",
				table.Name, table.FullTypeName,
				ftname,
			)
		}

		ndims := dm.GetArrayDim()
		if table.ArrayDim != ndims {
			t.Fatalf(
				"golang::gostring.%s: ndims error. want=%d. got=%d\n",
				table.Name, table.TypeName,
				ndims,
			)
		}
	}

}
