package main 

import "fmt" 

func main () {
	slice1 := make([]string, 2, 5)
	slice1[0] = "example"
	slice1[1] = "example"

	apen := append(slice1, "example lagi")
	fmt.Println(apen)

	copySlice := make([]string, len(slice1), cap(slice1))
	copy(copySlice, slice1)
	fmt.Println(copySlice)
}