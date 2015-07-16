package cgentype_test

import (
	"fmt"
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

func TestGenerate(t *testing.T) {
	for _, test := range []struct {
		Value interface{}
	}{
		{
			Value: int(0),
		},
		{
			Value: float64(0),
		},
		{
			Value: S1{},
		},
		{
			Value: S2{},
		},
		{
			Value: S3{},
		},
		{
			Value: S4{},
		},
		{
			Value: S5{},
		},
	} {
		fmt.Printf("===\n")
		rt := reflect.TypeOf(test.Value)
		str := cgentype.Generate(rt)
		fmt.Printf("code=\n%s\n", str)
	}

}
