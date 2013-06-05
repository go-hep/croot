package main

var (
	go_typemap = map[string]string{
		// C/C++ types
		"void":     "byte", //FIXME
		"uint64_t": "uint64",
		"uint32_t": "uint32",
		"uint16_t": "uint16",
		"uint8_t":  "uint8",
		"uint_t":   "uint32", // uint can not be decoded by encoding/binary.Read
		"int64_t":  "int64",
		"int32_t":  "int32",
		"int16_t":  "int16",
		"int8_t":   "int8",

		"bool":           "bool",
		"char":           "byte",
		"signed char":    "int8",
		"unsigned char":  "byte",
		"short":          "int16",
		"unsigned short": "uint16",
		"int":            "int32",  // int can not be decoded by encoding/binary.Read
		"unsigned int":   "uint32", // uint can not be decoded by encoding/binary.Read

		"char*":       "string",
		"const char*": "string",
		"char const*": "string",

		// FIXME: 32/64 platforms... (and cross-compilation)
		//"long":           "int32",
		//"unsigned long":  "uint32",
		"long":          "int64",
		"unsigned long": "uint64",

		"long long":          "int64",
		"unsigned long long": "uint64",

		"float":  "float32",
		"double": "float64",

		"float complex":  "complex64",
		"double complex": "complex128",

		// FIXME: 32/64 platforms
		//"size_t": "int",
		"size_t":      "int64",
		"std::size_t": "int64",

		// stl
		"std::string": "string",

		"std::ptrdiff_t": "int64",   //FIXME !!
		"std::ostream":   "uintptr", //FIXME !!

		// ROOT types
		"Char_t":   "byte",
		"UChar_t":  "byte",
		"Short_t":  "int16",
		"UShort_t": "uint16",
		"Int_t":    "int32",
		"UInt_t":   "uint32",

		"Seek_t":     "int64",
		"Long_t":     "int64",
		"ULong_t":    "uint64",
		"Float_t":    "float32",
		"Float16_t":  "float32", //FIXME
		"Double_t":   "float64",
		"Double32_t": "float64",

		"Bool_t":    "bool",
		"Text_t":    "byte",
		"Byte_t":    "byte",
		"Version_t": "int16",
		"Option_t":  "byte",
		"Ssiz_t":    "int64",
		"Real_t":    "float32",
		"Long64_t":  "int64",
		"ULong64_t": "uint64",
		"Axis_t":    "float64",
		"Stat_t":    "float64",
		"Font_t":    "int16",
		"Style_t":   "int16",
		"Marker_t":  "int16",
		"Width_t":   "int16",
		"Color_t":   "int16",
		"SCoord_t":  "int16",
		"Coord_t":   "float64",
		"Angle_t":   "float32",
		"Size_t":    "float32",
	}
)
