package cgentype_test

import (
	"reflect"
	"testing"

	"github.com/go-hep/croot/cgentype"
)

type S1 struct {
	Field1 float64
}

type S2 struct {
	Field1 float32
}

type S3 struct {
	Field1 S1
	Field2 S2
}

type S4 struct {
	Field1 string
	Field2 []float64
}

type S5 struct {
	Field1 [2]string
	Field2 [2][3]string
	Field3 [1][2][3]string
}

type S6 struct {
	Field1 []string
}

func TestGenerate(t *testing.T) {
	for _, test := range []struct {
		Value interface{}
		Exp   string
	}{
		{
			Value: int(0),
			Exp:   "int; // int Int\n",
		},
		{
			Value: float64(0),
			Exp:   "double; // float64 Double\n",
		},
		{
			Value: S1{},
			Exp: `struct S1 {
	double Field1; // float64 Double
};
`,
		},
		{
			Value: S2{},
			Exp: `struct S2 {
	float Field1; // float32 Float
};
`,
		},
		{
			Value: S3{},
			Exp: `struct S3 {
	S1 Field1; // S1 Struct
	S2 Field2; // S2 Struct
};
`,
		},
		{
			Value: S4{},
			Exp: `struct S4 {
	::golang::gostring Field1; // string Ptr
	::golang::goslice< double > Field2; //  Slice
};
`,
		},
		{
			Value: S5{},
			Exp: `struct S5 {
	::golang::gostring Field1[2]; //  Array
	::golang::gostring Field2[2][3]; //  Array
	::golang::gostring Field3[1][2][3]; //  Array
};
`,
		},
		{
			Value: S6{},
			Exp: `struct S6 {
	::golang::goslice< ::golang::gostring > Field1; //  Slice
};
`,
		},
	} {
		rt := reflect.TypeOf(test.Value)
		str := cgentype.Generate(rt)
		if str != test.Exp {
			t.Errorf(
				"error for %v:\n got=%q\nwant=%q\n",
				test.Value,
				str,
				test.Exp,
			)
		}
	}

}
