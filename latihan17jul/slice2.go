package main

import "fmt"

func main() {
	bulan := [...]string {"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	var slice1 = bulan[4:7]
	fmt.Println(slice1)

	//len function
	fmt.Println(len(slice1))
	//cap function
	fmt.Println(cap(slice1))

	slice1[1] = "juni diganti" //index 1 adalah juni
	fmt.Println(slice1)

	slice2 := bulan[10:]
	fmt.Println(slice2)

	//function append
	slice3 := append(slice2, "data baru")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	//make
	newSlice := make([]string, 2, 5)
	newSlice[0] = "dimas"
	newSlice[1] = "aditya"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))


}