package main

import "fmt"

func main () {
	var nilaiUjian = 90
	var nilaiPraktek = 70

	var hasilNilaiUjian = nilaiUjian >= 90
	var hasilNilaiPraktek = nilaiPraktek >= 75

	var hasil = hasilNilaiUjian || hasilNilaiPraktek
	hasil2 := hasilNilaiUjian && hasilNilaiPraktek
	fmt.Println(hasil)
	fmt.Println(hasil2)
}