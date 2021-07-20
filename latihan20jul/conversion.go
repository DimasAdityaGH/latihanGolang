package main

import "fmt"

func main () {
	var ubah = "nama"
	var f = ubah[0]
	var n = string(f)
	fmt.Println(n)

	var nama = "dimas"
	nama = "Aditya"
	fmt.Println(nama)

	var nilai32 int32 = 100000
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)
}