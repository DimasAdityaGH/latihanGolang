package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "dimas"
	newSlice[1] = "aditya"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))
}