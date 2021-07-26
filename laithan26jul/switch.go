package main

import "fmt"

func main () {
	name := "false"

	switch name {
	case "john":
		fmt.Println("hello, john")
	case "jord":
		fmt.Println("hello, jord")
	default:
		fmt.Println("hello, world")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}
}