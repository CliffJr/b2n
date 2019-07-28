# Package b2n provides functions for parsing byte arrays and converting them to

You can find full documentation [HERE](https://godoc.org/github.com/filipkroca/b2n#example-ValidateIMEI)

## Bytes slice to unsigned integer

### ParseBs2Uint8  

ParseBs2Uint8 takes a pointer to a byte slice, start byte and returns parsed Uint8

Performance per core:   0.46 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint8)

### ParseBs2Uint16  

ParseBs2Uint16 takes a pointer to a byte slice, start byte and returns parsed Uint16

Performance per core:   3.35 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint16)

### ParseBs2Uint32  

ParseBs2Uint32 takes a pointer to a byte slice, start byte and returns parsed Uint32

Performance per core:   4.97 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint32)

## Bytes slice encoded with Two's complement to signed integer  

Read more here [Two's complement](https://en.wikipedia.org/wiki/Two%27s_complement)  

### ParseBs2Int8TwoComplement  

ParseBs2Int8TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int8

Performance per core:   0.24 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int8TwoComplement)

### ParseBs2Int16TwoComplement  

ParseBs2Int16TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int16

Performance per core:   4.52 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int16TwoComplement)

### ParseBs2Int32TwoComplement  

ParseBs2Int32TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int32

Performance per core:   7.48 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int32TwoComplement) 

### ParseBs2Int64TwoComplement  

ParseBs2Int64TwoComplement takes a pointer to a byte slice coded with Two Complement Arithmetic, start byte and returns parsed signed Int64

Performance per core:   11.1 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Int64TwoComplement) 

## IMEI validator

ValidateIMEI takes pointer to 15 digits long IMEI string, calculate checksum using the Luhn algorithm and return validity as bool

Performance per core 193 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ValidateIMEI)
