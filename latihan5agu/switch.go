package main 

import "fmt" 

func main () {
	var nama = "a"

	switch nama {
	case "dimas":
		fmt.Println("dimas")
	case "aditya":
		fmt.Println("aditya")
	default:
		fmt.Println("no name")
	}

	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("name panjang")
	case false:
		fmt.Println("nama pendek")
	}

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("name very long")
	case length > 5:
		fmt.Println("name long")
	default:
		fmt.Println("good name")
	}
}