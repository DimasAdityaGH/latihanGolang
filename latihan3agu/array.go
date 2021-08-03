package main 

import "fmt" 

func main () {
	var hari = [...]string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(hari)

	var libur [2]string
	libur[0] = "sabtu"
	libur[1] = "minggu"

	fmt.Println(libur)
}