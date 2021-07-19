package main

import "fmt"

func main () {
	days := []string {"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	slice1 := days[2:5]
	fmt.Println(slice1)
	//mencari panjang menggunakan function len
	fmt.Println(len(slice1))
	//mencari kapasitas menggunakan cap
	fmt.Println(cap(slice1))

	slice2 := days[5:]
	//append
	slice3 := append(slice2, "senin lagi")
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "dimas"
	newSlice[1] = "fifi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copyslice := make([]string, len(newSlice), cap(newSlice))
	copy(copyslice, newSlice)
	fmt.Println(copyslice)

	var iniSlice = []string {
		"example",
		"example",
	}
	fmt.Println(iniSlice)

	var iniArray = []string {
		"example1",
		"example1",
	}
	fmt.Println(iniArray)
}