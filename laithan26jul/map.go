package main

import (
	"fmt"
)

func main () {
	var book map[string]string = map[string]string{
		"nama": "learn go",
		"penulis": "john",
	}
	fmt.Println(book)

	//mengambil data dimap dengan key
	fmt.Println(book["nama"])

	//membuat map baru dengan make
	buku := make(map[string]string)
	buku["nama"] = "sndadno"
	buku["penulis"] = "doasn"
	fmt.Println(buku)
	//mencari len di map
	fmt.Println(len(buku))
	//mengubah data dimap dengan key
	buku["nama"] = "learn python"
	fmt.Println(buku)
	fmt.Println(buku["nama"])

	data := make(map[string]string)
	data["important"] = "1e13n"
	data["false"] = "123e"

	delete(data, "false")
	fmt.Println(data)

}