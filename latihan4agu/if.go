package main 

import "fmt" 

func main () {
	bola := "besar"

	if bola == "kecil" {
		fmt.Println("bola pingpong")
	} else if bola == "besar" {
		fmt.Println("bola basket")
	} else {
		fmt.Println("bola")
	}

	// if length := len(bola); length > 5 {
	// 	fmt.Println("nama panjang")
	// } else {
	// 	fmt.Println("nama benar")
	// }

	var length = len(bola)
	if length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama benar")
	}
}