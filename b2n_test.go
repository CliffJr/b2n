// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package b2n

import (
	"fmt"
	"testing"
)

func ExampleParseBs2Uint8() {
	// sample data
	bs := []byte{0xAB, 0xFF, 0x12}

	// run test
	decoded := ParseBs2Uint8(&bs, 1)

	fmt.Printf("%T %v", decoded, decoded)
	// Output:
	// uint8 255
}

func ExampleParseBs2Uint16() {
	// sample data
	bs := []byte{0xAB, 0xCD, 0xAB, 0xFF}

	// run test
	decoded := ParseBs2Uint16(&bs, 0)

	fmt.Printf("%T %v", decoded, decoded)
	// Output:
	// uint16 43981
}

func ExampleParseBs2Uint32() {
	// sample data
	bs := []byte{0xAB, 0xCD, 0xAB, 0xFF}

	// run test
	decoded := ParseBs2Uint32(&bs, 0)

	fmt.Printf("%T %v", decoded, decoded)
	// Output:
	// uint32 2882382847
}

func ExampleParseBs2Uint64() {
	// sample data
	bs := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	// run test
	decoded := ParseBs2Uint64(&bs, 0)

	fmt.Printf("%T %v", decoded, decoded)
	// Output:
	// uint64 18446744073709551615
}

func ExampleParseBs2Int8TwoComplement() {
	// sample data
	bs := []byte{0xFF, 0x80, 0x51, 0x00}

	// run test
	decoded := ParseBs2Int8TwoComplement(&bs, 0)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int8TwoComplement(&bs, 1)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int8TwoComplement(&bs, 2)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int8TwoComplement(&bs, 3)
	fmt.Printf("%T %v\n", decoded, decoded)

	// Output:
	// int8 -1
	// int8 -128
	// int8 81
	// int8 0
}

func ExampleParseBs2Int16TwoComplement() {
	// sample data
	bs := []byte{0xFF, 0xFF, 0x80, 0x00, 0x7F, 0xFF, 0x00, 0x00}

	// run test
	decoded := ParseBs2Int16TwoComplement(&bs, 0)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int16TwoComplement(&bs, 2)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int16TwoComplement(&bs, 4)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int16TwoComplement(&bs, 6)
	fmt.Printf("%T %v\n", decoded, decoded)

	// Output:
	// int16 -1
	// int16 -32768
	// int16 32767
	// int16 0
}

func ExampleParseBs2Int32TwoComplement() {
	// sample data
	bs := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x00, 0x00, 0x00, 0x7F, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00}

	// run test
	decoded := ParseBs2Int32TwoComplement(&bs, 0)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int32TwoComplement(&bs, 4)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int32TwoComplement(&bs, 8)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int32TwoComplement(&bs, 12)
	fmt.Printf("%T %v\n", decoded, decoded)

	// Output:
	// int32 -1
	// int32 -2147483648
	// int32 2147483647
	// int32 0
}

func ExampleParseBs2Int64TwoComplement() {
	// sample data
	bs := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	// run test
	decoded := ParseBs2Int64TwoComplement(&bs, 0)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int64TwoComplement(&bs, 8)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int64TwoComplement(&bs, 16)
	fmt.Printf("%T %v\n", decoded, decoded)

	// run test
	decoded = ParseBs2Int64TwoComplement(&bs, 24)
	fmt.Printf("%T %v\n", decoded, decoded)

	// Output:
	// int64 -1
	// int64 -9223372036854775808
	// int64 9223372036854775807
	// int64 0
}

func ExampleValidateIMEI() {
	imei := "352094081852920"
	fmt.Println(ValidateIMEI(&imei))
	// Output:
	// true
}

func ExampleParseIMEI() {
	// Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002, IMEI is located starting byte 8

	var bs = []byte{0x00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08}

	imei, err := ParseIMEI(&bs, 8, 15)
	if err != nil {
		fmt.Println("ExampleParseIMEI error", err)
	}
	fmt.Printf("%T %v", imei, imei)

	// Output:
	// string 352094081672179
}

func BenchmarkValidateIMEI(b *testing.B) {
	imei := "352094081852920"
	for i := 0; i < b.N; i++ {
		ValidateIMEI(&imei)
	}
}

func BenchmarkParseIMEI(b *testing.B) {
	// Example packet 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

	var bs = []byte{0x00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08}

	for i := 0; i < b.N; i++ {
		ParseIMEI(&bs, 8, 15)
	}

}

func BenchmarkParseBs2Uint8(b *testing.B) {
	bs := []byte{0xAB, 0xFF, 0x12}
	for i := 0; i < b.N; i++ {
		ParseBs2Uint8(&bs, 1)
	}
}

func BenchmarkParseBs2Uint16(b *testing.B) {
	bs := []byte{0xAB, 0xCD, 0xAB, 0xFF}
	for i := 0; i < b.N; i++ {
		ParseBs2Uint16(&bs, 0)
	}
}

func BenchmarkParseBs2Uint32(b *testing.B) {
	bs := []byte{0xAB, 0xCD, 0xAB, 0xFF}
	for i := 0; i < b.N; i++ {
		ParseBs2Uint32(&bs, 0)
	}
}

func BenchmarkParseBs2Int8TwoComplement(b *testing.B) {
	bs := []byte{0xFF, 0x80, 0x51, 0x00}
	for i := 0; i < b.N; i++ {
		ParseBs2Int8TwoComplement(&bs, 0)
	}
}

func BenchmarkParseBs2Int16TwoComplement(b *testing.B) {
	bs := []byte{0xFF, 0x80, 0x51, 0x00}
	for i := 0; i < b.N; i++ {
		ParseBs2Int16TwoComplement(&bs, 0)
	}
}

func BenchmarkParseBs2Int32TwoComplement(b *testing.B) {
	bs := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	for i := 0; i < b.N; i++ {
		ParseBs2Int32TwoComplement(&bs, 0)
	}
}

func BenchmarkParseBs2Int64TwoComplement(b *testing.B) {
	bs := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	for i := 0; i < b.N; i++ {
		ParseBs2Int64TwoComplement(&bs, 0)
	}
}
