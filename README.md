# Package b2n provides functions for parsing byte arrays and converting them to

You can find full documentation [HERE](https://godoc.org/github.com/filipkroca/b2n#example-ValidateIMEI)

## Bytes slice to unsigned integer

### ParseBs2Uint8  

ParseBs2Uint8 takes a pointer to byte slice, start byte and returns parsed Uint8

Performance per core:   0.46 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ParseBs2Uint8)

### uint16

### uint32

## [Two's complement](https://en.wikipedia.org/wiki/Two%27s_complement)  

### int8  

### int16  

### int32  

### int64  

## IMEI validator

ValidateIMEI takes pointer to 15 digits long IMEI string, calculate checksum using the Luhn algorithm and return validity as bool

Performance per core 193 ns/op, 0 B/op, 0 allocs/op
[EXAMPLE](https://godoc.org/github.com/filipkroca/b2n#example-ValidateIMEI)
