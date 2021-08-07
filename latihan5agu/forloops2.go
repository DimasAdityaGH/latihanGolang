package main 

import "fmt" 

func main () {
	slice := []string {"python", "golang", "javascript", "flutter"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(i)
	}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	slice1 := []string {"html", "css", "js"}
	for index, value := range slice1 {
		fmt.Println(index, value)
	}

	map1 := make(map[string]string)
	map1["nama"] = "dimas"
	map1["kelas"] = "12"
	for key, value := range map1 {
		fmt.Println(key, value)
	}	


	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}
}