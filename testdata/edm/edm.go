// Copyright 2016 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package edm exposes a few simple types for testing ROOT I/O.
package edm

type Det struct {
	E float64
	T float64
}

type Event struct {
	I      int64
	A      Det
	B      Det
	ArrayI [2]int64
	ArrayD [2]float64
}

type DataSlice struct {
	I     int64
	Data  float64
	Slice []float64
}

type DataArray struct {
	I      int64
	Data   float64
	Array  [2]float64
	NdArrI [2][2]int64
	NdArrF [2][2]float64
}

type DataString struct {
	I      int64
	Data   float64
	String string
}

type DataStrings struct {
	I       int64
	Data    float64
	Strings []string
}

type DataStringArray struct {
	I      int64
	Data   float64
	Array  [2]string
	NdArrS [2][2]string
}
