package main

import "fmt"

func main () {
	newSlice := make([]string, 2, 5)
	newSlice[0] = "DONI"
	newSlice[1] = "dian"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

}