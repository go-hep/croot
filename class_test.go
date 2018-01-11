package croot_test

import (
	"testing"

	"github.com/go-hep/croot"
)

func TestClass(t *testing.T) {
	type MemberDescr struct {
		Name         string
		Offset       int
		TypeName     string
		FullTypeName string
		ArrayDim     int
	}

	for _, table := range []struct {
		Name    string
		Members []MemberDescr
	}{
		{
			Name: "golang::gostring",
			Members: []MemberDescr{
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
			},
		},
		{
			Name: "golang::goslice<double>",
			Members: []MemberDescr{
				{
					Name:         "Len",
					Offset:       0,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Cap",
					Offset:       4,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Data",
					Offset:       8,
					TypeName:     "double",
					FullTypeName: "double*",
					ArrayDim:     0,
				},
			},
		},
		{
			Name: "golang::goslice<float>",
			Members: []MemberDescr{
				{
					Name:         "Len",
					Offset:       0,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Cap",
					Offset:       4,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Data",
					Offset:       8,
					TypeName:     "float",
					FullTypeName: "float*",
					ArrayDim:     0,
				},
			},
		},
		{
			Name: "golang::goslice<int>",
			Members: []MemberDescr{
				{
					Name:         "Len",
					Offset:       0,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Cap",
					Offset:       4,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Data",
					Offset:       8,
					TypeName:     "int",
					FullTypeName: "int*",
					ArrayDim:     0,
				},
			},
		},
		{
			Name: "golang::goslice<golang::gostring>",
			Members: []MemberDescr{
				{
					Name:         "Len",
					Offset:       0,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Cap",
					Offset:       4,
					TypeName:     "int",
					FullTypeName: "int",
					ArrayDim:     0,
				},
				{
					Name:         "Data",
					Offset:       8,
					TypeName:     "golang::gostring",
					FullTypeName: "golang::gostring*",
					ArrayDim:     0,
				},
			},
		},
	} {
		cls := croot.GetClass(table.Name)
		if cls == nil {
			t.Fatalf("could not retrieve croot.Class for %q\n", table.Name)
		}

		cls.Print("")

		for _, m := range table.Members {

			dm := cls.GetDataMember(m.Name)
			if dm == nil {
				t.Fatalf(
					"could not retrieve %q data member from %q\n",
					m.Name,
					table.Name,
				)
			}

			offset := dm.GetOffset()
			if m.Offset != offset {
				t.Fatalf(
					"%s.%s: offset error. want=%d. got=%d\n",
					table.Name, m.Name, m.Offset,
					offset,
				)
			}

			tname := dm.GetTypeName()
			if m.TypeName != tname {
				t.Fatalf(
					"%s.%s: typename error. want=%q. got=%q\n",
					table.Name, m.Name, m.TypeName,
					tname,
				)
			}

			ftname := dm.GetFullTypeName()
			if m.FullTypeName != ftname {
				t.Fatalf(
					"%s.%s: full-typename error. want=%q. got=%q\n",
					table.Name, m.Name, m.FullTypeName,
					ftname,
				)
			}

			ndims := dm.GetArrayDim()
			if m.ArrayDim != ndims {
				t.Fatalf(
					"%s.%s: ndims error. want=%d. got=%d\n",
					table.Name, m.Name, m.ArrayDim,
					ndims,
				)
			}
		}
	}
}
