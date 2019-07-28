# Package b2n provides functions for parsing byte arrays and converting them to

Certain purpose:

Package b2n was created for parsing data from [Teltonika](https://wiki.teltonika.lt/view/Codec#Codec_8_Extended) UDP packets, package can be used for parsing values from data streams.

For example have binary packet bs
`
var dataString = "0x00A1CAFE001B000F3335363330373034323434313031338E010000013FEBDD19C8000F0E9FF0209A718000690000120000001E09010002000300040016014703F0001504C8000C0900910A00440B004D130044431555440000B5000BB60005422E9B180000CD0386CE000107C700000000F10000601A460000013C4800000BB84900000BB84A00000BB84C00000000024E0000000000000000CF000000000000000001"

bs, _ := hex.DecodeString(testDataString)

`

Full documentation [HERE](https://godoc.org/github.com/filipkroca/b2n#example-ValidateIMEI)

## Bytes slice to unsigned integer

### ParseBs2Uint8  

ParseBs2Uint8 takes a pointer to a byte slice, start byte and returns parsed Uint8

Performance per core:   0.46 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Uint8)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint8)

### ParseBs2Uint16  

ParseBs2Uint16 takes a pointer to a byte slice, start byte and returns parsed Uint16

Performance per core:   3.35 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Uint16)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint16)

### ParseBs2Uint32  

ParseBs2Uint32 takes a pointer to a byte slice, start byte and returns parsed Uint32

Performance per core:   4.97 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Uint32)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint32)

## Bytes slice encoded with Two's complement to signed integer  

Read more here [Two's complement](https://en.wikipedia.org/wiki/Two%27s_complement)  

### ParseBs2Int8TwoComplement  

ParseBs2Int8TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int8

Performance per core:   0.24 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Int8TwoComplement)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int8TwoComplement)

### ParseBs2Int16TwoComplement  

ParseBs2Int16TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int16

Performance per core:   4.52 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Int16TwoComplement)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int16TwoComplement)

### ParseBs2Int32TwoComplement  

ParseBs2Int32TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int32

Performance per core:   7.48 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Int32TwoComplement)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int32TwoComplement) 

### ParseBs2Int64TwoComplement  

ParseBs2Int64TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int64

Performance per core:   11.1 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ParseBs2Int64TwoComplement)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int64TwoComplement) 

## IMEI validator

ValidateIMEI takes pointer to 15 digits long IMEI string, calculate checksum using the Luhn algorithm and return validity as bool

Performance per core 193 ns/op, 0 B/op, 0 allocs/op

[DOCUMENTATION](https://godoc.org/github.com/filipkroca/b2n#ValidateIMEI)
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ValidateIMEI)
