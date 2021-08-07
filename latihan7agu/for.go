package main 

import "fmt" 

func main () {
	//for loops
	names := []string {"dimas", "jenny", "alif"}

	for i, value := range names {
		fmt.Println(i, "=", value)
	}

	nama := []string {"afifah", "masturina"}
	for r := 0; len(nama) > r; r++ {
		fmt.Println(r)
	}

	for r, value := range nama {
		fmt.Println(r, "=", value)
	}
}




































































































































