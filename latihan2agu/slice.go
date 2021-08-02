package main 

import "fmt" 

func main () {
	newSlice := make([]string, 2, 4)
	newSlice[0] = "afifah"
	newSlice[1] = "masturina"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	apend := append(newSlice, "cantik")
	fmt.Println(apend)
}