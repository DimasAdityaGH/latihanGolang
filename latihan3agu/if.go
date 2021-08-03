package main 

import "fmt" 

func main () {
	var nama = "fifi"

	if nama == "dimas" {
		fmt.Println("esa dimas aditya")
	} else if nama == "fifi" {
		fmt.Println("afifah masturina")
	} else {
		fmt.Println("nama anda")
	}

	// var length = len(nama); if length > 5 {
	// 	fmt.Println("kata lebih dari 5")
	// } else {
	// 	fmt.Println("kata cukup")
	// }

	if length := len(nama); length > 5 {
		fmt.Println("nama lebih dari 5")
	} else {
		fmt.Println("nama sudah cukup")
	}
}