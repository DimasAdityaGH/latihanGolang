package main

import (
	"fmt" 
)


func main () {
	var dimas = map[string]string {
		"umur": "17",
		"kelas": "12",
	}
	fmt.Println(dimas)

	hewan := make(map[string]string)
	hewan["nama"] = "harimau"
	hewan["makanan"] = "telur"

	fmt.Println(hewan)

	hewan["tempat"] = "hutan"
	hewan["warna"] = "oren"
	fmt.Println(hewan)

	hewan["warna"] = "putih"
	fmt.Println(hewan)

	delete(hewan, "warna")
	fmt.Println(hewan)
}