package main

import "fmt"

func main() {
	type str string
	type numb int8

	var produk str = "jus"
	var harga numb = 127

	fmt.Println("produk:", produk)
	fmt.Println("harga:", harga)
}