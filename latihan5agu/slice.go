package main 

import "fmt" 

func main () {
	slice := make([]string, 2, 5)
	slice[0] = "satu"
	slice[1] = "dua"
	fmt.Println(slice)

	sliceap := append(slice, "tiga")
	fmt.Println(sliceap)

	copyslice := make([]string, len(slice), cap(slice))
	copy(copyslice, slice)
	fmt.Println(copyslice)
}