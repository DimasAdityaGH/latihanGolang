package main 

import "fmt" 

func main () {
	var nama = "dimasaditya"

	switch nama {
	case "dimas":
		fmt.Println("dimas")
	case "dimasaditya":
		fmt.Println("dimas aditya")
	default:
		fmt.Println("a")
	}

	// switch length := len(nama); length > 5 {
	// case true:
	// 	fmt.Println("name is long")
	// case false:
	// 	fmt.Println("name is true")
	// }

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("name very long")
	case length > 5:
		fmt.Println("name is long")
	default:
		fmt.Println("name is normal")
	}
} 
