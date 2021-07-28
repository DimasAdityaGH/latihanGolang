package main

import "fmt"

func main () {
	var nama = "aditya"

	switch nama {
	case "aditya":
		fmt.Println("aditya")
	case "esa": 
		fmt.Println("esa")
	default:
		fmt.Println("username")
	}

	length := len(nama) 
	switch {
	case length > 10:
		fmt.Println("nama sangat panjang")
	case length > 5:
		fmt.Println("nama panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}