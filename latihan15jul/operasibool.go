package main

import "fmt"

func main() {
	var NilaiUjian = 100
	var NilaiPraktek = 90

	var kkmUjian = 80
	var kkmPraktek = 75

	var hasilUjian = NilaiUjian >= kkmUjian
	var hasilPraktek = NilaiPraktek >= kkmPraktek

	var lulus = hasilUjian || hasilPraktek

	fmt.Println(lulus)
	fmt.Println(hasilUjian && hasilPraktek)

	var benar bool = true
	fmt.Println(!benar)
}