package main 

import "fmt" 

func main () {
	var nama = "abcd"
	var e = nama[2]
	var eString = string(e)
	fmt.Println(eString)


	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)
	fmt.Println(nilai64)
}