// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Package b2n provides functions for parsing byte arrays and conversion into another data types
//home page https://github.com/filipkroca/b2n
package b2n

import (
	"log"
	"strconv"
)

//ParseBs2Uint8 a pointer to a byte slice, offset and returns parsed Uint8
func ParseBs2Uint8(bs *[]byte, offset int32) uint8 {
	//convert hex byte to Uint8
	return uint8((*bs)[offset])
}

//ParseBs2Uint16 a pointer to a byte slice, offset and returns parsed int16
func ParseBs2Uint16(bs *[]byte, offset int32) uint16 {
	var sum uint16
	var order uint32
	//convert hex byte slice to Uint64
	for i := offset + 1; i >= offset; i-- {
		//shift to the left by 8 bits every cycle
		sum += uint16((*bs)[i]) << order
		order += 8
	}
	return sum
}

//ParseBs2Uint32 a pointer to a byte slice, offset and returns parsed int32
func ParseBs2Uint32(bs *[]byte, offset int32) uint32 {
	var sum uint32
	var order uint32
	//convert hex byte slice to Uint64
	for i := offset + 3; i >= offset; i-- {
		//shift to the left by 8 bits every cycle
		sum += uint32((*bs)[i]) << order
		order += 8
	}
	return sum
}

//ParseBs2Uint64 a pointer to a byte slice, offset, stop byte and returns parsed int64
func ParseBs2Uint64(bs *[]byte, offset int32) uint64 {
	var sum uint64
	var order uint32
	//convert hex byte slice to Uint64
	for i := offset + 7; i >= offset; i-- {
		//shift to the left by 8 bits every cycle
		sum += uint64((*bs)[i]) << order
		order += 8
	}
	return sum
}

//ParseBs2Int8TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, offset and returns parsed signed Int8
func ParseBs2Int8TwoComplement(bs *[]byte, offset int32) int8 {
	var sum int8
	var signed bool
	//mask last Byte with mask (1000 0000) then shift by 7 bits and check sign bit
	if (*bs)[offset]&0x80>>7 == 1 {
		signed = true
	}
	//convert hex byte slice to int8
	cb := (*bs)[offset]
	//if signed do a XOR operation on every single Byte
	if signed {
		cb ^= 0xFF
	}
	sum = int8(cb)
	//finally if signed, increment with complement 1 and multiply by -1
	if signed {
		sum++
		sum = sum * -1
	}
	return sum
}

//ParseBs2Int16TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, offset and returns parsed signed Int16 coded with Two Complement Arithmetic
func ParseBs2Int16TwoComplement(bs *[]byte, offset int32) int16 {
	var sum int16
	var order uint16
	var signed bool
	//mask last Byte with mask (1000 0000) then shift by 7 bits and check sign bit
	if (*bs)[offset]&0x80>>7 == 1 {
		signed = true
	}
	//convert hex byte slice to int16
	for i := offset + 1; i >= offset; i-- {
		cb := (*bs)[i]
		//if signed do a XOR operation on every single Byte
		if signed {
			cb ^= 0xFF
		}
		//shift to the left by 8 bits every cycle
		sum += int16(cb) << order
		order += 8
	}
	//finally if signed, increment with complement 1 and multiply by -1
	if signed {
		sum++
		sum = sum * -1
	}
	return sum
}

//ParseBs2Int32TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, offset and returns parsed signed Int32 coded with Two Complement Arithmetic
func ParseBs2Int32TwoComplement(bs *[]byte, offset int32) int32 {
	var sum int32
	var order uint32
	var signed bool
	//mask last Byte with mask (1000 0000) then shift by 7 bits and check sign bit
	if (*bs)[offset]&0x80>>7 == 1 {
		signed = true
	}
	//convert hex byte slice to int32
	for i := offset + 3; i >= offset; i-- {
		cb := (*bs)[i]
		//if signed do a XOR operation on every single Byte
		if signed {
			cb ^= 0xFF
		}
		//shift to the left by 8 bits every cycle
		sum += int32(cb) << order
		order += 8
	}
	//finally if signed, increment with complement 1 and multiply by -1
	if signed {
		sum++
		sum = sum * -1
	}
	return sum
}

//ParseBs2Int64TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, offset and returns parsed signed Int64 coded with Two Complement Arithmetic
func ParseBs2Int64TwoComplement(bs *[]byte, offset int32) int64 {
	var sum int64
	var order uint32
	var signed bool
	//mask last Byte with mask (1000 0000) then shift by 7 bits and check sign bit
	if (*bs)[offset]&0x80>>7 == 1 {
		signed = true
	}
	//convert hex byte slice to int32
	for i := offset + 7; i >= offset; i-- {
		cb := (*bs)[i]
		//if signed do a XOR operation on every single Byte
		if signed {
			cb ^= 0xFF
		}
		//shift to the left by 8 bits every cycle
		sum += int64(cb) << order
		order += 8
	}
	//finally if signed, increment with complement 1 and multiply by -1
	if signed {
		sum++
		sum = sum * -1
	}
	return sum
}

//ValidateIMEI takes pointer to 15 digits long IMEI string, calculate checksum using the Luhn algorithm and return validity as bool
func ValidateIMEI(imei *string) bool {
	bs := []byte((*imei))

	if len(bs) != 15 {
		log.Printf("Should validate only 15chars long Imei, got %v", len(bs))
		return false
	}

	parsed, err := strconv.ParseInt(string(bs[len(bs)-1]), 10, 8)
	if err != nil {
		log.Printf("Unable to parse IMEI digits %v", err)
		return false
	}
	checkSumDigit := int8(parsed)
	var checkSum uint64

	//make buffer array for Luhn algorithm with len 14 bytes and cap 31 bytes
	digits := make([]uint8, 14, 31)

	//count Luhn algorithm
	for i := 0; i < 14; i++ {

		parsed, err = strconv.ParseInt(string(bs[i]), 10, 8)
		if err != nil {
			log.Printf("Unable to parse IMEI digits %v", err)
			return false
		}

		digits[i] = uint8(parsed)
		if i%2 != 0 {
			digits[i] = digits[i] * 2
		}

		if digits[i] >= 10 {
			digits = append(digits, 1)
			digits[i] = digits[i] % 10
		}
	}

	for _, val := range digits {
		checkSum += uint64(val)
	}

	//return divider to 10 is same as checkSumDigit
	return ((10 - checkSum%10) == uint64(checkSumDigit))
}
