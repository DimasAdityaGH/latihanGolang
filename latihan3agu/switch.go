package main

import ("fmt")


func main () {
	var nama = "fifi"

	switch nama {
	case "dimas":
		fmt.Println("dimasaditya_")
	case "fifi":
		fmt.Println("afifahmasturina_")
	default:
		fmt.Println("default")
	}

	// length := len(nama)
	// switch {
	// case length > 10:
	// 	fmt.Println("nama terlalu panjang")
	// case length > 5:
	// 	fmt.Println("nama cukup")
	// default:
	// 	fmt.Println("nama bagus")
	// }

	length := len(nama)
	switch length > 5 {
	case true:
		fmt.Println("nama panjang")
	case false:
		fmt.Println("nama bagus")
	}

}