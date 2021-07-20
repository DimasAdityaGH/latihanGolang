package main

import "fmt"

func main () {
	//constant

	const (
		firstName = "dimas"
		lastName = "aditya"
	)
	fmt.Println(firstName, lastName)

	const nama = "Esa dimas aditya"
	fmt.Println(nama)// menggunakan constant caranya sama seperti variabel, yang membedakanya adalah constant nilainya tetap
}