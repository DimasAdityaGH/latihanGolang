package main

import "fmt"

func main () {
	//if expression
	nama := "dimas"

	if nama == "dimas" {
		fmt.Println("iam dimas")
	} else if nama == "fifi" {
		fmt.Println("iam fifi")
	} else {
		fmt.Println("i am anonim")
	}

	//short statement
	length := len(nama); if length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}