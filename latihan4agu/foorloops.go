package main 

import "fmt" 

func main () {
	// dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
	//salah satu fitur perulangan adalah for loops

	// for statement
	//dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan di for
	//init statement, yaitu statement sebelum for dieksekusi
	//post statement, yaitu statement yang selalu dieksekusi diakhir setiap pengulangan

	//for range
	//for bisa digunakan untuk melakukan literasi terhadap semu data collection
	//data collection contohnya array, slice, map
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("perulangan", counter)
	// 	counter++
	// }

	for counter := 1; counter < 10; counter+= 3 {
		fmt.Println(counter)
	}

	slice := []string{"dimas", "aditya", "esa"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	nama := []string{"afifah", "masturina"}

	for i, value := range nama {
		fmt.Println(i, value)
	}

	


}	