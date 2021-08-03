package main 

import "fmt" 

func main () {
	var nama = "fifi"

	switch nama {
	case "dimas":
		fmt.Println("dimas")
	case "fifi":
		fmt.Println("afifah")
	default:
		fmt.Println("null")
	}

	// length := len(nama)
	// switch length > 5 {
	// case true:
	// 	fmt.Println("nama panjang")
	// case false:
	// 	fmt.Println("nama pendek")
	// }

	// switch length := len(nama); length > 5 {
	// case true:
	// 	fmt.Println("nama panjang")
	// case false:
	// 	fmt.Println("nama pendek")
	// }

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("nama lengkap")
	case length > 5:
		fmt.Println("nama pendek")
	default:
		fmt.Println("nama bagus")
	}
}