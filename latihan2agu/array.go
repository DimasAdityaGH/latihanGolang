package main 

import "fmt" 

func main () {
	hari := [...]string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(hari)

	var angka [3]string
	angka[0] = "satu"
	angka[1] = "dua"
	angka[2] = "tiga"
	fmt.Println(angka)
}
