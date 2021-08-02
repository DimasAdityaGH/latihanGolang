package main 

import "fmt" 

func main () {
	var bahasa = "indonesia"

	if bahasa == "indonesia" {
		fmt.Println("indonesia")
	} else if bahasa == "jawa" {
		fmt.Println("jawa")
	} else {
		fmt.Println("default")
	}

	//short statement

	// if length := len(bahasa); length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	//no short

	var length = len(bahasa)
	if length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama benar")
	}


}