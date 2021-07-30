package main 

import "fmt" 

func main () {
	var nama = "dimas"

	if nama == "dimas" {
		fmt.Println("dimas")
	} else if nama == "jenny" {
		fmt.Println("jenny")
	} else {
		fmt.Println("nama")
	}

	// length := len(nama); if length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	if length := len(nama); length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama benar")
	}
}