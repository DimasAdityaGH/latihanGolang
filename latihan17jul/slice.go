package main

import (
	"fmt"
)

func main() {
	months := [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	// function len
	fmt.Println(len(slice1))
	//function cap
	fmt.Println(cap(slice1))

	// months[5] = "diubah"
	// fmt.Println(months)

	slice1[0] = "mei lagi"
	fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	//function append
	var slice3 = append(slice2, "eko")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//make
	newSlice := make([]string, 2, 5)

	newSlice[0] = "dimas"
	newSlice[1] = "adit"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copyslice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5} //ini adalah array
	iniSlice := []int{1, 2, 3, 4, 5} //ini adalah slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}