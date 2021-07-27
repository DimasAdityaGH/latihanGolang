package main

import (
	"fmt"
)

func main () {
	//tipe data map
	buku := map[string]string {
		"title": "golang",
		"penulis": "john",
	}
	fmt.Println(buku)
	fmt.Println(len(buku))

	books := make(map[string]string)
	books["name"] = "python easy"
	books["page"] = "112"
	//mengubah data dengan map key
	books["name"] = "python is awesome"
	fmt.Println(books)
	fmt.Println(len(books))

	//menghapus data di map dengan key
	delete(books, "page")
	fmt.Println(books)
	//menambah data di map dengan key
	books["halaman"] = "123"
	fmt.Println(books)

}