package main

import "fmt"

func main () {
	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)

	slice := days[1:6]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice1 :=  []string {"satu", "dua"}
	fmt.Println(slice1)

	sliceappend := append(slice1, "tiga")
	fmt.Println(sliceappend)

	sliceNew := make([]string, 2, 5)
	sliceNew[0] = "example"
	sliceNew[1] = "example"
	fmt.Println(sliceNew)
	fmt.Println(len(sliceNew))
	fmt.Println(cap(sliceNew))

	copyS := make([]string, len(sliceNew), cap(sliceNew))
	copy(copyS, sliceNew)
	fmt.Println(copyS)
	fmt.Println(len(copyS))
	fmt.Println(cap(copyS))
}