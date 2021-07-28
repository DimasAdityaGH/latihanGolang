package main

import "fmt"

func main () {
	var name = "dimas"

	switch name {
	case "dimas":
		fmt.Println("hello dimas")
	case "john":
		fmt.Println("hello john")
	default:
		fmt.Println("hello world!")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah cukup")
	}
}