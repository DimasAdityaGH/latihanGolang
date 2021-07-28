package main

import "fmt"

func main () {
	var salam = "world"

	if salam == "world" {
		fmt.Println("hello, world")
	} else if salam == "you" {
		fmt.Println("hello, you")
	} else if salam == "bro" {
		fmt.Println("hello, bro!")
	} else {	
		fmt.Println("hello")
	}

	// var length = len(salam); if length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	if length := len(salam); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}