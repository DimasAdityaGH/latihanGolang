package main 

import "fmt" 

func main () {
	var nama = "jokowi"

	if nama == "dimas" {
		fmt.Println("iam dimas")
	} else if nama == "joko" {
		fmt.Println("iam joko")
	} else {
		fmt.Println("iam human")
	}

	// var lenth = len(nama); if lenth > 5 {
	// 	fmt.Println("nama panjang")
	// } else {
	// 	fmt.Println("nama bagus")
	// }

	if length := len(nama); length > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama bagus")
	}
}