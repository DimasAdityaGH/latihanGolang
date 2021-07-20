package main

import "fmt"

func main () {
	var days [7]string 
	days[0] = "senin"
	days[1] = "selasa"
	days[2] = "rabu"
	days[3] = "kamis"
	days[4] = "jum'at"
	days[5] = "sabtu"
	days[6] = "minggu"
	fmt.Println(days)

	nominal := [...]int {
		500,
		1000,
		10000,
		20000,
		50000,
		100000,
	}
	fmt.Println(nominal)
	fmt.Println(nominal[0])
	fmt.Println(len(nominal))
	fmt.Println(cap(nominal))
}