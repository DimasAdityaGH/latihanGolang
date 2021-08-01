package main 

import "fmt" 

func main () {
	nama := ""

	if nama == "dimas" {
		fmt.Println("ini dimas")
	} else if nama == "fifi" {
		fmt.Println("ini fifi")
	} else {
		fmt.Println("ini aku")
	}

	// length := len(nama)
	// if length > 5 {
	// 	fmt.Println("mama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	if length := len(nama); length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama pendek")
	}
}