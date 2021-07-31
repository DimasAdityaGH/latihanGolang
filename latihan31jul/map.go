package main 

import "fmt" 

func main () {
	hewan := map[string]string {
		"nama": "kucing",
		"makanan": "ikan",
		"kaki": "4",
	}

	fmt.Println(hewan)

	//tambah
	hewan["warna"] = "oren"
	fmt.Println(hewan)

	//ubah
	hewan["kaki"] = "empat"
	fmt.Println(hewan)

	//hapus
	delete(hewan, "kaki")
	fmt.Println(hewan)
}