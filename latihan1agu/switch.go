package main 

import "fmt" 

func main () {
	var nama = "dimas"

	switch nama {
	case "dimas":
		fmt.Println("hi dimas")
	case "afifah":
		fmt.Println("hi afifah")
	default:
		fmt.Println("hi baby")
	}

	// length := len(nama)
	// switch length > 5 {
	// case true:
	// 	fmt.Println("nama panjang")
	// case false:
	// 	fmt.Println("nama bagus")
	// }


	// switch length := len(nama); length >= 5 {
	// case true:
	// 	fmt.Println("nama panjang")
	// case false: 
	// 	fmt.Println("nama pendek")
	// }

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama panjang")
	default:
		fmt.Println("nama bagus")
	}


}