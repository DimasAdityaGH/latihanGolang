package main 

import "fmt" 

func main () {
	// for i := 1; i < 10; i++ {
	// 	fmt.Println("perulangan ke", i)
	// }


	//for range	

	slice := []string{"esa", "dimas", "aditya"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println(i, "=", value)
	}

	names := []string {"john", "alex", "yoshida"}
	for index, value := range names {
		fmt.Println(index, value)
	}

	slice2 := make(map[string]string)
	slice2["nama"] = "dimas"
	slice2["code"] = "112"

	for key, value := range slice2 {
		fmt.Println(key, "=", value)
	}
}