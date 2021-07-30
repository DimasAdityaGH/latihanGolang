package main 

import "fmt" 

func main () {
	hewan := make(map[string]string)
	hewan["nama"] = "kucing"
	hewan["makanan"] = "ikan"
	hewan["warna"] = "hitam"
	hewan["salah"] = "ups"

	fmt.Println(hewan)

	//ubah
	hewan["warna"] = "oren"
	fmt.Println(hewan)

	//tambah
	hewan["kaki"] = "empat"
	fmt.Println(hewan)

	//hapus
	delete(hewan, "salah")
	fmt.Println(hewan)
}