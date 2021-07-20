package main

import "fmt"

func main () {
	type str string
	type num int

	var produk str = "susu"
	var harga num = 100000
	fmt.Println(produk)
	fmt.Println(harga)
}