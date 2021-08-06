package main 

import "fmt" 

func main () {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	slice := []string {"dimas", "dwi"}

	for index, value := range slice {
		fmt.Println(index, value)
	}
}