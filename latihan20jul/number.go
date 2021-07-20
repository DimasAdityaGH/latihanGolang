package main

import "fmt"

func main () {
	//type data number 
	var nilai8 int8 = 127 //int8 nilainya 0 - 127
	var nilai16 int16 = 32767 //int16 nilainya 127 -32767
	var nilai32 int32 = 400000 //int32 nilainya 32767 - 2147483647
	var nilai64 int64 = 2121212333 //int64 nilainya 2147483648	- 9223372036854775807	
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	fmt.Println(ui8)
	fmt.Println(ui16)
	fmt.Println(ui32)
	fmt.Println(ui64)

	
}