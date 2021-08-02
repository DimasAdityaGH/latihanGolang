package main 

import "fmt" 

func main () {
	var hewan = "singa"

	switch hewan {
	case "singa":
		fmt.Println("predator")
	case "rusa":
		fmt.Println("mangsa")
	default:
		fmt.Println("input tidak sesuai")
	}

	// length := len(hewan)
	// switch length > 5 {
	// case true:
	// 	fmt.Println("kata lebih dari 5")
	// case false:
	// 	fmt.Println("kata cukup")
	// }

	// switch length := len(hewan); {
	// case length > 10:
	// 	fmt.Println("kata lebih dari 10")
	// case length > 5:
	// 	fmt.Println("kata lebih dari 5")
	// }
	
	length := len(hewan)
	switch {
	case length > 5:
		fmt.Println("nama panjang")
	case length > 10:
		fmt.Println("nama sangat panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}