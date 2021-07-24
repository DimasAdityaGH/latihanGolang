package main

import "fmt"

func main () {
	hari := []string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"minggu",
	}
	fmt.Println(hari)
	slice1 := hari[1:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := hari[5:]
	//append
	slice3 := append(slice2, "senin lagi")
	fmt.Println(slice3)
	//newslice

	newSlice := make([]string, 2, 5)
	newSlice[0] = "dimas"
	newSlice [1]= "fifi"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}