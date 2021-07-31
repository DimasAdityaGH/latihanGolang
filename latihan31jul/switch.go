package main 

import "fmt" 

func main () {
	var name = "uus"

	switch name {
	case "dimas":
		fmt.Println("dimas")
	case "uus":
		fmt.Println("uus")
	default:
		fmt.Println("esa dimas aditya")
	}

	// length := len(name)
	// switch length > 5 {
	// case true:
	// 	fmt.Println("name long")
	// case false:
	// 	fmt.Println("name good")
	// }

	length := len(name) 
	switch {
	case length > 10:
		fmt.Println("name is very long")
	case length > 5:
		fmt.Println("name is long")
	default: 
		fmt.Println("name is good")
	}


}