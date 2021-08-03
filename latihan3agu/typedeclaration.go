package main 

import "fmt" 

func main () {
	type str string
	type numb int8

	var nama str = "dimas"
	var nomor numb = 12
	fmt.Println(nama, nomor)
}