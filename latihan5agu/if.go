package main

import (
	"fmt" 
)


func main () {
	var nama = "world"

	if nama == "world" {
		fmt.Println("hello world!")
	} else if nama == "guys" {
		fmt.Println("hello guys")
	} else {
		fmt.Println("nama")
	}

	if length := len(nama); length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama pendek")
	}

}